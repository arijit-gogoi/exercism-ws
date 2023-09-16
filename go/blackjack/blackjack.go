package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	case "ace":
		return 11
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	first := ParseCard(card1)
	second := ParseCard(card2)
	dealer := ParseCard(dealerCard)

	sum := first + second

	switch {
	case sum == 22:
		return "P"

	case sum == 21:
		if dealer != 10 && dealer != 11 {
			return "W"
		}
		return "S"

	case 17 <= sum && sum <= 21:
		return "S"

	case 12 <= sum && sum <= 16:
		if dealer >= 7 {
			return "H"
		}
		return "S"

	case sum <= 11:
		return "H"
	}
	return "X"
}
