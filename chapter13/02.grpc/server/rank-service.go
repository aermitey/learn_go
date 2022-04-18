package main

import (
	"context"
	context2 "golang.org/x/net/context"
	"io"
	"learngo/chapter12/02.practice/frInterface"
	"learngo/pkg/apis"
	"log"
	"sync"
)

var _ apis.RankServiceServer = &RankServer{}

type RankServer struct {
	sync.Mutex
	rankS    frInterface.ServeInterface
	personCh chan *apis.PersonalInformation
	apis.UnimplementedRankServiceServer
}

func (r *RankServer) Update(ctx context.Context, information *apis.PersonalInformation) (*apis.PersonalInformationFatRate, error) {
	r.regPerson(information)
	return r.rankS.UpdatePersonalInformation(information)
}

func (r *RankServer) GetFatRate(ctx context.Context, information *apis.PersonalInformation) (*apis.PersonalRank, error) {
	return r.rankS.GetFatRate(information.Name)
}

func (r *RankServer) GetTop(ctx context.Context, null *apis.Empty) (*apis.PersonalRanks, error) {
	top, err := r.rankS.GetTop()
	if err != nil {
		log.Println("获取榜单时出错:", err)
	}
	return &apis.PersonalRanks{PersonalRanks: top}, nil
}

func (r *RankServer) Register(ctx context2.Context, information *apis.PersonalInformation) (*apis.PersonalInformation, error) {
	r.regPerson(information)
	log.Printf("收到新注册人：%s\n", information.String())
	return information, nil
}

func (r *RankServer) RegisterPersons(server apis.RankService_RegisterPersonsServer) error {
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
		r.regPerson(pi)
	}
	log.Println("注册清单", pis.String())
	return server.SendAndClose(pis)
}

func (r *RankServer) WatchPersons(empty *apis.Empty, server apis.RankService_WatchPersonsServer) error {
	for pi := range r.personCh {
		if err := server.Send(pi); err != nil {
			log.Println("发送失败，结束", err)
			return err
		}
	}
	return nil
}

func (r *RankServer) regPerson(pi *apis.PersonalInformation) {
	err := r.rankS.RegisterPersonalInformation(pi)
	if err != nil {
		log.Println("注册失败:", err)
	}
	r.personCh <- pi
}
