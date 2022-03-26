package main

import (
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	proto2 "learngo/pkg/apis"
	"log"
	"time"
)

func main() {
	counter := 5000
	persons := make([]*proto2.PersonalInformation, 0, counter)
	for i := 0; i < counter; i++ {
		persons = append(persons, &proto2.PersonalInformation{
			Name:   "小强",
			Sex:    "男",
			Tall:   1.7,
			Weight: 71,
			Age:    35,
		})
	}
	{
		fmt.Println("序列化json")
		startTime := time.Now()
		data, err := json.Marshal(persons)
		if err != nil {
			log.Fatal(err)
		}
		finishMarshalTime := time.Now()
		err = ioutil.WriteFile("/Users/chenxi/Desktop/data.json", data, 0777)
		if err != nil {
			return
		}
		finishWriteTime := time.Now()
		fmt.Println("序列化耗时:", finishMarshalTime.Sub(startTime))
		fmt.Println("写文件耗时:", finishWriteTime.Sub(finishMarshalTime))
	}
	{
		fmt.Println("反序列化json")
		startTime := time.Now()
		data, err := ioutil.ReadFile("/Users/chenxi/Desktop/data.json")
		if err != nil {
			return
		}
		err = json.Unmarshal(data, &persons)
		if err != nil {
			return
		}
		finishUnMarshalTime := time.Now()
		fmt.Println("反序列化json耗时:", finishUnMarshalTime.Sub(startTime))
	}
	{
		fmt.Println("序列化yaml")
		startTime := time.Now()
		data, err := yaml.Marshal(persons)
		if err != nil {
			log.Fatal(err)
		}
		finishMarshalTime := time.Now()
		err = ioutil.WriteFile("/Users/chenxi/Desktop/data.yaml", data, 0777)
		if err != nil {
			return
		}
		finishWriteTime := time.Now()
		fmt.Println("序列化耗时:", finishMarshalTime.Sub(startTime))
		fmt.Println("写文件耗时:", finishWriteTime.Sub(finishMarshalTime))
	}
	{
		fmt.Println("序列化proto")
		Lister := &proto2.PersonalInformationList{Items: persons}
		startTime := time.Now()
		data, err := proto.Marshal(Lister)
		if err != nil {
			log.Fatal(err)
		}
		finishMarshalTime := time.Now()
		err = ioutil.WriteFile("/Users/chenxi/Desktop/data.proto", data, 0777)
		if err != nil {
			return
		}
		finishWriteTime := time.Now()
		fmt.Println("序列化耗时:", finishMarshalTime.Sub(startTime))
		fmt.Println("写文件耗时:", finishWriteTime.Sub(finishMarshalTime))
	}
	{
		fmt.Println("反序列化proto")
		Lister := &proto2.PersonalInformationList{}
		startTime := time.Now()
		data, err := ioutil.ReadFile("/Users/chenxi/Desktop/data.proto")
		if err != nil {
			return
		}
		err = proto.Unmarshal(data, Lister)
		if err != nil {
			return
		}
		finishUnMarshalTime := time.Now()
		fmt.Println("反序列化proto耗时:", finishUnMarshalTime.Sub(startTime))
	}
}
