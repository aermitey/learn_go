package main

import (
	"crypto/rand"
	"learngo/pkg/apis"
	"math/big"
)

type ClientInterface interface {
	ReadPersonalInformation() (apis.PersonalInformation, error)
}

var _ ClientInterface = &fakeInterFace{}

type fakeInterFace struct {
	name       string
	sex        string
	baseWeight float64
	baseTall   float64
	baseAge    int
}

func (f *fakeInterFace) ReadPersonalInformation() (apis.PersonalInformation, error) {
	r, _ := rand.Int(rand.Reader, big.NewInt(1000))
	out := float64(r.Int64()) / 1000
	if r.Int64()%2 == 0 {
		out = 0 - out
	}
	personInfo := apis.PersonalInformation{
		Name:   f.name,
		Sex:    f.sex,
		Tall:   float32(f.baseTall),
		Weight: float32(f.baseWeight),
		Age:    int64(f.baseAge),
	}
	f.baseWeight += out
	return personInfo, nil
}
