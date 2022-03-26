package main

import (
	"encoding/base64"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"learngo/pkg/apis"
)

func main() {
	r := gin.Default()
	pprof.Register(r)
	r.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte("你好,gin"))
	})
	r.GET("/history", func(c *gin.Context) {
		c.JSON(200, []*apis.PersonalInformation{ //如果使用.JSON，就可以返回json格式，并自带content-type:application/json
			{
				Name:   "小强",
				Sex:    "男",
				Tall:   1.70,
				Weight: 80,
				Age:    30,
			},
			{
				Name:   "小李",
				Sex:    "男",
				Tall:   1.90,
				Weight: 80,
				Age:    20,
			},
		})
	})
	r.POST("/register", func(c *gin.Context) {
		pi := &apis.PersonalInformation{}
		if err := c.BindJSON(pi); err != nil {
			c.JSON(400, gin.H{
				"message": "无法读取个人注册信息",
			})
			return
		}
		//todo 注册到排行榜
	})
	r.GET("/getWithQuery", func(c *gin.Context) {
		name := c.Query("name")
		c.JSON(400, gin.H{
			"payload": base64.StdEncoding.EncodeToString([]byte(name)),
		})
	})
	r.GET("/getWithParam/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(400, gin.H{
			"payload": base64.StdEncoding.EncodeToString([]byte(name)),
		})
	})
	r.GET("/getWithParam/111", func(c *gin.Context) { //更高优先级
		name := c.Param("name")
		c.JSON(400, gin.H{
			"special": base64.StdEncoding.EncodeToString([]byte(name)),
		})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
