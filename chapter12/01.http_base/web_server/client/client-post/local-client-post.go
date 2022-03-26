package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	r := strings.NewReader("fooooooooo")
	response, err := http.Post("http://localhost:8080", "*/*", r)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println("reps:", string(data))
}
