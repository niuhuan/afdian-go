package afdian_test

import (
	"encoding/json"
	"io/ioutil"
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
