package main

import "fmt"

var str string

func main() {
	fmt.Print(str + " " + string(rune(96)))
	fmt.Println(str + " " + string(rune(96)))
	fmt.Println(string(rune(41)))
}

func init() {
	str = `
package main

import "fmt"

var str string

func main() {
	fmt.Print(str + " " + string(rune(96)))
	fmt.Println(str + " " + string(rune(96)))
	fmt.Println(string(rune(41)))
}

func init() {
		str = `
}
