package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Record struct {
	FilePath string
}

func (r *Record) SaveInformation(personInfo *PersonalInformation) error {
	data, err := json.Marshal(personInfo)
	if err != nil {
		log.Println("Marshal出错", err)
		return err
	}
	err = r.WriteFileWithAppend(data)
	if err != nil {
		log.Println("写入文件失败", err)
		return err
	}
	return nil
}

func (r *Record) WriteFileWithAppend(data []byte) error {
	file, err := os.OpenFile(r.FilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal("打开文件失败:", err)
	}
	defer file.Close()

	_, err = file.Write(append(data, '\n'))
	if err != nil {
		log.Println("写入文件失败", err)
		return err
	}
	return nil
}

func (r *Record) UpdateInfo(name string, tall float64, weight float64, age int) (err error) {
	personInfo, err := r.Unmarshal(err)
	if err != nil {
		return err
	}
	for k, _ := range *personInfo {
		if (*personInfo)[k].Name == name {
			(*personInfo)[k].Tall = tall
			(*personInfo)[k].Weight = weight
			(*personInfo)[k].Age = age
		}
	}

	dataJson, err := json.Marshal(personInfo)
	if err != nil {
		log.Println("Marshal出错", err)
		return err
	}
	r.writeFile(dataJson)
	//fmt.Printf("值为:%+v\n", personInfo)
	//fmt.Println("unmarshal 的结果是(string)：", personInfo)
	return
}

func (r *Record) Unmarshal(err error) (*[]PersonalInformation, error) {
	n, b := r.readFile()
	personInfo := &[]PersonalInformation{}
	data := string(b[:n])
	//data = []string{`{"name":"陈曦","sex":"男","tall":1.65,"weight":80,"age":27}`, `{"name":"小强...'","sex":"男","tall":1.7,"weight":71,"age":35}`}
	//data = []string{string(b[:n])}
	//fmt.Println(data)
	err = json.Unmarshal([]byte(data), personInfo)
	if err != nil {
		return nil, err
	}
	return personInfo, nil
}

func (r *Record) readFile() (int, []byte) {
	file, err := os.Open(r.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b := make([]byte, 1000)
	n, err := file.Read(b)
	//ioutil.ReadFile()
	//ioutil.ReadAll()

	//fmt.Println("切片内容", string(b[:n]))
	//for i, data := range string(b[:n]) {
	//	fmt.Println(i, data)
	//}
	if err != nil {
		log.Fatal(err)
	}
	return n, b
}

func (r *Record) writeFile(data []byte) error {
	file, err := os.Create(r.FilePath) //没有文件的话使用create创建
	if err != nil {
		fmt.Println("打开失败", err)
		os.Exit(1) //给01为正常运行
	}
	defer file.Close()

	_, err = file.Write(append(data, '\n'))
	if err != nil {
		log.Println("写入文件失败", err)
		return err
	}
	return nil
}
