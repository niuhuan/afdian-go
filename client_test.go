package afdian_test

import (
	"encoding/json"
	"github.com/niuhuan/afdian-go"
	"io/ioutil"
	"testing"
)

func client() *afdian.Client {
	buff, err := ioutil.ReadFile("afdian.json")
	if err != nil {
		panic(err)
	}
	var client afdian.Client
	err = json.Unmarshal(buff, &client)
	if err != nil {
		panic(err)
	}
	return &client
}

func printResult[T any](t *T, err error) {
	if err != nil {
		panic(err)
	}
	buff, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	println(string(buff))
}

func TestClient_Ping(t *testing.T) {
	client := client()
	err := client.Ping()
	if err != nil {
		panic(err)
	}
}

func TestClient_QueryOrder(t *testing.T) {
	printResult(client().QueryOrder(1))
}

func TestClient_QuerySponsor(t *testing.T) {
	printResult(client().QuerySponsor(1))
}
