package main

import "math"

// InterestCalculator provides methods for calculating interest
type InterestCalculator struct {
	AnnualRate float64
}

// NewInterestCalculator creates a new instance with default 12% annual rate
func NewInterestCalculator() *InterestCalculator {
	return &InterestCalculator{
		AnnualRate: 0.12, // 12% annual rate
	}
}

// CalculateSimpleInterest calculates simple interest for a given principal and time (in years)
func (ic *InterestCalculator) CalculateSimpleInterest(principal float64, years float64) float64 {
	return principal * ic.AnnualRate * years
}

// CalculateCompoundInterest calculates compound interest for a given principal and time (in years)
func (ic *InterestCalculator) CalculateCompoundInterest(principal float64, years float64) float64 {
	return principal * math.Pow(1+ic.AnnualRate, years)
}

// CalculateMonthlyPayment calculates the monthly payment for a loan
func (ic *InterestCalculator) CalculateMonthlyPayment(principal float64, years float64) float64 {
	monthlyRate := ic.AnnualRate / 12
	numberOfPayments := years * 12

	// Monthly payment formula: P * r * (1 + r)^n / ((1 + r)^n - 1)
	// where P is principal, r is monthly rate, n is number of payments
	numerator := principal * monthlyRate * math.Pow(1+monthlyRate, numberOfPayments)
	denominator := math.Pow(1+monthlyRate, numberOfPayments) - 1

	return numerator / denominator
}
