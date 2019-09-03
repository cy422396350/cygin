package api

import (
	"github.com/cy422396350/cygin/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
	p := models.Person{Id: 1}
	person, err := p.GetPerson()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"person": person,
		"secret": 111,
	})
}
