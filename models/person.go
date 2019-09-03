package models

import (
	redis2 "github.com/chasex/redis-go-cluster"
	"github.com/cy422396350/cygin/redis"
	"log"
)

type HandlePerson interface {
	GetKey(key string) (name string, err error)
}

type Person struct {
	Id       int
	Name     string
	Password string
}

func (p *Person) GetKey(key string) (name string, err error) {
	name, err = redis2.String(redis.RedisCy.Do("GET", key))
	if err != nil {
		log.Fatal(err)
	}
	return
}
