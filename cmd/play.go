package cmd

import (
	"cli-blackjack/database"
	"cli-blackjack/game"
	"errors"
	"github.com/manifoldco/promptui"
	"strconv"
)

var playOpt = Opt{
	Name: "Play",
	Run: func() (bool, error) {
		bet, err := betPrompt()
		if err != nil {
			return false, err
		}

		result, err := game.Play(bet)
		if err != nil {
			return false, err
		}

		database.CurrentStats.Balance += result

		if result > 0 {
			database.CurrentStats.Wins++
		} else if result < 0 {
			database.CurrentStats.Losses++
		} else {
			database.CurrentStats.Draws++
		}

		if err := database.SaveStats(); err != nil {
			return false, err
		}

		return true, nil
	},
}

func betPrompt() (int, error) {
	max := database.CurrentStats.Balance

	prompt := promptui.Prompt{
		Label: "Bet",
		Validate: func(s string) error {
			value, err := strconv.Atoi(s)
			if err != nil {
				return errors.New("invalid number")
			} else if value <= 0 {
				return errors.New("value must be greater than 0")
			} else if value > max {
				return errors.New("value cannot be larger than full balance")
			}

			return nil
		},
	}

	betString, err := prompt.Run()
	if err != nil {
		return 0, err
	}

	bet, err := strconv.Atoi(betString)
	if err != nil {
		return 0, err
	}

	return bet, nil
}
