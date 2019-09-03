package api

import (
	"github.com/cy422396350/cygin/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func IndexApi(c *gin.Context) {
	p := models.Person{Id: 1}
	person, err := p.GetPerson()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"person": person,
	})
}
