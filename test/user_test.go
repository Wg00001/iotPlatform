package test

import (
	"encoding/json"
	"fmt"
	"iotPlatform/common"
	"iotPlatform/helper"
	"testing"
)

var userServiceAddr = "http://127.0.0.1:14000"

func TestUserLogin(t *testing.T) {
	m := common.M{
		"username": "get",
		"password": "123456",
	}
	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAddr+"/user/login", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}
