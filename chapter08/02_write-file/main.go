package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "/Users/chenxi/go/src/learngo/chapter08/testWrite"
	writeFile(filePath)
	//writeFileWithAppend(filePath)
}

func writeFile(filePath string) {
	file, err := os.Create(filePath) //没有文件的话使用create创建
	if err != nil {
		fmt.Println("打开失败", err)
		os.Exit(1) //给01为正常运行
	}
	defer file.Close()

	_, err = file.Write([]byte("This is first line---"))
	fmt.Println(err)
	//_, err = file.Write([]byte("This is first line\n"))
	//fmt.Println(err)
	//_, err = file.WriteString("第二行内容\n")
	//fmt.Println(err)
	_, err = file.WriteAt([]byte("CHANGED"), 0) //在指定位置写入，替换原位置数据
	file.Sync()                                 //立刻刷盘
	fmt.Println(err)
}

func writeFileWithAppend(filePath string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) //0777最大权限,此方法为追加，无法修改原来的数据
	if err != nil {
		fmt.Println("打开失败", err)
		os.Exit(1) //给01为正常运行
	}
	defer file.Close()

	//_, err = file.Write([]byte("This is first line---"))
	//fmt.Println(err)
	//_, err = file.Write([]byte("This is first line\n"))
	//fmt.Println(err)
	//_, err = file.WriteString("第二行内容\n")
	//fmt.Println(err)
	_, err = file.WriteAt([]byte("123321"), 0) //在指定位置写入，替换原位置数据
	//file.Sync() //立刻刷盘
	fmt.Println(err)
}
