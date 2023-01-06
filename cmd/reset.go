package cmd

import (
	"cli-blackjack/database"
	"fmt"
	"github.com/manifoldco/promptui"
)

var resetOpt = Opt{
	Name: "Reset",
	Run: func() (bool, error) {
		prompt := promptui.Prompt{
			Label:     "Are you sure",
			IsConfirm: true,
		}

		_, err := prompt.Run()
		if err != nil {
			fmt.Println(Red("✗ Aborted reset operation."))
			return true, nil
		}

		stats := database.CurrentStats

		stats.Balance = database.DefaultBalance
		stats.Wins = 0
		stats.Draws = 0
		stats.Losses = 0

		fmt.Println(Green("✔ Reset all play history."))

		return true, nil
	},
}
