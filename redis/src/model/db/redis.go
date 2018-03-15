package db

import (
	"fmt"
	"sync"

	Redis "github.com/garyburd/redigo/redis"
)

var SingleRedisDb RedisDbStu
var SingleRedisDbOnce sync.Once

func GetSingleRedisDbStu() *RedisDbStu {
	SingleRedisDbOnce.Do(
		func() {
			SingleRedisDb.Init()
		})
	return &SingleRedisDb
}

type RedisDbStu struct {
	Redis.Conn
}

func (r *RedisDbStu) Init() (err error) {
	r.Conn, err = Redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis dial error")
		return
	}
	fmt.Println("redis dial success")
	return
}
