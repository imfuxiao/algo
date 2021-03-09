package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c1 := make(chan struct{}, 1)
	c2 := make(chan struct{}, 1)
	c3 := make(chan struct{}, 1)
	c4 := make(chan struct{}, 1)

	defer func() {
		close(c1)
		close(c2)
		close(c3)
		close(c4)
	}()

	// goroutine 1
	go func() {
		for {
			select {
			case <-c1:
				fmt.Println("1")
				c2 <- struct{}{}
			}
		}
	}()

	// goroutine 2
	go func() {
		for {
			select {
			case <-c2:
				fmt.Println("2")
				c3 <- struct{}{}
			}
		}
	}()

	// goroutine 3
	go func() {
		for {
			select {
			case <-c3:
				fmt.Println("3")
				c4 <- struct{}{}
			}
		}
	}()

	// goroutine 4
	go func() {
		for {
			select {
			case <-c4:
				fmt.Println("4")
				fmt.Println("------")
			}
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-time.After(1 * time.Second):
			c1 <- struct{}{}
		case <-c:
			os.Exit(0)
		}
	}
}
