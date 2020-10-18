package main

import "fmt"

func main() {
	fmt.Println("hello world this is fist code me in go an enter your name please : ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("hello " + name)
	fmt.Scanln(&name)
}
