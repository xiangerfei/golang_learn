package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
	//"golang_learn/go_gin/utils/string_utils"
	"gofirst/go_bin/utils/string_utils"
)

func RandomString(n int) string{
	//var letters = []byte("abcdefghjklmnopq")
	var letters string = "abcdefghjklmnopq"
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}


func main(){
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/api/auth/register", func(c *gin.Context){
		// 获取参数
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")

		// 数据验证
		if len(telephone) != 11{
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号不能低于11位" + string(telephone)})
			fmt.Printf("hello")
		}

		if len(password) < 6{
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 423, "msg": "密码不能少于16位"})
			return
		}

		if len(name) == 0{
			name = string_utils.RandomString(10)
			fmt.Printf("%s\n", name)
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 424, "msg": name})
			return
		}
	})
	r.Run("0.0.0.0:8081")
}


