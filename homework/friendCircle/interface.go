package main

import (
	"learngo/pkg/apis"
)

type CircleInterface interface {
	PostingStatus(fc *apis.FriendCircle) error
	DeleteStatus(id int) error
	Browse() ([]*apis.FriendCircle, error)
}

type RankInterface interface {
	RegisterPersonalInformation(pi *apis.PersonalInformation) error
	UpdatePersonalInformation(pi *apis.PersonalInformation) (*apis.PersonalInformationFatRate, error)
	GetFatRate(name string) (*apis.PersonalRank, error)
	GetTop() ([]*apis.PersonalRank, error)
}

type RankInitInterface interface {
	Init() error
}
