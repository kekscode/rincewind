package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case <-term:
		log.Println("Received SIGTERM, exiting gracefully...")
		os.Exit(0)
	}
}
