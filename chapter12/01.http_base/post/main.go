package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	r := strings.NewReader("fool")
	response, err := http.Post("http://www.baidu.com", "*/*", r)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println(string(data))
}
