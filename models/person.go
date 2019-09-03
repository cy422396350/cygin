package models

import (
	redis2 "github.com/chasex/redis-go-cluster"
	"github.com/cy422396350/cygin/database"
	"github.com/cy422396350/cygin/redis"
	"log"
)

type HandlePerson interface {
	GetPerson() (person Person, err error)
	GetUser() (name string, err error)
}

type Person struct {
	Id       int
	Name     string
	Password string
}

func (p *Person) GetPerson() (person Person, err error) {
	err = database.SqlDB.QueryRow("SELECT id, name, password FROM person WHERE id=?", p.Id).Scan(
		&person.Id, &person.Name, &person.Password,
	)
	return
}
func (p *Person) GetKey(key string) (name string, err error) {
	name, err = redis2.String(redis.RedisCy.Do("GET", key))
	if err != nil {
		log.Fatal(err)
	}
	return
}
