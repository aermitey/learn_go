package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"learngo/chapter12/02.practice/frInterface"
	"learngo/pkg/apis"
	"net/http"
	"strings"
)

func main() {
	var rankServer frInterface.ServeInterface = NewFatRateRank()
	m := http.NewServeMux()
	m.Handle("/Register", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !strings.EqualFold(request.Method, "post") { // strings.EqualFold大小写不敏感
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		if request.Body == nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		defer request.Body.Close()
		payLoad, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("无法读取数据:%s", err)))
			return
		}
		var pi *apis.PersonalInformation
		err = json.Unmarshal(payLoad, &pi)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("无法解析数据:%s", err)))
			return
		}
		if err := rankServer.RegisterPersonalInformation(pi); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf("注册失败:%s", err)))
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(`success`))
	}))

	m.Handle("/personalInfo", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !strings.EqualFold(request.Method, "post") { // strings.EqualFold大小写不敏感
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		if request.Body == nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		defer request.Body.Close()
		payLoad, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("无法读取数据:%s", err)))
			return
		}
		var pi *apis.PersonalInformation
		err = json.Unmarshal(payLoad, &pi)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("无法解析数据:%s", err)))
			return
		}
		if fr, err := rankServer.UpdatePersonalInformation(pi); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf("更新失败:%s", err)))
			return
		} else {
			writer.WriteHeader(http.StatusOK)
			data, err := json.Marshal(fr)
			if err != nil {
				return
			}
			writer.Write(data)
		}
	}))

	m.Handle("/rank", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !strings.EqualFold(request.Method, "get") { // strings.EqualFold大小写不敏感
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		name := request.URL.Query().Get("name")
		if name == "" {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("name参数未设置"))
			return
		}
		if fr, err := rankServer.GetFatRate(name); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf("获取排行失败:%s", err)))
			return
		} else {
			writer.WriteHeader(http.StatusOK)
			data, err := json.Marshal(fr)
			if err != nil {
				return
			}
			writer.Write(data)
		}
	}))

	m.Handle("/ranktop", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !strings.EqualFold(request.Method, "get") { // strings.EqualFold大小写不敏感
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		if topFrs, err := rankServer.GetTop(); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf("获取排行失败:%s", err)))
			return
		} else {
			writer.WriteHeader(http.StatusOK)
			data, err := json.Marshal(topFrs)
			if err != nil {
				return
			}
			writer.Write(data)
		}
	}))
	http.ListenAndServe(":8080", m)
}
