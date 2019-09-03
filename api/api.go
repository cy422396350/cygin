package api

import (
	"github.com/cy422396350/cygin/models"
	"github.com/gin-gonic/gin"
	"log"
)

func GetAuthUserInfo() (name string, password string) {
	p := models.Person{}
	name, err := p.GetKey(gin.AuthUserKey)
	if err != nil {
		log.Fatalln(err)
	}
	password, err = p.GetKey("password")
	return
}

func IndexApi(c *gin.Context) {
	i := models.IndexData{}
	res, err := i.GetCategory()
	if err != nil {
		log.Fatalln(err)
	}
	c.String(200, res)
}
