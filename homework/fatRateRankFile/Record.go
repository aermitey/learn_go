package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

	_, err = file.Write(append(data, '\t', '\n'))
	if err != nil {
		log.Println("写入文件失败", err)
		return err
	}
	return nil
}

func (r *Record) UpdateInfo(name string, tall float64, weight float64, age int) (fatRate float64, err error) {
	file, err := os.OpenFile(r.FilePath, os.O_RDWR, 0)
	if err != nil {
		fmt.Println("打开失败", err)
		os.Exit(1) //给01为正常运行
	}
	defer file.Close()
	personInfo := &PersonalInformation{}
	//info := bufio.NewScanner(file)
	//err = os.Truncate(r.FilePath, 0)
	reader := bufio.NewReader(file)
	lineCnt := 0
	seekP := 0
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			log.Println(errors.New("读取完成"))
			break
		}
		lineCnt = len(line) + 1
		json.Unmarshal(line, personInfo)
		if err != nil {
			log.Println("unmarshal失败", err)
			return 0, err
		}
		if personInfo.Name == name {
			personInfo.Weight = weight
			personInfo.Age = age
			personInfo.Tall = tall
			fatRate, err = personInfo.CalcFatRate()
			if err != nil {
				log.Println(err)
				return 0, err
			}
			marshal, err := json.Marshal(personInfo)
			if err != nil {
				log.Println(err)
				return 0, err
			}
			file.WriteAt(append(marshal, '\n'), int64(seekP))
			if err != nil {
				log.Println(err)
				return 0, err
			}
			lineCnt = len(marshal) + 1
		}
		seekP += lineCnt
	}
	return fatRate, nil
}
