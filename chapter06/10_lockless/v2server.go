package main

import (
	"fmt"
	"time"
)

type WebServerV2 struct {
	config *Config
}

func (v *WebServerV2) reload() {
	//v.config = &Config{
	//	fmt.Sprintf("%d", time.Now().UnixNano()),
	//}
	//v.config.Content = "1"
	v.config.Content = fmt.Sprintf("%d", 1)
}

func (v *WebServerV2) ReloadWorker() {
	for {
		time.Sleep(10 * time.Millisecond)
		v.reload()
	}
}

func (v *WebServerV2) Visit() string {
	return v.config.Content
}
