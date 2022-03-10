package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "/Users/chenxi/go/src/learngo/chapter08/test"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开失败", err)
		os.Exit(1) //给01为正常运行
	}
	defer file.Close()

	b := make([]byte, 100)
	var n int
	for i := 0; i < 2; i++ {
		n, err = file.Read(b) //将读取到的数据放入b切片
		if err != nil {
			fmt.Println("打开失败", err)
			os.Exit(2)
		}
		//fmt.Println("读出的内容", string(b))     //如果不转成string，输出的是一大堆数字
		fmt.Println("读出的内容", string(b[:n])) //一定要记住给后续程序使用时截取到实际读取的数据，而不是全部，否则后续程序会把无效读取也做正常处理
		fmt.Println("n的大小", n)
		//file.Seek(0, io.SeekStart)    //重新定位到开头
		//file.Seek(1, io.SeekStart)    //重新定位到开头并前进一位
		//file.Seek(0, io.SeekCurrent)  // 返回值ret表示当前游标的位置
		file.Seek(-1, io.SeekCurrent) // 从目前游标的位置开始并给出相对位置
		//file.Seek(1, io.SeekEnd)
		// todo handle error
	}

}
