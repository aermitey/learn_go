package main

import (
	context2 "golang.org/x/net/context"
	"learngo/pkg/apis"
)

var _ apis.RankServiceServer = &RankServer{}

type RankServer struct {
}

func (r RankServer) Register(ctx context2.Context, information *apis.PersonalInformation) (*apis.PersonalInformation, error) {
	//TODO implement me
	panic("implement me")
}

func (r RankServer) mustEmbedUnimplementedRankServiceServer() {
	//TODO implement me
	panic("implement me")
}
