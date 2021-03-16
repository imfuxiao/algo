package main

import "fmt"

func main() {
	var flag uint8 = 0
	for i := 0; i < 10; i++ {
		fmt.Println(flag)
		flag = ^flag
	}
}
