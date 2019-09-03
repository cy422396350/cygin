package models

import (
	"github.com/chasex/redis-go-cluster"
	redis2 "github.com/cy422396350/cygin/redis"
	"log"
)

type HandleIndex interface {
	GetCategory() (res string, err error)
}

type IndexData struct {
}

func (i *IndexData) GetCategory() (res string, err error) {
	res, err = redis.String(redis2.RedisCy.Do("GET", "indexc"))
	if err != nil {
		log.Fatal(err)
	}
	return
}
