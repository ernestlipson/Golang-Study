package investmentcalc

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmt = 1000
	var interestRate = 0.05
	var years = 5

	var futureValue = float64(investmentAmt) * math.Pow(1+interestRate, float64(years))
	fmt.Println("Future value of investment:", futureValue)
}
