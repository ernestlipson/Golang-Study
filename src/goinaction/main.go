package main

import (
	"fmt"
	"io"
	"os"
)

var name string = "Ernest Darko"
var age uint64 = 100876

func main (){
	fmt.Fprint(io.Discard)
	
	message:= fmt.Sprintf("Hello, My name is %s.", name)
	name_length:= len(name)
	fmt.Println("Length of String,", name_length)
	fmt.Println(message)
	fmt.Println("Numeric types:", 2+9)
	os.Exit(200)
}