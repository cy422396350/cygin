package redis

import (
	"github.com/chasex/redis-go-cluster"
	"log"
	"time"
)

var RedisCy *redis.Cluster

func init() {
	var err error
	RedisCy, err = redis.NewCluster(
		&redis.Options{
			StartNodes:   []string{"127.0.0.1:6391", "127.0.0.1:6392", "127.0.0.1:6393", "127.0.0.1:6394", "127.0.0.1:6395", "127.0.0.1:6396"},
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})
	if err != nil {
		log.Fatal(err.Error())
	}
}
