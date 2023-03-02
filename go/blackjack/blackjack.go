package blackjack

import "fmt"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "king":
		return 10
	case "queen":
		return 10
	case "jack":
		return 10
	case "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	case "one":
		return 1
	default:
		return 0

	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	c1 := ParseCard(card1)
	c2 := ParseCard(card2)
	dealer := ParseCard(dealerCard)
	sum := c1 + c2
	fmt.Printf("Player has %d. Dealer has %d\n", sum, dealer)

	if c1 == 11 && c2 == 11 {
		return "P"
	} else if sum == 21 && dealer < 10 {
		return "W"
	} else if sum >= 17 {
		return "S"
	} else if sum >= 12 && dealer < 7 {
		return "S"
	} else {
		return "H"
	}
}
