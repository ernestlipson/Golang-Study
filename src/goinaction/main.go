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
	doubleNum()
	fmt.Println(convert_temp())
	fmt.Println(print_num(100))
	os.Exit(200)
}

func doubleNum()  {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2 * math.Pi
	fmt.Println(output)
}

func convert_temp() (float64, error) {
	fmt.Print("Enter a Temperature: ")
	var temp float64
	fmt.Scanf("%f", &temp)
	if temp < -459.67 {
		return 0, fmt.Errorf("Invalid Temperature: %v", temp)
	}
	if temp >= -273.15 && temp <= 100 {
		return temp * 9/5 + 32, nil
	}
	return (temp -32) * 5/9, nil
}

func print_num(val int) (int,int, error){
	var evenCount int
	var oddCount int
	if (val < 0) {
		return 0, 0, fmt.Errorf("Invalid number. Value must be greater than 0")
	}
	for i := 0; i <= val; i++ {
		if i % 2 == 0 {
			evenCount += i
		} else {
			oddCount += i
		}
		fmt.Println(i)
	}
	return evenCount, oddCount, nil
}