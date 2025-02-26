package afdian_test

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/niuhuan/afdian-go"
)

func login(client *afdian.AppClient, t *testing.T) {
	data, err := ioutil.ReadFile("account.json")
	if err != nil {
		t.Fatal(err)
	}
	var accountJsonStruct struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		AuthToken string `json:"auth_token"`
	}
	err = json.Unmarshal(data, &accountJsonStruct)
	if err != nil {
		t.Fatal(err)
	}
	// login
	// uset and password
	// rsp, err := client.Login(accountJsonStruct.Username, accountJsonStruct.Password)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// set token
	client.SetAuthToken(accountJsonStruct.AuthToken)
	// login
	println("====================================")
	client.DumpCookies()
	println("====================================")
}

func TestAppClient_DumpCookies(t *testing.T) {
	client := afdian.NewAppClient()
	login(client, t)
	ma, err := client.MyAccount()
	if err != nil {
		t.Fatal(err)
	}
	txt, _ := json.Marshal(ma)
	println(string(txt))
}

func TestAppClient_Plans(t *testing.T) {
	data, err := ioutil.ReadFile("account_id.txt")
	if err != nil {
		t.Fatal(err)
	}
	accountId := string(data)
	accountId = strings.TrimSpace(accountId)
	client := afdian.NewAppClient()
	login(client, t)
	plains, err := client.Plans(accountId)
	if err != nil {
		t.Fatal(err)
	}
	buff, err := json.Marshal(plains)
	if err != nil {
		t.Fatal(err)
	}
	println(string(buff))
}

func TestAppClient_PlanSkus(t *testing.T) {
	data, err := ioutil.ReadFile("account_plan_id.txt")
	if err != nil {
		t.Fatal(err)
	}
	planId := string(data)
	planId = strings.TrimSpace(planId)
	client := afdian.NewAppClient()
	login(client, t)
	plains, err := client.PlanSkus(planId)
	if err != nil {
		t.Fatal(err)
	}
	buff, err := json.Marshal(plains)
	if err != nil {
		t.Fatal(err)
	}
	println(string(buff))
}
