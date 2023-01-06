package game

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Suit string

const (
	Spade   Suit = "spade"
	Heart   Suit = "heart"
	Club    Suit = "club"
	Diamond Suit = "diamond"
)

var suits = []Suit{Spade, Heart, Club, Diamond}

type Card struct {
	Suit  Suit
	Value int
}

func newCard() Card {
	suit := suits[rand.Intn(len(suits))]
	value := rand.Intn(13) + 1
	return Card{suit, value}
}

func (c Card) String() string {
	var value string
	switch c.Value {
	case 1:
		value = "A"
	case 11:
		value = "J"
	case 12:
		value = "Q"
	case 13:
		value = "K"
	default:
		value = strconv.Itoa(c.Value)
	}

	var suit string
	switch c.Suit {
	case Spade:
		suit = "♠"
	case Heart:
		suit = "♥"
	case Club:
		suit = "♣"
	case Diamond:
		suit = "♦"
	default:
		suit = " "
	}

	return fmt.Sprintf("%s%s", value, suit)
}

func (c Card) GetValue() int {
	if c.Value > 10 {
		return 10
	}
	return c.Value
}
