package main

import (
	"awesomeProject/mq"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ch := make(chan int)
	go mq.CreateConnection(ch)

	mq.SendMess("Hello, Its my first message from go!", &ctx)

	fmt.Println("finish")
}
