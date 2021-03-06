package main

import (
	"context"
	"google.golang.org/grpc"
	"learngo/pkg/apis"
	"log"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not conect: %v", err)
	}
	c := apis.NewRankServiceClient(conn)
	ret, err := c.Register(context.TODO(), &apis.PersonalInformation{
		Name:   "tom",
		Sex:    "男",
		Tall:   1.7,
		Weight: 70,
		Age:    30,
	})
	if err != nil {
		log.Fatal("注册失败")
	}
	log.Println("注册成功", ret)

	top, err := c.GetTop(context.TODO(), &apis.Empty{})
	if err != nil {
		log.Fatal("获取榜单失败:", err)
	}
	log.Println("榜单:", top.String())

	regCli, err := c.RegisterPersons(context.TODO())
	if err != nil {
		log.Fatal("获取批量注册客户端失败:", err)
	}
	for i := 0; i < 5; i++ {
		if err := regCli.Send(&apis.PersonalInformation{
			Name:   "tom",
			Sex:    "男",
			Tall:   1.7,
			Weight: 70,
			Age:    30,
		}); err != nil {
			log.Fatal("注册时失败:", err)
		}
	}
	resp, err := regCli.CloseAndRecv()
	if err != nil {
		log.Fatal("无法接受结果:", err)
	}
	log.Println("批量注册结果:", resp.String())
}
