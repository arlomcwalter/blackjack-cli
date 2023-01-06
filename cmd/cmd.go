package cmd

import (
	"errors"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type Opt struct {
	Name string
	Run  func() (bool, error)
}

var rootCmd = &cobra.Command{
	Use:   "blackjack",
	Short: "A simple blackjack game.",
}

var options = []Opt{playOpt, resetOpt, statsOpt, quitOpt}

func Execute() {
	optionNames := make([]string, len(options))
	for i, opt := range options {
		optionNames[i] = opt.Name
	}

	repeat := true

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label:    "Action",
			Items:    optionNames,
			HideHelp: true,
		}

		_, result, err := prompt.Run()
		if err != nil {
			Quit("Error creating option selector prompt.")
		}

		option, err := getOption(result)
		if err != nil {
			Quit("Invalid option selected.")
		}

		repeat, err = option.Run()
		if err != nil {
			Quit("Error running option.")
		}
	}

	for {
		if err := rootCmd.Execute(); err != nil {
			Quit("Error executing command.")
		}

		if !repeat {
			break
		}
	}
}

func getOption(name string) (Opt, error) {
	for _, opt := range options {
		if opt.Name == name {
			return opt, nil
		}
	}

	return Opt{}, errors.New("invalid option")
}
