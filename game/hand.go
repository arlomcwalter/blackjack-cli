package game

import (
	"fmt"
	"strconv"
)

type Hand []Card

func newHand() Hand {
	return Hand{newCard(), newCard()}
}

func (h Hand) cardString(dealer bool) string {
	var cards string

	for i, card := range h {
		if i != 0 && dealer {
			cards += "??"
		} else {
			cards += card.String()
		}

		if i != len(h)-1 {
			cards += " "
		}
	}

	return cards
}

func (h Hand) valueString(dealer bool) string {
	if h.isBlackjack() {
		return "21"
	}

	value1, value2 := h.getValues(dealer)

	if value2 == -1 || value2 > 21 {
		return strconv.Itoa(value1)
	}

	return fmt.Sprintf("%d/%d", value1, value2)
}

func (h Hand) getValues(dealer bool) (int, int) {
	var value1, value2 int

	for i, card := range h {
		if i > 0 && dealer {
			break
		}

		if card.GetValue() == 1 {
			value1 += card.GetValue()
			value2 += 11
		} else {
			value1 += card.GetValue()
			value2 += card.GetValue()
		}
	}

	if value1 == value2 {
		value2 = -1
	}

	return value1, value2
}

func (h Hand) isBlackjack() bool {
	value1, value2 := h.getValues(false)
	return value1 == 21 || value2 == 21
}

func (h Hand) isBust() bool {
	value1, value2 := h.getValues(false)
	return value1 > 21 && (value2 == -1 || value2 > 21)
}
