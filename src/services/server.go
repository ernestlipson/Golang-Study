package main

import (
	"fmt"
	"strings"
)

// TODO: define the 'OvenTime' constant
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	extraCar := carsCount % 10
	carsOfTen := carsCount / 10
	return uint(carsOfTen)*95000 + uint(extraCar)*10000
}

func CanFastAttack(knightIsAwake bool) bool {
	if knightIsAwake == true {
		return false
	} else {
		return true
	}
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if petDogIsPresent && !archerIsAwake {
		return true
	}
	if prisonerIsAwake && !petDogIsPresent && !archerIsAwake && !knightIsAwake {
		return true
	}
	return false
}

// Alternative function to the CanFreePrisoner
func FreePrisonerAlt(knightIsAwake,
	archerIsAwake,
	prisonerIsAwake,
	petDogIsPresent bool) bool {
	return !knightIsAwake && !archerIsAwake && prisonerIsAwake || !archerIsAwake && petDogIsPresent

}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	return border + "\n" + welcomeMsg + "\n" + border
}

func CleanupMessage(oldMsg string) string {
	message := strings.ReplaceAll(oldMsg, "*", "")
	output := strings.TrimSpace(message)
	return output
}

func main() {
	fmt.Println(RemainingOvenTime(12))
	fmt.Println(PreparationTime(8))
	fmt.Println(ElapsedTime(8, 12))
	fmt.Println(CalculateWorkingCarsPerHour(1547, 90))
	fmt.Println(CalculateWorkingCarsPerMinute(1105, 90))
	fmt.Println(CalculateCost(37))
	fmt.Println(CalculateCost(6))
	fmt.Println(CanFastAttack(false))

	fmt.Println(CanSpy(false, true, false))
	fmt.Println(CanSpy(false, false, false))
	fmt.Printf("CanFreePrisoner: %v\n", CanFreePrisoner(false, true, false, true))
	fmt.Printf("FreePrisonerAlternative: %v\n", FreePrisonerAlt(false, true, false, true))

	strings.ToLower("ANY AnabEL")
	fmt.Println(AddBorder("Akwaaba", 12))
	message := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`
	fmt.Println(CleanupMessage(message))
}
