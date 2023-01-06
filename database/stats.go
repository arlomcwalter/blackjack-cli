package database

import (
	"strconv"
)

type Stats struct {
	Balance int
	Wins    int
	Draws   int
	Losses  int
}

var CurrentStats *Stats
var DefaultBalance = 5000

func FetchStats() {
	balance := getField("balance", DefaultBalance)
	wins := getField("wins", 0)
	draws := getField("draws", 0)
	losses := getField("losses", 0)

	CurrentStats = &Stats{
		Balance: balance,
		Wins:    wins,
		Draws:   draws,
		Losses:  losses,
	}
}

func SaveStats() error {
	err := setField("balance", CurrentStats.Balance)
	if err != nil {
		return err
	}
	err = setField("wins", CurrentStats.Wins)
	if err != nil {
		return err
	}
	err = setField("draws", CurrentStats.Draws)
	if err != nil {
		return err
	}
	err = setField("losses", CurrentStats.Losses)
	if err != nil {
		return err
	}
	return nil
}

func getField(key string, fallback int) int {
	bytes, err := DB.Get([]byte(key), nil)
	if err != nil {
		return 0
	}

	value, err := strconv.Atoi(string(bytes))
	if err != nil {
		return fallback
	}

	return value
}

func setField(key string, value int) error {
	return DB.Put([]byte(key), []byte(strconv.Itoa(value)), nil)
}
