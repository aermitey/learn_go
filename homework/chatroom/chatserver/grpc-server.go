package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learngo/homework/chatroom/apis/proto"
	"log"
	"net"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	startGRPCServer(ctx)
}

func startGRPCServer(ctx context.Context) {
	conn := connectDB()
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer([]grpc.ServerOption{}...)
	proto.RegisterChatServiceServer(s, &ChatServer{
		conn: conn,
	})
	go func() {
		select {
		case <-ctx.Done():
			s.Stop()
		}
	}()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func connectDB() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("chenxi:123456@tcp(127.0.0.1:3306)/test"))
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	fmt.Println("连接数据库成功")
	return conn
}
