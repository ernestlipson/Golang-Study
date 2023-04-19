package main

import "fmt"

var (
	actorNmae string = "Elizabeth"
	companion string = "Smith"
	season    int    = 12
)

func main() {
	var declare int = 78
	newDeclare := 99

	defer fmt.Println(fibonacciSequence(1000))
	fmt.Printf("Address And location of newDeclare is %p", &newDeclare)
	fmt.Println(declare, newDeclare)

	fmt.Println(halfIntVal(89))
	fmt.Println(findLargeNumber(89, 45, 6))
	hasNeg, negNums := hasNegative(-56, 89, 45, 6, -4, -8, -56, 10000)
	fmt.Println(hasNeg, negNums)
	fmt.Println(hasNegative(89, 45, 6, -4, -8, -56, 10000))
}

func halfIntVal(e int) (int, bool) {
	if e%2 == 0 {
		return 1, true
	}
	return 0, false
}

func findLargeNumber(args ...int) int {
	initNum := args[0]
	for _, v := range args {
		if v > initNum {
			initNum = v
		}
	}
	return initNum
}

func generateOddNumber() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 3
		return
	}
}

func fibonacciSequence(variable int) int {

	if variable <= 1 {
		return variable
	}

	var a, b, c = 0, 1, 0
	for i := 2; i <= variable; i++ {
		c = a + b
		a = b
		b = c
	}

	return c
}

func hasNegative(args ...int) (bool, []int) {
	negativeList := []int{}
	for _, v := range args {
		if v < 0 {
			negativeList = append(negativeList, v)
		}
	}
	if len(negativeList) > 0 {
		return true, negativeList
	}
	return false, nil
}

type CircleDef struct {
	x float64
	y float64
	z float64
}

func person_def(diameter float64, circumference float64, radius float64) CircleDef {
	person_def := CircleDef{x: diameter, y: circumference, z: radius}
	fmt.Println(person_def.x)
	return person_def
}
