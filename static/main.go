package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	BirthDay time.Time `form:"birthday" time_format:"2006-01-02 15:04:05"`
}

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523433"},
}

func bindTest(ctx *gin.Context) {
	var person Person
	if err := ctx.ShouldBind(&person); err != nil {
		ctx.String(http.StatusBadRequest, "bind Person err:%v", err)
		ctx.Abort()
	}
	ctx.String(http.StatusOK, "%v", person)
}

func getSecrets(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	if secret, ok := secrets[user]; ok {
		c.JSON(200, gin.H{"user": user, "secret": secret})
	} else {
		c.JSON(200, gin.H{"user": user, "secret": "No SECRET :("})
	}
}

func main() {
	r := gin.Default()
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar", //用户名：密码
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	authorized.POST("/secrets", getSecrets)
	r.Static("/images", "./images")
	r.GET("/n:name/:id", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"name": context.Param("name"),
			"id":   context.Param("id"),
		})
	})
	r.GET("/user", func(context *gin.Context) {
		firstName := context.Query("first")
		lastName := context.DefaultQuery("last", "default_last")
		context.String(http.StatusOK, "%s %s", firstName, lastName)
	})
	r.POST("/body", func(context *gin.Context) {
		bytes, e := ioutil.ReadAll(context.Request.Body)
		if e != nil {
			context.String(http.StatusBadRequest, e.Error())
			context.Abort()
		}
		context.String(http.StatusOK, string(bytes))
	})
	r.GET("/testing", bindTest)
	r.POST("/testing", bindTest)

	r.Run(":8080")
}
