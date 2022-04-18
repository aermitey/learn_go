package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"learngo/pkg/apis"
	"log"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not conect: %v", err)
	}
	c := apis.NewRankServiceClient(conn)
	w, err := c.WatchPersons(context.TODO(), &apis.Empty{})
	if err != nil {
		log.Fatal("启动watcher失败", err)
	}
	for {
		pi, err := w.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("服务器告知返回数据完成")
				break
			}
			log.Fatal("接收异常", err)
		}
		log.Println("收到新变动", pi.String())
	}
}
