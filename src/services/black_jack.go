package main

func ParseCard(card string) int {
	var cards = map[string]int{
		"ace":   11,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
	}
	return cards[card]
}

func FirstTurn(card1, card2, dealerCard string) string {
	var response string
	dealerCardVal := ParseCard(dealerCard)
	sumOfPlayerCard := ParseCard(card1) + ParseCard(card2)
	if ParseCard(card1) == 11 && ParseCard(card2) == 11 {
		response = "P"
	}
	switch {
	case sumOfPlayerCard == 21 && (dealerCardVal != 11 && dealerCardVal != 10):
		response = "W"
	case (sumOfPlayerCard == 21) && (dealerCardVal == 11 || dealerCardVal == 10):
		response = "S"
	case (sumOfPlayerCard >= 17 && sumOfPlayerCard <= 20):
		response = "S"
	case (sumOfPlayerCard >= 12 && sumOfPlayerCard <= 16) && (dealerCardVal < 7):
		response = "S"
	case (sumOfPlayerCard >= 12 && sumOfPlayerCard <= 16) && dealerCardVal >= 7:
		response = "H"
	case sumOfPlayerCard <= 11:
		response = "H"
	}
	return response
}

func AFirstTurn(card1, card2, dealerCard string) string {
	var response string
	var dealerCardValue int = ParseCard(dealerCard)
	var myCardSum int = ParseCard(card1) + ParseCard(card2)

	switch {
	case ParseCard(card1) == 11 && ParseCard(card2) == 11:
		response = "P"
	case myCardSum == 21 && (dealerCardValue != 11 && dealerCardValue != 10):
		response = "W"
	case (myCardSum == 21) && (dealerCardValue == 11 || dealerCardValue == 10):
		response = "S"
	case (myCardSum >= 17 && myCardSum <= 20):
		response = "S"
	case (myCardSum >= 12 && myCardSum <= 16) && (dealerCardValue < 7):
		response = "S"
	case (myCardSum >= 12 && myCardSum <= 16) && (dealerCardValue >= 7):
		response = "H"
	case (myCardSum <= 11):
		response = "H"
	}
	return response
}
