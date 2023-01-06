package cmd

var quitOpt = Opt{
	Name: "Quit",
	Run: func() (bool, error) {
		return false, nil
	},
}
