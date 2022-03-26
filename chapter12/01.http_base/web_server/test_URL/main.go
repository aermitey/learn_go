package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//time.Sleep(2 * time.Second)
		qp := request.URL.Query()
		data, err := json.Marshal(qp)
		if err != nil {
			return
		}
		writer.Write([]byte("hello" + string(data)))
	}))
}
