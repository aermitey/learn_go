package main

import "sync"

//type Rank struct {
//	standard []string
//}
//
//var globalRank = &Rank{}
//var globalRankInit bool = false
//var globalRankInitLock sync.Mutex
//
//func init() {
//	globalRank.standard = []string{"Asia"}
//}
//
//func initGlobalRankStandard(standard []string) {
//	globalRankInitLock.Lock()
//	defer globalRankInitLock.Unlock()
//	if globalRankInit {
//		return
//	}
//	globalRank.standard = standard
//	globalRankInit = true
//}
//
//func main() {
//	standard := []string{"asia"}
//	for i := 0; i < 10; i++ {
//		go func() {
//			initGlobalRankStandard(standard)
//		}()
//	}
//}//使用锁类型进行只做一次的操作

type Rank struct {
	standard []string
}

var globalRank = &Rank{}
var once = sync.Once{}

func initGlobalRankStandard(standard []string) {
	once.Do(func() {
		globalRank.standard = standard
	})
}

var facStore = &dbFactoryStore{}

type dbFactoryStore struct {
	store map[string]DBFactory
}
type Conn struct {
}
type DBFactory interface {
	GetConnection() *Conn
}

func initMysqlFac(connStr string) DBFactory {
	return &MySqlDBFactory{}
}

type MySqlDBFactory struct {
	once sync.Once
}

func (f MySqlDBFactory) GetConnection() *Conn {
	once.Do(func() {
		initMysqlFac("")
	})
	//todo
	return nil
}

func main() {
	//standard := []string{"asia"}
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		initGlobalRankStandard(standard)
	//	}()
	//}

}
