package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
        case "ace" :
        	return 11
        case "jack", "queen" , "king" , "ten" :
        	return 10
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
        default:
        	return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    totalCard := ParseCard(card1) + ParseCard(card2)
    dCardVal := ParseCard(dealerCard)
	switch {
        case totalCard == 22:
        	return "P"
        case totalCard == 21 && dCardVal < 10:
            return "W"
        case (totalCard <= 20 && totalCard >=12 && dCardVal <7) || (totalCard>16 && dCardVal >= 10):
        	return "S"
		case (totalCard <= 16 && totalCard >=12 && dCardVal >= 7 ) || totalCard <=11 :
        	return "H"
        default:
        	return "error"
    }
}
