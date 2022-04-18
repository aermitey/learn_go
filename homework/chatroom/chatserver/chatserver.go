package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"learngo/homework/chatroom/apis/proto"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var _ proto.ChatServiceServer = &ChatServer{}

//var _ proto.ChatCoupleServer = &ChatCouple{}

type ChatServer struct {
	conn      *gorm.DB
	watchChan chan *proto.ChatLog
	proto.UnimplementedChatServiceServer
	ChatCouple
}

type ChatCouple struct {
	conn           *gorm.DB
	sourceAccount int64
	targetAccount int64
	sourceToTarget chan *proto.ChatLog
	targetToSource chan *proto.ChatLog
	proto.UnimplementedChatCoupleServer
}

func (c *ChatCouple) MakeChatRoom(ctx context.Context, num *proto.Num) (*ChatCouple, error) {
	value := ctx.Value("SourceAccount")
	if user, ok := value.(proto.User); ok {
		return &ChatCouple{
			conn:                          c.conn,
			sourceAccount:                 user.Account,
			targetAccount:                 num.Num,
			sourceToTarget:                make(chan *proto.ChatLog),
			targetToSource:                make(chan *proto.ChatLog),
		}, nil
	}
	err := fmt.Errorf("未获取正确的参数")
	return nil, err
}

func (c *ChatCouple) SendChatContent(ctx context.Context, ch chan *proto.ChatLog, content *proto.Content) (*proto.Empty, error) {
	output := proto.User{}
	value := ctx.Value("SourceAccount")
	if user, ok := value.(proto.User); ok {
		resp := c.conn.Where(&proto.Account{Account: content.Account}).Find(&output)
		if err := resp.Error; err != nil {
			log.Println("账号输入错误", err)
			return nil, err
		}
		chatLog := proto.ChatLog{}
		chatLog.TalkerAccount = user.Account
		chatLog.TalkerNickname = user.Nickname
		chatLog.ListenerAccount = content.Account
		chatLog.ListenerNickname = output.Nickname
		chatLog.ChatContent = content.Content
		chatLog.ChatTime = fmt.Sprintf("%s", time.Now())
		ch <- &chatLog
		return nil, nil
	}
	err := fmt.Errorf("未获取正确的参数")
	return nil, err
}

func (c *ChatCouple) ReceiveChatContent(ctx context.Context, ch chan *proto.ChatLog) (*proto.ChatLog, error) {
	return <-ch, nil
}

func (c *ChatServer) Register(ctx context.Context, account *proto.Account) (*proto.User, error) {
	randNum := GenValidateCode()
	account.Account = randNum
	data, err := json.Marshal(account)
	if err != nil {
		log.Println("marshal出错:", err)
		return nil, err
	}
	file, err := os.Create("/Users/chenxi/go/src/learngo/homework/userInfo.json")
	if err != nil {
		log.Println("写入失败", err)
		return nil, err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return nil, err
	}
	resp := c.conn.Create(&proto.Account{
		Account:  randNum,
		Password: account.Password,
		Nickname: account.Nickname,
	})
	if err := resp.Error; err != nil {
		fmt.Println("创建失败", err)
		return nil, err
	}
	return &proto.User{Account: randNum}, nil
}

func (c *ChatServer) Login(ctx context.Context, account *proto.Account) (*proto.User, error) {
	output := proto.Account{}
	user := proto.User{}
	if account.Password == "" {
		file, err := os.Open("/Users/chenxi/go/src/learngo/homework/userInfo")
		if err != nil {
			log.Println("读取密码失败", err)
			return nil, err
		}
		defer file.Close()
		b := make([]byte, 1000)
		_, err = file.Read(b)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(b, account)
		if err != nil {
			return nil, err
		}
	}
	user.Account = account.Account
	user.Nickname = account.Nickname
	ctx = context.WithValue(ctx, "1", user)
	resp := c.conn.Where(&proto.Account{Account: account.Account}).Find(&output)
	if err := resp.Error; err != nil {
		log.Println("账号输入错误", err)
		return nil, err
	}
	if output.Password != account.Password {
		err := fmt.Errorf("密码输入错误")
		return nil, err
	} else {
		return &user, nil
	}
}

func (c *ChatServer) ChatList(ctx context.Context, num *proto.Num) (*proto.UserList, error) {
	output := proto.UserList{}
	resp := c.conn.Where(&proto.Account{Online: "在线"}).Limit(20).Offset(int(num.Num-1) * 20).Find(&output)
	if err := resp.Error; err != nil {
		log.Println("查询在线用户失败")
		return nil, err
	}
	return &output, nil
}

func (c *ChatServer) ChatHistory(ctx context.Context, num *proto.Num) (*proto.ChatLogList, error) {
	output := proto.ChatLogList{}
	value := ctx.Value("SourceAccount")
	if user, ok := value.(proto.User); ok {
		resp := c.conn.Where("talkerAccount = ? and listenerAccount = ?", user.Account, num).Order("chatTime desc").Limit(20).Find(&output)
		if err := resp.Error; err != nil {
			log.Println("查询聊天记录失败")
			return nil, err
		}
		return &output, nil
	}
	err := fmt.Errorf("未获取正确的参数")
	return nil, err
}

func (c *ChatServer) ChatWith(ctx context.Context, num *proto.Num, content *proto.Content) (*proto.Empty, error) {
	_, err := c.ChatCouple.SendChatContent(ctx, c.ChatCouple.sourceToTarget, content)
	if err != nil {
		return nil, err
	}
	chatContent, err := c.ChatCouple.ReceiveChatContent(ctx, c.ChatCouple.targetToSource)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%s", chatContent)
	return nil, nil
}

func (c *ChatServer) ReceiveSend(ctx context.Context, num *proto.Num, content *proto.Content) (*proto.Empty, error) {
	_, err := c.ChatCouple.SendChatContent(ctx, c.ChatCouple.targetToSource, content)
	if err != nil {
		return nil, err
	}
	chatContent, err := c.ChatCouple.ReceiveChatContent(ctx, c.ChatCouple.sourceToTarget)
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("%s", chatContent)
	return nil, nil
}


func (c *ChatServer) WatchChatRoom(ctx context.Context, empty *proto.Empty) (*proto.ChatLog, error) {
	value := ctx.Value("SourceAccount")
	if user, ok := value.(proto.User); ok {
		for chatLog := range c.watchChan {
			if user.Account == chatLog.ListenerAccount {
				return chatLog, nil
			}
		}
	}
	err := fmt.Errorf("未获取正确的参数")
	return nil, err
}

func GenValidateCode() int64 {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < 6; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	num, _ := strconv.Atoi(sb.String())
	return int64(num)
}
