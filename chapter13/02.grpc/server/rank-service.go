package main

import (
	context2 "golang.org/x/net/context"
	"io"
	"learngo/pkg/apis"
	"log"
	"sync"
)

var _ apis.RankServiceServer = &RankServer{}

type RankServer struct {
	sync.Mutex
	persons map[string]*apis.PersonalInformation
	apis.UnimplementedRankServiceServer
}

type null struct {
}

func (r *RankServer) Register(ctx context2.Context, information *apis.PersonalInformation) (*apis.PersonalInformation, error) {
	r.Lock()
	defer r.Unlock()
	r.persons[information.Name] = information
	log.Printf("收到新注册人：%s\n", information.String())
	return information, nil
}

func (r RankServer) RegisterPersons(server apis.RankService_RegisterPersonsServer) error {
	pis := &apis.PersonalInformationList{}
	for {
		pi, err := server.Recv()
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			log.Printf("WARNING: 获取人员注册时失败:%v\n", err)
			return err
		}
		pis.Items = append(pis.Items, pi)
		r.Lock()
		r.persons[pi.Name] = pi
		r.Unlock()
	}
	log.Println("注册清单", pis.String())
	return server.SendAndClose(pis)
}
