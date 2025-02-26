package afdian_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/niuhuan/afdian-go"
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

func TestClient_QueryOrderByNo(t *testing.T) {
	printResult(client().QueryOrderByNo("202502122353435510010232860"))
}
