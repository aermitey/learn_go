package main

import (
	"fmt"
	"time"
)

type DownLoader struct {
	fileNameCh chan string
	finishCh   chan string
}

func main() {
	fileNameCh := make(chan string)
	finishCh := make(chan string, 100) // 如果这个finished chan 没有buffer，则在download完成后无法将结果放进该channel，从而导致死锁。没有buffer的channel必须有消费者才可以向channel内装数据。
	workerCount := 50
	for i := 0; i < workerCount; i++ {
		//启动协程准备下载
		go func() {
			downLoader := &DownLoader{
				fileNameCh: fileNameCh,
				finishCh:   finishCh,
			}
			downLoader.Download()
		}()
	}

	fileNumber := 100
	fileNames := make([]string, 0, fileNumber)
	for i := 0; i < fileNumber; i++ {
		fileNames = append(fileNames, fmt.Sprintf("file_%d.txt", i))
	}

	for _, fileItem := range fileNames {
		fileNameCh <- fileItem
	}
	close(fileNameCh)

	finishedFileCount := 0
	go func() {
		for finishedFile := range finishCh {
			fmt.Println("文件:", finishedFile, "处理完毕")
			finishedFileCount++
			if finishedFileCount == fileNumber {
				close(finishCh)
			}
		}
	}()

	fmt.Println("所有文件处理完毕，结束")
}

func (d *DownLoader) Download() {
	for fileName := range d.fileNameCh {
		fmt.Println("开始下载文件:", fileName)
		time.Sleep(1 * time.Second)
		fmt.Println("开始处理文件:", fileName)
		time.Sleep(10 * time.Millisecond)
		fmt.Println("保存文件:", fileName)
		d.finishCh <- fileName
	}
}
