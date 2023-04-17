package main

import (
	"fmt"
	"time"
)

func hello(){
	fmt.Println("Trying out a goroutine")
}
func main(){
	go hello()
	fmt.Println("main function")
	time.Sleep(1 * time.Second)
	fmt.Println("Another Main fxn")
	fmt.Println("Hello World!")
}