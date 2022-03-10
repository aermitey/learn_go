package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type Rank interface {
	UpdateFR(name string, fr float64)
	GetRank(name string) int
}

type Client interface {
	UpdateFR(name string, fr float64)
	GetRank(name string) int
}

func randFloat() float64 {
	rand.Seed(time.Now().UnixNano())
	num := rand.Float64()
	for {
		if num > 0.4 {
			num = num - 0.4
		} else {
			return num
		}
	}
	return num
}

func TestRank(t *testing.T) {
	//var rank Rank        //todo 实例化
	var clients []Client //todo 实例化

	for i := 0; i < len(clients); i++ {
		go func(i int) {
			//todo add context control exit
			go func() {
				for {
					clients[i].UpdateFR(fmt.Sprintf("%d", i), randFloat())
				}
			}()
			go func() {
				for  {
					clients[i].GetRank(fmt.Sprintf("%d", i))
				}
			}()
		}(i)
	}
}
