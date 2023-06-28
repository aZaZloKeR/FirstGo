package listener

import (
	"awesomeProject/main/service"
	"awesomeProject/util/config"
	"awesomeProject/util/mq"
	"context"
)

func ListenTestQueue() {
	ch := make(chan string)
	go func() {
		for true {
			mq.ReadMess(context.Background(), config.Get().TestQueue)
		}
	}()
	for true {
		body := <-ch
		if body != "" {
			service.DoSmth(body)
		}
	}
}
