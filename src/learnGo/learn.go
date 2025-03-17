package main

import (
	"fmt"
	"math"
)

var (
	actorNmae string = "Elizabeth"
	companion string = "Smith"
	season    int    = 12
)

func main() {
	var declare int = 78
	newDeclare := 99

	var c CircleDef
	circumference := new(CircleDef)
	// circle := CircleDef{x: 90, y: 23, z: 12}
	circumf := CircleDef{x: 2, y: 8, z: 18}
	fmt.Printf("c: %v\n", c)
	fmt.Printf("circumf: %v\n", circumf)
	fmt.Printf("circumference: %v\n", circumference)

	defer fmt.Println(fibonacciSequence(1000))
	fmt.Printf("Address And location of newDeclare is %p", &newDeclare)
	fmt.Println(declare, newDeclare)

	fmt.Println(halfIntVal(89))
	fmt.Println(findLargeNumber(89, 45, 6))
	hasNeg, negNums := hasNegative(-56, 89, 45, 6, -4, -8, -56, 10000)
	fmt.Println(hasNeg, negNums)
	fmt.Println(hasNegative(89, 45, 6, -4, -8, -56, 10000))

	fmt.Println(circleArea(circumf))
	fmt.Printf("ptrsCircleArea: %v\n", ptrsCircleArea(&circumf))
	fmt.Printf("c.area(): %v\n", c.area())
	rectangle := Rectangle{topSide: 100, rightSide: 200, leftSide: 200}
	fmt.Printf("The Area of the Rectangle Is: %v\n", rectangle.rectDist())

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

func personDef(diameter float64, circumference float64, radius float64) CircleDef {
	personDef := CircleDef{x: diameter, y: circumference, z: radius}
	fmt.Println(personDef.x)
	return personDef
}

func circleArea(param CircleDef) float64 {
	return math.Pi * (param.x * param.x)
}

func ptrsCircleArea(circle *CircleDef) float64 {
	return math.Pi * circle.x * circle.x
}

func (c *CircleDef) area() float64 {
	return math.Pi * c.x * c.x
}

type Rectangle struct {
	topSide, rightSide, bottomSide, leftSide float64
}

func (rectangle *Rectangle) rectDist() float64 {
	lenth := rectangle.topSide
	width := rectangle.leftSide
	return lenth * width
}
