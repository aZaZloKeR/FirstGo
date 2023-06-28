package main

import (
	"awesomeProject/util/config"
	"awesomeProject/util/mq"
	"awesomeProject/util/syncer"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
)

func main() {
	log.Printf("Config: %v", *config.Get())
	ctx, done := context.WithCancel(context.Background())
	go func() {
		mq.ReadMess(ctx, config.Get().TestQueue)
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for true {
		scanner.Scan()
		go mq.SendMess(ctx, scanner.Text(), config.Get().TestQueue)
	}

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Kill, os.Interrupt)
	signal := <-sigChan
	fmt.Printf("Get signal: %v, stop service", signal)
	done()
	syncer.Wait()
}
