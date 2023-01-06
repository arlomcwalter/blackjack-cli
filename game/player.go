package game

import "fmt"

type Player struct {
	name   string
	hand   Hand
	stand  bool
	bust   bool
	dealer bool
}

func newPlayer(dealer bool) Player {
	var name string
	if dealer {
		name = "Dealer"
	} else {
		name = "Player"
	}

	return Player{
		name:   name,
		hand:   newHand(),
		stand:  false,
		bust:   false,
		dealer: dealer,
	}
}

func (p Player) handString(hide bool) string {
	return fmt.Sprintf("%s: %s[%s]", p.name, p.hand.cardString(hide), p.hand.valueString(hide))
}

func (h Hand) finalValue() int {
	value1, value2 := h.getValues(false)
	if value2 != -1 && value2 <= 21 {
		return value2
	}
	return value1
}

func (p Player) finalScore() int {
	value1, value2 := p.hand.getValues(false)
	if value2 != -1 && value2 <= 21 {
		return value2
	}
	return value1
}
