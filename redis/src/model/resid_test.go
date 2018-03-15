package model

import (
	"fmt"
	Cli "model/client"
	Db "model/db"
	"testing"

	Redis "github.com/garyburd/redigo/redis"
)

func Test_Redis(t *testing.T) {
	var Client Cli.ClientStu
	var Clients Cli.ClientsStu
	Clients.Clients = append(Clients.Clients, Client)

	Clients.Id = "clients01"
	Clients.ClientsStr = "ClientsStr001"
	c := Db.GetSingleRedisDbStu()
	c.Do("HMSET", Redis.Args{}.Add(Clients.GetId()).Add("Id").Add("clients").AddFlat(&Clients)...)

	var Clients02 Cli.ClientsStu
	v, err := Redis.Values(c.Do("HGETALL", Redis.Args{}.Add(Clients.GetId())...))
	if err != nil {
		panic(err)
	}
	fmt.Println(Clients02)
	if err := Redis.ScanStruct(v, &Clients02); err != nil {
		panic(err)
	}
	fmt.Println(Clients02)
}
