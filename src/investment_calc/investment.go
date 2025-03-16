package investmentcalc

import (
	"fmt"
	"math"
)

func CalculateInvestmentFutureValues(investment, rate, years float64) (float64, float64) {
	future := investment * math.Pow(1+rate, years)

	// Calculate future real value with 1% rate increase per year
	futureRealValue := investment
	currentRate := rate
	for i := 0; i < int(years); i++ {
		futureRealValue *= (1 + currentRate)
		currentRate += 0.01 // Increase rate by 1%
	}

	fmt.Println("Future value of investment:", future)
	fmt.Println("Future real value (with 1% annual rate increase):", futureRealValue)

	return future, futureRealValue
}
