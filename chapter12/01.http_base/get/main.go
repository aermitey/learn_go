package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	reps, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	defer reps.Body.Close()
	data, err := ioutil.ReadAll(reps.Body)
	if err != nil {
		return
	}
	fmt.Println(string(data))
}
