package main

import (
	"awesomeProject/main/listener"
	"awesomeProject/util/file"
	"awesomeProject/util/mq"
	"fmt"
	"time"
)

func main() {
	go mq.SendMess("Hello, Its my first message from go!", file.GetConf("test_queue"))

	go listener.ListenTestQueue()
	time.Sleep(2 * time.Minute)
	fmt.Println("finish:")

}
