package cmd

import (
	"cli-blackjack/database"
	"fmt"
)

var statsOpt = Opt{
	Name: "Stats",
	Run: func() (bool, error) {
		stats := database.CurrentStats

		fmt.Printf("%s %s\n", Bold("Balance:"), Cyan(fmt.Sprintf("$%d", stats.Balance)))
		fmt.Printf("%s %s\n", Bold("Wins:"), Green(stats.Wins))
		fmt.Printf("%s %s\n", Bold("Draws:"), Faint(stats.Draws))
		fmt.Printf("%s %s\n", Bold("Losses:"), Red(stats.Losses))

		return true, nil
	},
}
