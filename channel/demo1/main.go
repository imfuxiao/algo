package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan struct{})
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println("test", v)
			}
		}
	}()
	c <- struct{}{}
	<-time.After(5 * time.Second)
	close(c)
	<-time.After(5 * time.Second)
}
func t(c chan struct{}) {
	close(c)
}
