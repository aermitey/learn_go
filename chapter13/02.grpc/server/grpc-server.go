package main

import (
	"context"
	"google.golang.org/grpc"
	"learngo/pkg/apis"
	"log"
	"net"
)

func startGRPCServer(ctx context.Context) {
	lis, err := net.Listen("tcp", nodePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer([]grpc.ServerOption{}...)
	apis.RegisterRankServiceServer(s, &PiServer{
		pis: map[string]*apis.PersonalInformation{},
	})
	go func() {

	}()
}
