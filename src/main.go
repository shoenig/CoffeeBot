package main

import "flag"
import . "fmt"
import "os/user"
import "syscall"

import "bot"

func main() {
	var config *string = flag.String("config", "", "Location of configuration file.")
	var def *bool = flag.Bool("default", false, "Enable to create a default config file")
	var daemonize *bool = flag.Bool("daemon", false, "Run in daemon mode")
	flag.Parse()
	if *def {
		bot.CreateDefaultConfig()
		return
	}
	if *daemonize {
		cbotU, err := user.Lookup("cbot")
		if err != nil {
			Printf("ERROR, No user cbot found, maybe you forgot to create it?")
			return
		}
		syscall.Setuid(cbotU.Uid) // drop root
	}
	ircbot := bot.NewBot(bot.ReadJSONConfig(*config))
	ircbot.Run()
}
