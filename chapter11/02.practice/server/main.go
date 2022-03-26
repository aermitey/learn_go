package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"learngo/chapter02/15.fatrate_refactor/calc"
	"learngo/pkg/apis"
	"log"
	"net"
	"time"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "配置启动端口")
	flag.Parse()

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	rank := NewFatRateRank()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("wanring: 建立连接失败：", err)
			continue
		}
		fmt.Println(conn)

		// talk(conn)
		go talk(conn, rank)
	}
}

func talk(conn net.Conn, rank *FatRateRank) {
	defer fmt.Println("结束链接：", conn)
	defer conn.Close()
	for {
		//读取到完整的数据
		finalReceivedMessage := make([]byte, 0, 1024)
		encounterError := false
		for {
			buf := make([]byte, 1024)
			valid, err := conn.Read(buf)
			if err != nil {
				log.Println("WARNING：读数据时出问题，重新读取", err)
				encounterError = true
				time.Sleep(1 * time.Second)
				break
			}
			if valid == 0 {
				break
			}
			validContent := buf[:valid]
			finalReceivedMessage = append(finalReceivedMessage, validContent...)
			if valid < len(buf) {
				break
			}
		}
		if encounterError {
			_, err := conn.Write([]byte(`服务器获取数据失败，请重新输入`))
			if err != nil {
				return
			}
			continue
		}

		personInfo := &apis.PersonalInformation{}
		err := json.Unmarshal(finalReceivedMessage, personInfo)
		if err != nil {
			_, err := conn.Write([]byte(`输入的数据无法解析，请重新输入`))
			if err != nil {
				return
			}
			continue
		}
		bmi, err := calc.CalcBMI(float64(personInfo.Tall), float64(personInfo.Weight))
		if err != nil {
			_, err := conn.Write([]byte(`无法计算bmi，请重新输入`))
			if err != nil {
				return
			}
			continue
		}
		fatRate, err := calc.FatRateFromBMI(bmi, int(personInfo.Age), personInfo.Sex)
		if err != nil {
			_, err := conn.Write([]byte(`无法计算体脂，请重新输入`))
			if err != nil {
				return
			}
			continue
		}
		rank.inputRecord(personInfo.Name, fatRate)
		rankId, _ := rank.getRank(personInfo.Name)

		_, err = conn.Write([]byte(fmt.Sprintf("姓名: %s, BMI: %f, 体脂率: %f, 排名: %d", personInfo.Name, bmi, fatRate, rankId)))
		if err != nil {
			return
		}
		break
	}
}
