package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

var name string = "Ernest Darko"
var age uint64 = 100876

func main (){
	fmt.Fprint(io.Discard)
	
	message:= fmt.Sprintf("Hello, My name is %s.", name)
	name_length:= len(name)
	val_mult:= 32132 * 42452
	fmt.Println("Mult is,", val_mult)
	fmt.Println("Length of String,", name_length)
	fmt.Println(message[3])
	fmt.Println("Numeric types:", 2+9)
	os.Exit(200)
	doubleNum()
}

func doubleNum()  {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2 * math.Pi
	fmt.Println(output)
}