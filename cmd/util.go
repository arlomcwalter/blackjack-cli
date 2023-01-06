package cmd

import (
	"cli-blackjack/database"
	"github.com/manifoldco/promptui"
	"log"
)

var (
	Bold  = promptui.Styler(promptui.FGBold)
	Faint = promptui.Styler(promptui.FGFaint)
	Green = promptui.Styler(promptui.FGGreen)
	Red   = promptui.Styler(promptui.FGRed)
	Cyan  = promptui.Styler(promptui.FGCyan)
)

func Quit(err string) {
	database.SaveStats()
	log.Fatal(Red(err))
}

func getConfirm() bool {
	prompt := promptui.Prompt{
		Label:     "Are you sure",
		IsConfirm: true,
	}

	_, err := prompt.Run()
	if err != nil {
		return false
	}

	return true
}
