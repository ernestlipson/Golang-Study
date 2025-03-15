package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

var name string = "Ernest Darko"

func main() {
	fmt.Fprint(io.Discard)
	defer os.Exit(0)

	message := fmt.Sprintf("Hello, My name is %s.", name)
	fmt.Println("Mult is,", 32132*42452)
	fmt.Println("Length of String,", len(name))
	fmt.Println(message[3])
	fmt.Println("Numeric types:", 2+9)
	doubleNum()

	temp, err := convertTemp()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(temp)
	}
	evenSum, oddSum, err := printNum(100)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Even Sum:", evenSum, "Odd Sum:", oddSum)
	}
	fizzBuzz(90)
	fizzBuzzSwitch(99)

	makeSlice()
	goMap()

	arr := []int{8, 12, 89, 12, 9, 8, 9, 4, 4, 2, 67, 89, 9000}
	fmt.Println("Smallest number:", findSmallest(arr))
	fmt.Println("Largest number:", findLargest(arr))

	xarray := []float64{8, 12, 89, 12, 9, 8, 9, 4, 4, 2, 67, 89, 9000}
	small, large := findSmallestLargest(xarray)
	fmt.Println("Smallest and largest numbers:", small, large)

	xs := []int{4, 78, 99, 1000, 89, 99, 67, -100}
	sum, mult := sumWithMinMaxProduct(xs...)
	fmt.Println("Sum and product:", sum, mult)

	nextEven := generateEvenNumber()
	fmt.Println(nextEven(), nextEven(), nextEven())

	fmt.Println("Factorial of 2:", factorial(2))
	fmt.Println("Factorial of 67:", factorial(67))

	xPtr := new(int)
	assignNumber(xPtr)
	fmt.Println(*xPtr)
	fmt.Println(strings.Contains("ernest", "ern"))
}

func doubleNum() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	fmt.Println(input * 2 * math.Pi)
}

func convertTemp() (float64, error) {
	fmt.Print("Enter a Temperature: ")
	var temp float64
	fmt.Scanf("%f", &temp)
	if temp < -459.67 {
		return 0, fmt.Errorf("Invalid Temperature: %v", temp)
	}
	if temp >= -273.15 && temp <= 100 {
		return temp*9/5 + 32, nil
	}
	return (temp - 32) * 5 / 9, nil
}

func printNum(val int) (int, int, error) {
	if val < 0 {
		return 0, 0, fmt.Errorf("Invalid number. Value must be greater than 0")
	}
	var evenSum, oddSum int
	for i := 0; i <= val; i++ {
		if i%2 == 0 {
			evenSum += i
		} else {
			oddSum += i
		}
		fmt.Println(i)
	}
	return evenSum, oddSum, nil
}

func fizzBuzz(val int) {
	for i := 1; i <= val; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func fizzBuzzSwitch(val int) {
	for i := 1; i <= val; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func makeSlice() {
	slice := append(make([]int, 19), 99, 88, 12, 78)
	fmt.Println(slice)
}

func goMap() {
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	if chemical, ok := elements["Uriah"]; ok {
		fmt.Println(ok, chemical)
	}
}

func findSmallest(arr []int) int {
	smallest := arr[0]
	for _, v := range arr {
		if v < smallest {
			smallest = v
		}
	}
	return smallest
}

func findLargest(arr []int) int {
	largest := arr[0]
	for _, v := range arr {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func findSmallestLargest(xs []float64) (float64, float64) {
	smallest, largest := xs[0], xs[0]
	for _, v := range xs {
		if v < smallest {
			smallest = v
		}
		if v > largest {
			largest = v
		}
	}
	return smallest, largest
}

func sumWithMinMaxProduct(args ...int) (int, int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()
	for _, v := range args {
		if v < 0 {
			panic("ERROR: Signed Negative Integer. Number should be a signed Positive")
		}
	}
	least, most := args[0], args[0]
	for _, v := range args {
		if v < least {
			least = v
		}
		if v > most {
			most = v
		}
	}
	total := 0
	for _, v := range args {
		total += v
	}
	return total, least * most
}

func generateEvenNumber() func() uint {
	i := uint(0)
	return func() uint {
		ret := i
		i += 2
		return ret
	}
}

func factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func assignNumber(xPtr *int) {
	*xPtr = 4
}
