package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func main() {
	fmt.Printf("此时goroutine的个数: %d\n", runtime.NumGoroutine())
	num := 6
	for index := 0; index < num; index++ {
		http.Get("https://www.baidu.com")
		//_, _ = ioutil.ReadAll(resp.Body)
		//_ = resp.Body.Close()
	}
	fmt.Printf("此时goroutine的个数: %d\n", runtime.NumGoroutine())
}
