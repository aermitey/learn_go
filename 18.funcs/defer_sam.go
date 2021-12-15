package main

import (
	"fmt"
	"os"
	"time"
)

func deferGuess() {
	startTime := time.Now()
	defer func() {
		fmt.Println("duration:", time.Now().Sub(startTime))
	}()
	fmt.Println("start time:", startTime)
	time.Sleep(5 * time.Second)
	fmt.Println("end time:", time.Now())
}

func openFile() {
	fileName := "/"
	fileObj, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	defer fileObj.Close()
}
