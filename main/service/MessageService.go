package service

import (
	"fmt"
	"time"
)

func DoSmth(body string) {
	time.Sleep(3 * time.Second)
	fmt.Println("I am receive message: ", body)
}
