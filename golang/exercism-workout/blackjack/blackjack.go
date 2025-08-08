package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch {
	case card == "ace":
		return 11
	case card == "king", card == "queen", card == "jack", card == "ten":
		return 10
	case card == "two":
		return 2
	case card == "three":
		return 3

	case card == "four":
		return 4

	case card == "five":
		return 5

	case card == "six":
		return 6

	case card == "seven":
		return 7

	case card == "eight":
		return 8

	case card == "nine":
		return 9

	default:
		return 0
	}

}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"

	case ParseCard(card1)+ParseCard(card2) == 21:
		if ParseCard(dealerCard) >= 10 {
			return "S"
		} else {
			return "W"
		}
	case ParseCard(card1)+ParseCard(card2) >= 17 && ParseCard(card1)+ParseCard(card2) <= 20:
		return "S"

	case ParseCard(card1)+ParseCard(card2) >= 12 && ParseCard(card1)+ParseCard(card2) <= 16:
		if ParseCard(dealerCard) >= 7 {
			return "H"
		} else {
			return "S"
		}
	default:
		return "H"
	}
}
