package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second) //在此等待五秒
	defer cancel()
	go httpDirectGet()
	go httpGetWithContext(ctx)
	<-ctx.Done()
}

func httpDirectGet() {
	reps, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer reps.Body.Close()
	data, err := ioutil.ReadAll(reps.Body)
	if err != nil {
		return
	}
	fmt.Println(string(data))
}

//带有上下文控制的http操作
func httpGetWithContext(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	req, err := http.NewRequest("get", "http://localhost:8080", nil)
	if err != nil {
		log.Fatal("无法生成请求", err)
	}
	//此处要赋值给新的req，因为原本的req经过WithContext之后未发生任何变化
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("无法发送请求", err)
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("无法读取返回内容", err)
	}
	fmt.Println(string(data))
}
