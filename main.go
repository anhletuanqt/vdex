package main

import (
	"fmt"
	"time"
)

const (
	defaultConfigPath = "./conf.json"
)

func main() {
	// config.Init(defaultConfigPath)
	// redisClient.NewRedis()

	// matching.StartEngine()

	// worker.NewFillExecutor().Start()
	// worker.NewSettlementExecutor().Start()
	// event_listener.NewOnchainEvent().Start()
	// worker.Start()

	// pushing.StartServer()
	// api.StartServer()
	testt()
}

func testt() {
	c := make(chan int)

	go func() {
		for s := range c {
			fmt.Println("g1", s)
		}
	}()
	go func() {
		for s := range c {
			fmt.Println("g2", s)
		}
	}()
	c <- 1
	c <- 2

	time.Sleep(20 * time.Second)
}
