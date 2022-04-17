package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"learngo/homework/chatroom/apis/proto"
	"log"
	"os"
)

func main() {
	chatRegisterCmd := flag.NewFlagSet("chat register", flag.ExitOnError)
	chatRegisterNickname := chatRegisterCmd.String("nickname", "", "nickname required for registration")
	chatRegisterPassword := chatRegisterCmd.String("password", "", "password required for registration")

	chatLoginCmd := flag.NewFlagSet("chat login", flag.ExitOnError)
	chatLoginAccount := chatLoginCmd.Int("account", 0, "password required for login")
	chatLoginPassword := chatLoginCmd.String("password", "", "password required for registration")

	chatListCmd := flag.NewFlagSet("chat list", flag.ExitOnError)
	chatListPage := chatListCmd.Int("page", 0, "Query the exact number of pages in the chat log")

	chatWithCmd := flag.NewFlagSet("chat with", flag.ExitOnError)
	chatWithAccount := chatWithCmd.Int("account", 0, "chat with somebody")

	chatHistoryCmd := flag.NewFlagSet("chat history", flag.ExitOnError)
	chatHistoryWith := chatHistoryCmd.Int("account", 0, "Check the history of chatting with an account")

	chatSubscribeCmd := flag.NewFlagSet("chat history", flag.ExitOnError)

	conn, err := grpc.Dial("0.0.0.0:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not conect: %v", err)
	}
	c := proto.NewChatServiceClient(conn)

	account := &proto.Account{}
	user := &proto.User{}
	switch os.Args[1] {
	case "chat register":
		err := chatRegisterCmd.Parse(os.Args[3:])
		if err != nil {
			fmt.Println("参数错误")
			return
		}
		account.Nickname = *chatRegisterNickname
		account.Password = *chatRegisterPassword
		userInfo, err := c.Register(context.TODO(), account)
		if err != nil {
			log.Println("注册失败")
			return
		}
		user = userInfo
		fmt.Printf("%d", user.Account)
	case "chat login":
		chatLoginCmd.Parse(os.Args[2:])
		account.Account = int64(*chatLoginAccount)
		account.Password = *chatLoginPassword
		userInfo, err := c.Login(context.TODO(), account)
		if err != nil {
			fmt.Println("登录失败")
			return
		}
		user = userInfo
	case "chat list":
		chatListCmd.Parse(os.Args[2:])
		num := *chatListPage
		userList, err := c.ChatList(context.TODO(), &proto.Num{Num: int64(num)})
		if err != nil {
			return
		}
		fmt.Printf("在线用户列表第%d页%v", num, userList)
	case "chat with":
		chatWithCmd.Parse(os.Args[2:])
		ctx := context.WithValue(context.TODO(), "SourceAccount", user)
		go func() {
			c
		}()
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
