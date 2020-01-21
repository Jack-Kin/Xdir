package main

import "fmt"

func main() {
	var messages chan string = make(chan string)
	go func(message string) {
		messages <- message // 存消息
	}("h")
	fmt.Println(<-messages) // 取消息
	go func(message string) {
		messages <- message // 存消息
	}("W!")
	fmt.Println(<-messages) // 取消息
}