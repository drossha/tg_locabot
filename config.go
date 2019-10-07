package main

import "os"

const (
	ApiUrl = "https://api.telegram.org/bot"
)

type Config struct {
	Token string
}

func LoadConfig() Config {
	var config = Config{
		Token: os.Getenv("TOKEN"),
	}

	return config
}
