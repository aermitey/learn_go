package main

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learngo/chapter12/02.practice/frInterface"
	"learngo/pkg/apis"
	"log"
	"net/http"
)

func main() {
	conn := connectDB()
	var rankServer frInterface.ServeInterface = NewDBRank(conn, NewFatRateRank())
	if initRank, ok := rankServer.(frInterface.RankInitInterface); ok {
		err := initRank.Init()
		if err != nil {
			log.Fatal("初始化失败", err)
		}
	}
	r := gin.Default()
	pprof.Register(r)
	r.POST("/register", func(c *gin.Context) {
		var pi *apis.PersonalInformation
		err := c.BindJSON(&pi)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "无法解析注册信息",
			})
			return
		}
		if err := rankServer.RegisterPersonalInformation(pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "注册失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "注册成功",
		})
	})

	r.PUT("/personalInfo", func(c *gin.Context) {
		var pi *apis.PersonalInformation
		if err := c.BindJSON(&pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "无法解析更新信息",
			})
			return
		}
		if reps, err := rankServer.UpdatePersonalInformation(pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "更新失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, reps)
		}
	})

	r.GET("/rank/:name", func(c *gin.Context) {
		name := c.Param("name")
		if fr, err := rankServer.GetFatRate(name); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "获取个人体脂率失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, fr)
		}
	})

	r.GET("/ranktop", func(c *gin.Context) {
		if topFrs, err := rankServer.GetTop(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "获取体脂率排行失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, topFrs)
		}
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("启动端口失败", err)
	}
}

func connectDB() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("chenxi:123456@tcp(127.0.0.1:3306)/test"))
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	fmt.Println("连接数据库成功")
	return conn
}
