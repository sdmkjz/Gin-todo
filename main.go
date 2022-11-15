package main

import (
	"Gin_todo/conf"
	"Gin_todo/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
