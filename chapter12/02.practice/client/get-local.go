package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type frClient struct {
	handRing ClientInterface
}

func (f frClient) register() {
	pi, _ := f.handRing.ReadPersonalInformation()
	data, err := json.Marshal(pi)
	if err != nil {
		return
	}
	r := bytes.NewReader(data)
	resp, err := http.Post("http://localhost:8080/Register", "application/json", r)
	if err != nil {
		log.Println("注册失败", err)
		return
	}
	if resp.Body != nil {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		fmt.Println("返回:", string(data))
	}
}

func (f frClient) update() {
	pi, _ := f.handRing.ReadPersonalInformation()
	data, err := json.Marshal(pi)
	if err != nil {
		return
	}
	r := bytes.NewReader(data)
	resp, err := http.Post("http://localhost:8080/personalInfo", "application/json", r)
	if err != nil {
		log.Println("更新失败", err)
		return
	}
	if resp.Body != nil {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		fmt.Println("返回:", string(data))
	}
}

func main() {
	frCli := &frClient{handRing: &fakeInterFace{
		name:       fmt.Sprintf("小强%d", time.Now().UnixNano()),
		sex:        "男",
		baseWeight: 71.0,
		baseTall:   1.7,
		baseAge:    35,
	}}
	frCli.register()
	for {
		frCli.update()
		time.Sleep(1 * time.Second)
	}
}
