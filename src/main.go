package main

import "flag"
import "fmt"

import "bot"
import "utils"

func main() {
	var config *string = flag.String("config", "", "Location of configuration file.")
	var def *bool = flag.Bool("default", false, "Enable to create a default config file")
	flag.Parse()
	if *def {
		bot.CreateDefaultConfig()
		return
	}
	fmt.Printf("CoffeeBot 2011 v%s\n", utils.VERSION)
	ircbot := bot.NewBot(bot.ReadJSONConfig(*config))
	ircbot.Run()
}
