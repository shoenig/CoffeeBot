package main

import "flag"

import "bot"

func main() {
	var config *string = flag.String("config", "", "Location of configuration file.")
	var def *bool = flag.Bool("default", false, "Enable to create a default config file")
	flag.Parse()
	if *def {
		bot.CreateDefaultConfig()
		return
	}
	ircbot := bot.NewBot(bot.ReadJSONConfig(*config))
	ircbot.Run()
}
