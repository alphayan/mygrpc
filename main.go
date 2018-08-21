package main

import (
	"mygrpc/client"
	"mygrpc/server"
	"time"
)

func main() {
	go func() {
		for {
			client.Pub()
			time.Sleep(time.Second)
		}
	}()
	go client.Sub()
	go client.Run()
	defer client.Conn().Close()
	server.Run()
}
