package main

import (
	"awesomeProject/main/listener"
	"awesomeProject/util/file"
	"awesomeProject/util/mq"
	"fmt"
	"log"
	"time"
)

func main() {
	log.Printf("Config: %v", *file.GetConf())
	go mq.SendMess("Hello, Its my first message from go!", file.GetConf().TestQueue)
	go listener.ListenTestQueue()
	time.Sleep(2 * time.Minute)
	fmt.Println("finish:")

}
