package main

import (
	"awesomeProject/main/listener"
	"awesomeProject/util/file"
	"awesomeProject/util/mq"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
)

func main() {
	log.Printf("Config: %v", *file.GetConf())
	ctx, cancel := context.WithCancel(context.Background())

	scanner := bufio.NewScanner(os.Stdin)
	go listener.ListenTestQueue()
	go func() {
		for true {
			scanner.Scan()
			go mq.SendMess(scanner.Text(), file.GetConf().TestQueue, ctx)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Kill, os.Interrupt)
	sig := <-sigChan
	fmt.Printf("Get sig: %v, stop service", sig)
	fmt.Println("finish:")
	cancel()

}
