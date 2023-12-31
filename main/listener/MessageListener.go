package listener

import (
	"awesomeProject/main/service"
	"awesomeProject/util/file"
	"awesomeProject/util/mq"
)

func ListenTestQueue() {
	ch := make(chan string)
	go func() {
		for true {
			mq.ReadMess(file.GetConf().TestQueue, ch)
		}
	}()
	for true {
		body := <-ch
		if body != "" {
			service.DoSmth(body)
		}
	}
}
