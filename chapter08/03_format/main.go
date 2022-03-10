package main

import (
	"encoding/json"
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
	"io/ioutil"
	"learngo/chapter02/15.fatrate_refactor/calc"
	"learngo/pkg/apis/proto"
	"log"
	"os"
)

func main() {
	input := &inputFromStd{}
	record := newRecord("/Users/chenxi/go/src/learngo/chapter08/testFatRate")
	calc := &Calc{}
	fatRateRank := &FatRateRank{}

	for {
		personInfo := input.GetInPut()

		err := record.savePersonalInformation(personInfo)
		if err != nil {
			log.Fatal("保存失败", err)
		}

		fatRate, err := calc.FatRate(personInfo)
		if err != nil {
			log.Fatal("计算体脂率失败:", err)
		}

		fatRateRank.inputRecord(personInfo.Name, fatRate)
		rankResult, _ := fatRateRank.getRank(personInfo.Name)
		fmt.Println("排名结果:", rankResult)
	}
}

func caseStudy() {
	filePath := "/Users/chenxi/go/src/learngo/chapter08/testFatRate.json"

	personalInformation := proto.PersonalInformation{
		Name:   "小强...'",
		Sex:    "男",
		Tall:   1.70,
		Weight: 71,
		Age:    35,
	}

	fmt.Printf("%v\n", personalInformation)
	data, err := json.Marshal(personalInformation)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("marshal 的结果是(原生)：", data)
	fmt.Println("marshal 的结果是(string)：", string(data))

	writeFile(filePath, data)
	readFile(filePath)
}

func readFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}
	personalInformation := proto.PersonalInformation{}
	err = json.Unmarshal(data, &personalInformation)
	if err != nil {
		log.Fatal("错误：", err)
		return
	}
	fmt.Println("开始计算体脂信息:", personalInformation)
	bmi, _ := gobmi.BMI(float64(personalInformation.Weight), float64(personalInformation.Tall))
	fmt.Printf("%s的BMI是：%v\n", personalInformation.Name, bmi)
	fatRate, _ := calc.FatRateFromBMI(bmi, int(personalInformation.Age), personalInformation.Sex)
	fmt.Printf("%s的体脂率是：%v", personalInformation.Name, fatRate)
}

func writeFile(filePath string, data []byte) {
	file, err := os.Create(filePath) //没有文件的话使用create创建
	if err != nil {
		fmt.Println("打开失败", err)
		os.Exit(1) //给01为正常运行
	}
	defer file.Close()

	_, err = file.Write(data)
	fmt.Println(err)
}
