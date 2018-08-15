package main

import (
	"mygrpc/client"
	"mygrpc/server"
)

func main() {
	go client.Run()
	server.Run()
}
