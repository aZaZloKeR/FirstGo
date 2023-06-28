package main

import (
	"awesomeProject/util/file"
	"awesomeProject/util/mq"
	"fmt"
)

func main() {
	mq.SendMess("Hello, Its my first message from go!", file.GetConf("test_queue"))

	a := mq.ReadMess(file.GetConf("test_queue"))
	fmt.Println("get message: ", a)
}
