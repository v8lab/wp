package model

import (
	Cli "model/client"
	DB "model/db"
)

func Try() {
	var Clients Cli.ClientsStu
	Clients.Id = "clients01"
	Redis := DB.GetSingleRedisDbStu()
	Redis.Do("HSET")
}
