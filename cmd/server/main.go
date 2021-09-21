package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kekscode/rincewind/internal/timekeeper"
)

func main() {
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-term:
			log.Println("Received SIGTERM, exiting gracefully...")
			os.Exit(0)
		default:
			t := timekeeper.New("work", time.Duration(time.Minute*25))
			t.Start()
			fmt.Printf("Begin: %v\n", t.Begin)
			fmt.Printf("End: %v\n", t.End)
			fmt.Printf("Elapsed: %v\n", t.TimeElapsed())
			fmt.Printf("Left: %v\n", t.TimeLeft())
			fmt.Printf("Timer done: %v\n", t.Expired())
		}
	}
}
