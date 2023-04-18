package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

var name string = "Ernest Darko"
var age uint64 = 100876

func main() {
	fmt.Fprint(io.Discard)
	defer os.Exit(200)

	message := fmt.Sprintf("Hello, My name is %s.", name)
	name_length := len(name)
	val_mult := 32132 * 42452
	fmt.Println("Mult is,", val_mult)
	fmt.Println("Length of String,", name_length)
	fmt.Println(message[3])
	fmt.Println("Numeric types:", 2+9)
	doubleNum()

	fmt.Println(convert_temp())
	fmt.Println(print_num(100))
	fizz_buzz(90)
	fizz_buzz_s(99)

	make_slice()
	go_map()

	classic_array([]int{8, 12, 89, 12, 9, 8, 9, 4, 4, 2, 67, 89, 9000})
	fmt.Println(small_number([]int{0, 8, 12, 89, 12, 9, 8, 9, 4, 4, 2, 67, 89, 9000}))
	fmt.Println(largest_number([]int{0, 8, 12, 89, 12, 9, 8, 9, 4, 4, 2, 67, 89}))

	xarray := []float64{8, 12, 89, 12, 9, 8, 9, 4, 4, 2, 67, 89, 9000}
	fmt.Println(arraySmLg(xarray))
	small, large := arraySmLg(xarray)
	fmt.Println("The small and large Numbers: ", small, large)

	xs := []int{4, 78, 99, 1000, 89, 99, 67, -100}
	sum, mult := sumWithMinMaxProduct(xs...)
	fmt.Println("Sum of XS is & multiply is ", sum, mult)
	sum_A, mult_A := sumWithMinMaxProduct(2, 8, 19, 100000, 868, 996, -2)
	fmt.Println("Sum of SOME STRANGE ARRAY ", sum_A, mult_A)

	nextEven := generateEvenNumber()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	fmt.Println("The factorial is:", facotrial(2))

	fmt.Println("The factorial is:", facotrial(67))

}

func doubleNum() {
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
		return temp*9/5 + 32, nil
	}
	return (temp - 32) * 5 / 9, nil
}

func print_num(val int) (int, int, error) {
	var evenCount int
	var oddCount int
	if val < 0 {
		return 0, 0, fmt.Errorf("Invalid number. Value must be greater than 0")
	}
	for i := 0; i <= val; i++ {
		if i%2 == 0 {
			evenCount += i
		} else {
			oddCount += i
		}
		fmt.Println(i)
	}
	return evenCount, oddCount, nil
}

func fizz_buzz(val int) {
	for i := 0; i < val; i++ {
		if i == 0 {
			continue
		}
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}

}

func fizz_buzz_s(val int) {
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

func classic_array(inputArray []int) {
	var totalUser int = 0
	for _, value := range inputArray {
		totalUser += value
	}
	fmt.Println("Total Array elements: ", totalUser)
	fmt.Println("Mean/Average: ", totalUser/len(inputArray))
}

func make_slice() {
	slice := append(make([]int, 19), 99, 88, 12, 78)
	fmt.Println(slice)
}

func copy_func() {
	slice := []int{9, 56, 88, 78, 33}
	slice_two := []int{9, 56, 88, 78, 33}
	copy(slice, slice_two)
}

func go_map() {
	new_map := make(map[int]string)
	new_map[8] = "Seventy Eight"
	fmt.Println(new_map)

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
		"Ne": "Neon"}

	if chemical, ok := elements["Uriah"]; ok {
		fmt.Println(ok, chemical)
	}
}

func small_number(inputArray []int) int {
	small_val := inputArray[0]
	for _, value := range inputArray {
		if value < small_val {
			small_val = value
		}
	}
	return small_val
}

func largest_number(inputArray []int) int {
	large_val := inputArray[0]
	for i := 0; i < len(inputArray); i++ {
		if inputArray[i] > large_val {
			large_val = inputArray[i]
		}
	}
	return large_val
}

func arraySmLg(xs []float64) (float64, float64) {
	sm_number := xs[0]
	lg_number := xs[len(xs)-1]
	for _, v := range xs {
		if v < sm_number {
			sm_number = v
		} else if v > lg_number {
			lg_number = v
		}
	}
	return sm_number, lg_number
}

func sumWithMinMaxProduct(args ...int) (int, int) {
	for _, v := range args {
		if v < 0 {
			defer func() {
				str := recover()
				fmt.Println(str)
			}()
			panic("ERROR: Signed Negative Integer. Number should be a signed Positive")
		}
	}
	least, most := args[0], args[len(args)-1]
	var multiply int
	for _, v := range args {
		if v < least {
			least = v
		} else if v > most {
			most = v
		}
		multiply = least * most
	}
	total := 0
	for _, v := range args {
		total += v
	}
	return total, multiply
}

func generateEvenNumber() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func facotrial(fact uint) uint {
	if fact == 0 {
		return 1
	}
	return fact * facotrial(fact-1)
}
