package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	for i := 0; i <= 5000; i++ {
		fmt.Printf("开始数第%d页，", i)
		//time.Sleep(1 * time.Second)
		r, _ := rand.Int(rand.Reader, big.NewInt(800))
		fmt.Printf("有%d个字\n", r.Int64())
		totalCount += r.Int64()
	}
	fmt.Printf("共有%d字", totalCount)
}
