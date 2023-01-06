package game

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"strconv"
)

type Game struct {
	player Player
	dealer Player
	bet    int
}

type move string

const (
	hit    move = "hit"
	stand  move = "stand"
	double move = "double"
)

func Play(bet int) (int, error) {
	game := Game{
		player: newPlayer(false),
		dealer: newPlayer(true),
		bet:    bet,
	}

	player := game.player
	dealer := game.dealer

	fmt.Println(player.handString(false))

	if player.hand.isBlackjack() {
		player.stand = true
		fmt.Println("You have a blackjack!")
	}

	if dealer.hand.isBlackjack() {
		dealer.stand = true
		fmt.Println(dealer.handString(false))
		fmt.Println("Dealer has blackjack.")
	} else {
		fmt.Println(dealer.handString(true))
	}

playerLoop:
	for !player.stand {
		playerMove, err := game.promptMoves()
		if err != nil {
			return 0, err
		}

		switch playerMove {
		case hit:
			player.hand = append(player.hand, newCard())
		case stand:
			player.stand = true
			break playerLoop
		case double:
			game.bet *= 2
			player.hand = append(player.hand, newCard())
			player.stand = true
		}

		fmt.Println(player.handString(false))

		if player.hand.isBlackjack() {
			player.stand = true
			fmt.Println("Blackjack!")
			break
		}

		if player.hand.isBust() {
			player.bust = true
			fmt.Println("Bust!")
			break
		}
	}

	for {
		playerScore := player.finalScore()
		dealerScore := dealer.finalScore()

		if player.bust {
			dealer.stand = true
			break
		}

		if playerScore > dealerScore {
			dealer.hand = append(dealer.hand, newCard())
			fmt.Println(dealer.handString(false))
		} else {
			dealer.stand = true
			break
		}

		if dealer.hand.isBlackjack() {
			dealer.stand = true
			fmt.Println("Dealer got a blackjack!")
			break
		}

		if dealer.hand.isBust() {
			dealer.bust = true
			fmt.Println("Dealer bust!")
			break
		}
	}

	playerScore := player.finalScore()
	dealerScore := dealer.finalScore()

	fmt.Println("Player: " + strconv.Itoa(playerScore))
	fmt.Println("Dealer: " + strconv.Itoa(dealerScore))

	if player.bust {
		fmt.Printf("You lost $%d.\n", game.bet)
		return -game.bet, nil
	}

	if playerScore > dealerScore || dealer.bust {
		if playerScore == 21 {
			game.bet *= 2
		}

		fmt.Printf("You won $%d!\n", game.bet)
		return game.bet, nil
	}

	if playerScore == dealerScore {
		fmt.Println("Push.")
		return 0, nil
	}

	fmt.Printf("You lost $%d.\n", game.bet)
	return -game.bet, nil
}

func (g Game) promptMoves() (move, error) {
	moves := []move{hit, stand}
	if len(g.player.hand) == 2 {
		moves = append(moves, double)
	}

	prompt := promptui.Select{
		Label:        "Move",
		Items:        moves,
		HideHelp:     true,
		HideSelected: true,
	}

	_, moveStr, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return move(moveStr), nil
}
