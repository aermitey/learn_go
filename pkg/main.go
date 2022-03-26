package main

import (
	"encoding/json"
	"fmt"
	"learngo/pkg/apis"
)

func main() {
	fc := &apis.FriendCircle{
		Id:         1,
		PersonId:   1,
		Name:       "陈曦",
		Sex:        "男",
		Tall:       1.7,
		Weight:     70,
		Age:        25,
		FatRate:    0.2,
		Content:    "123456",
		CreateTime: "123456",
		IsDeleted:  0,
	}
	data, err := json.Marshal(fc)
	if err != nil {
		return
	}
	fmt.Println(string(data))
}
