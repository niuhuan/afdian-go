package afdian

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

const AfDianAppUrl = "https://ifdian.net/api"

type AppClient struct {
	jar    *cookiejar.Jar
	client *http.Client
}

func RequestApp[R any](c *AppClient, method, path string, form url.Values) (*R, error) {
	finalUrl := AfDianAppUrl + path
	req, err := http.NewRequest(method, finalUrl, nil)
	if err != nil {
		return nil, err
	}
	if method == http.MethodGet {
		if form != nil {
			req.URL.RawQuery = form.Encode()
		}
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Body = ioutil.NopCloser(strings.NewReader(form.Encode()))
	}
	rsp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	buff, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	var response AfdianResponse[R]
	err = json.Unmarshal(buff, &response)
	if err != nil {
		return nil, err
	}
	if response.Ec != 200 {
		return nil, errors.New(response.Em)
	}
	return &response.Data, nil
}

func (c *AppClient) Login(username, password string) (*AuthToken, error) {
	at, err := RequestApp[AuthToken](c, http.MethodPost, "/passport/login", url.Values{
		"account":  {username},
		"password": {password},
	})
	if err == nil {
		c.SetAuthToken(at.AuthToken)
	}
	return at, err
}

func (c *AppClient) MyAccount() (*MyAccount, error) {
	return RequestApp[MyAccount](c, http.MethodGet, "/my/account", nil)
}

func (c *AppClient) Plans(userId string) (*Plans, error) {
	return RequestApp[Plans](
		c,
		http.MethodGet,
		"/creator/get-plans",
		url.Values{
			"user_id":         {userId},
			"album_id":        {""},
			"unlock_plan_ids": {""},
			"diy":             {""},
			"affiliate_code":  {""},
		},
	)
}

func (c *AppClient) PlanSkus(planId string) (*PlanSkus, error) {
	return RequestApp[PlanSkus](
		c,
		http.MethodGet,
		"/creator/get-plan-skus",
		url.Values{
			"plan_id": {planId},
			"is_ext":  {""},
		},
	)
}

func (c *AppClient) SetAuthToken(token string) {
	c.jar.SetCookies(&url.URL{Scheme: "https", Host: "ifdian.net"}, []*http.Cookie{
		{
			Name:    "auth_token",
			Value:   token,
			Path:    "/",
			Domain:  "ifdian.net",
			Expires: time.Now().AddDate(10, 0, 0),
		},
	})
}

func (c *AppClient) DumpCookies() {
	urlInstance, err := url.Parse(AfDianAppUrl)
	if err != nil {
		panic(err)
	}
	cookies := c.jar.Cookies(urlInstance)
	serialized, err := SerializeCookies(cookies)
	if err != nil {
		panic(err)
	}
	println(string(serialized))
}

func NewAppClient() *AppClient {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
		Transport: &TransportWithUA{
			Transport: http.DefaultTransport,
			UA:        "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:135.0) Gecko/20100101 Firefox/135.0",
		},
	}
	return &AppClient{
		jar:    jar,
		client: client,
	}
}

// Cookie 的中间序列化结构
type SerializableCookie struct {
	Name       string `json:"name"`
	Value      string `json:"value"`
	Path       string `json:"path,omitempty"`
	Domain     string `json:"domain,omitempty"`
	Expires    string `json:"expires,omitempty"`
	RawExpires string `json:"raw_expires,omitempty"`
	HttpOnly   bool   `json:"http_only,omitempty"`
	Secure     bool   `json:"secure,omitempty"`
}

// 将 []*http.Cookie 转换为 []SerializableCookie
func SerializeCookies(cookies []*http.Cookie) ([]byte, error) {
	serializable := make([]SerializableCookie, len(cookies))
	for i, cookie := range cookies {
		serializable[i] = SerializableCookie{
			Name:       cookie.Name,
			Value:      cookie.Value,
			Path:       cookie.Path,
			Domain:     cookie.Domain,
			Expires:    cookie.Expires.Format("2006-01-02T15:04:05Z07:00"),
			RawExpires: cookie.RawExpires,
			HttpOnly:   cookie.HttpOnly,
			Secure:     cookie.Secure,
		}
	}
	return json.Marshal(serializable) // 序列化为 JSON
}

// 将序列化后的 JSON 恢复为 []*http.Cookie
func DeserializeCookies(data []byte) ([]*http.Cookie, error) {
	var serializable []SerializableCookie
	err := json.Unmarshal(data, &serializable) // 反序列化 JSON 字符串
	if err != nil {
		return nil, err
	}

	cookies := make([]*http.Cookie, len(serializable))
	for i, sCookie := range serializable {
		cookies[i] = &http.Cookie{
			Name:       sCookie.Name,
			Value:      sCookie.Value,
			Path:       sCookie.Path,
			Domain:     sCookie.Domain,
			RawExpires: sCookie.RawExpires,
			HttpOnly:   sCookie.HttpOnly,
			Secure:     sCookie.Secure,
		}
		// 解析时间字符串为 time.Time 类型
		if sCookie.Expires != "" {
			if parsedTime, err := http.ParseTime(sCookie.Expires); err == nil {
				cookies[i].Expires = parsedTime
			}
		}
	}
	return cookies, nil
}

// 自定义 Transport 设置统一的 User-Agent
type TransportWithUA struct {
	Transport http.RoundTripper
	UA        string
}

// 实现 RoundTripper 接口，添加统一 User-Agent
func (t *TransportWithUA) RoundTrip(req *http.Request) (*http.Response, error) {
	// req.Header.Set("User-Agent", t.UA) // 设置统一的 User-Agent
	// req.Header.Set("Referer", "https://ifdian.net/")
	return t.Transport.RoundTrip(req) // 转发请求
}
