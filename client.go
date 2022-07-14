package afdian

import (
	"bytes"
	"crypto/md5"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"io/ioutil"
	"net/http"
	"time"
)

var json jsoniter.API

func init() {
	extra.RegisterFuzzyDecoders()
	json = jsoniter.ConfigCompatibleWithStandardLibrary
}

const AfDianOpenApiUri = "https://afdian.net/api/open"

type Client struct {
	http.Client
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

// QueryAfdian Go1.18 暂不支持成员函数使用泛型
func QueryAfdian[T any, R any](c *Client, requestPath string, params *T) (*R, error) {
	paramsBuff, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	paramsString := string(paramsBuff)
	ts := time.Now().Unix()
	sign := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%vparams%vts%vuser_id%v", c.Token, paramsString, ts, c.UserId))))
	var requestParams = map[string]any{}
	requestParams["user_id"] = c.UserId
	requestParams["params"] = paramsString
	requestParams["ts"] = ts
	requestParams["sign"] = sign
	buff, err := json.Marshal(requestParams)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%v%v", AfDianOpenApiUri, requestPath), bytes.NewBuffer(buff))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	rsp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	buff, err = ioutil.ReadAll(rsp.Body)
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

// Ping 就很蠢, 非要有参数
func (c *Client) Ping() error {
	_, err := QueryAfdian[struct{ A int }, any](c, "/ping", &struct{ A int }{A: 333})
	return err
}

// QueryOrder 查订单
func (c *Client) QueryOrder(page int64) (*PageData[Order], error) {
	return QueryAfdian[map[string]any, PageData[Order]](c, "/query-order", &map[string]any{
		"page": page,
	})
}

// QuerySponsor 查赞助者
func (c *Client) QuerySponsor(page int64) (*PageData[Sponsor], error) {
	return QueryAfdian[map[string]any, PageData[Sponsor]](c, "/query-sponsor", &map[string]any{
		"page": page,
	})
}

//////////////////// 爱发电调服务器相关 ////////////////////

// ParseOrder 爱发电调服务器的请求内容, 解析成Order
func ParseOrder(body []byte) (*Order, error) {
	var parseBody AfdianCall
	err := json.Unmarshal(body, &parseBody)
	if err != nil {
		return nil, err
	}
	return &parseBody.Data.Order, nil
}

// 应该返回给爱发电一个json

func CallResponseString() string {
	return "{\"ec\":200,\"em\":\"\"}"
}

func CallResponseStruct() *AfdianCallResponse {
	var rsp AfdianCallResponse
	err := json.Unmarshal([]byte(CallResponseString()), &rsp)
	if err != nil {
		panic(err)
	}
	return &rsp
}
