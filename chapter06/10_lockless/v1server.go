package main

import (
	"fmt"
	"sync"
	"time"
)

type WebServerV1 struct {
	config     Config
	configLock sync.RWMutex
}

func (v *WebServerV1) reload() {
	v.configLock.Lock()
	defer v.configLock.Unlock()
	v.config.Content = fmt.Sprintf("%d", time.Now().UnixNano())
	//v.config.Content = "1"
}

func (v *WebServerV1) ReloadWorker() {
	for {
		time.Sleep(10 * time.Millisecond)
		v.reload()
	}
}

func (v *WebServerV1) Visit() string {
	v.configLock.RLock()
	defer v.configLock.RUnlock()
	return v.config.Content
}
