package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// define the 'OvenTime' constant
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

func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

func HappyBirthday(name string, age int) string {
	message := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
	return message
}

func AssignTable(name string, table int,
	neighbor, direction string, distance float64) string {
	message := fmt.Sprintf("Welcome to my party %s\nYou have been assigned to table %03d. Your table is %s, exactly %.2f meters from here.\nYou will be sitting next to %s", name, table, direction, distance, neighbor)
	return message
}

func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return option1 + " is clearly the better choice."
	} else {
		return option2 + " is clearly the better choice"
	}
}

func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return 0.8 * originalPrice
	} else if age >= 10 {
		return 0.5 * originalPrice
	} else {
		return 0.7 * originalPrice
	}
}

func findNumAtIndex(num int, nums ...int) {
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			return
		}
	}
	fmt.Println(num, "not found in ", nums)
}

func FavoriteCards() []int {
	magiCards := []int{2, 6, 9}
	return magiCards
}

func GetItem(slice []int, index int) int {
	if index >= len(slice) || index < 0 {
		return -1
	}
	value := slice[index]
	return value
}

func SetItem(slice []int, index, value int) []int {
	if index >= len(slice) || index < 0 {
		return append(slice, value)
	}
	slice[index] = value
	return slice
}

func PrependItems(slice []int, values ...int) []int {
	if len(values) == 0 {
		return slice
	}
	values = append(values, slice...)
	return values
}

func RemoveItem(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	slice = append(slice[:index], slice[index+1:]...)
	return slice
}

// Structs
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

type Track struct {
	distance int
}

type Employee struct {
	firstName      string
	lastName       string
	salary         int
	jobDuration    bool
	yearsOfService string
}

func NewTrack(trackDistance int) Track {
	return Track{distance: trackDistance}
}

func NewCar(carSpeed int, batteryPerctge int) Car {
	return Car{
		speed:        carSpeed,
		batteryDrain: batteryPerctge,
		battery:      100,
	}
}

func NewEmployee(initName, midName string, salary int) Employee {
	return Employee{
		firstName:      initName,
		lastName:       midName,
		salary:         salary,
		jobDuration:    true,
		yearsOfService: "2024-09-18",
	}
}

func Drive(car Car) Car {
	distance := car.speed
	if car.batteryDrain == 0 {
		return NewCar(0, 0)
	}
	return Car{
		speed:        car.speed,
		batteryDrain: car.batteryDrain,
		battery:      100 - car.batteryDrain,
		distance:     distance,
	}
}

// CanFinish checks if a car is able to finish a certain track.
// func CanFinish(car Car, track Track) bool {
// }

func RollADie() int {
	return rand.Intn(20) + 1
}

func GenerateWandEnergy() float64 {
	return rand.Float64() + 12
}

func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})
	return animals
}

func TotalBirdCount(birdCount []int) int{
	var totalBirdCount int
	for _, v := range birdCount {
		totalBirdCount += v
	}
	return totalBirdCount
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	startIndex := (week - 1) * 7
	endIndex := startIndex + 7
	return TotalBirdCount(birdsPerDay[startIndex:endIndex])

}

func FixBirdCount(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] = birdsPerDay[i] + 1
	}
	return birdsPerDay
}

func main() {
	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	fmt.Println(TotalBirdCount(birdsPerDay))
	fmt.Println("Birds Per Specified Week is: ",BirdsInWeek(birdsPerDay, 2))
	// Declaration of anonymous structs
	installments := struct {
		userName, referredName string
		salary                 int
		jobType                bool
	}{
		userName:     "Monica",
		referredName: "Seyram",
		salary:       1200008,
		jobType:      true,
	}

	newEmployee := &Employee{
		firstName:   "Ross",
		lastName:    "Bing",
		salary:      13000,
		jobDuration: true,
	}

	car := NewCar(5, 2)
	track := NewTrack(12)
	xElement := []string{"a", "b", "c", "d", "e"}
	rand.Shuffle(len(xElement), func(i, j int) {
		xElement[i], xElement[j] = xElement[j], xElement[i]
	})

	fmt.Println(installments)
	fmt.Println(newEmployee)
	fmt.Println(newEmployee.firstName)
	fmt.Println(NewEmployee("Ernest", "Darko", 45600))
	fmt.Println(car)
	fmt.Println(track)
	fmt.Println(NewTrack(34))
	fmt.Println("Roll A die with THIS:", RollADie())
	fmt.Println("Wand Energy die with THIS:", GenerateWandEnergy())

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
	fmt.Println(NeedsLicense("car"))
	fmt.Println(HappyBirthday("Ernest", 34))
	fmt.Println(AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))

	fmt.Println(ChooseVehicle("Amanda Hongguang", "Amanda Corolla"))
	fmt.Println(CalculateResellPrice(1000, 1))
	slice := []int{56, 67, 45, 90, 109}
	findNumAtIndex(45, slice...)
	fmt.Println(GetItem([]int{5, 2, 10, 6, 8, 7, 0, 9}, 8))
	fmt.Println(SetItem([]int{1, 2, 4, 1}, -1, 6))
	fmt.Println(PrependItems([]int{3, 2, 6, 4, 8}))
	fmt.Println(RemoveItem([]int{3, 2, 6, 4, 8}, 2))
}
