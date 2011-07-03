package irc

import "fmt"
import "strings"
import "time"

import "utils"

func (c *IRCClient) showHelp() {
	fmt.Printf("< showing help\n")
	c.ogmHandler <- NOM("", "PRIVMSG", c.channel, "cmds: !coffee, !help, !about, !8ball, !weather, !uptime, !wiki")
}

func (c *IRCClient) showAbout() {
	fmt.Printf("< showing about\n")
	mess1 := "CoffeeBot v0.1, A Bot for Coffee"
	mess2 := "Seth Hoenig, June 2011"
	mess3 := "Source code: https://github.com/Queue29/CoffeeBot"
	c.ogmHandler <- NOM("", "PRIVMSG", c.channel, mess1)
	c.ogmHandler <- NOM("", "PRIVMSG", c.channel, mess2)
	c.ogmHandler <- NOM("", "PRIVMSG", c.channel, mess3)
}

func (c *IRCClient) sendWeather() {
	fmt.Printf("sending weather " + utils.SimpleTime() + "\n")
	mess := "cloudy with a chance of meatballs"
	c.ogmHandler <- NOM("", "PRIVMSG", c.channel, mess)
}

func (c *IRCClient) searchWiki(term string) {
	fmt.Printf("searching wikipedia\n")
	mess := utils.Wikify(term)
	c.ogmHandler <- NOM("", "PRIVMSG", c.channel, mess)
}

func (c *IRCClient) do8Ball() {
	fmt.Printf("8 balling\n")
	opts := ops_8ball()
	choice := utils.RandInt(0, len(opts))
	reply := opts[choice]
	c.ogmHandler <- NOM("", "PRIVMSG", c.channel, reply)
}

func (c *IRCClient) coffeeTime() {
	fmt.Printf("< it's coffee time\n")
	c.pushNickList = true
	st := utils.SimpleTime()
	c.ogmHandler <- NOM("", "PRIVMSG", c.channel, "☕☕☕☕☕☕ COFFEE TIME! ☕☕☕☕☕☕ "+st)
	c.ogmHandler <- NOM("", "NAMES", c.channel, "") //this will trip the icmh
}

func (c *IRCClient) doCoffeePSA(arg string) {
	if !c.pushNickList {
		return
	}
	fmt.Printf("< Coffee PSA\n")
	fmt.Printf("<< arg: %s\n", arg)
	psa := ""
	nicks := strings.Fields(arg)
	for _, nick := range nicks {
		if nick[0] == '+' || nick[0] == '@' {
			nick = nick[1:]
			if nick == "ChanServ" {
				continue
			}
		}
		psa += (nick + " ")
	}
	c.pushNickList = false
	c.ogmHandler <- NOM("", "PRIVMSG", c.channel, psa)
}

func (c *IRCClient) sendUptime() {
	fmt.Printf("uptiming\n")
	tP := time.Seconds()
	uptimeSecs := tP - c.t0
	mess := utils.SecsToTime(uptimeSecs)
	c.ogmHandler <- NOM("", "PRIVMSG", c.channel, mess)
}

func (c *IRCClient) sendJoin() {
	fmt.Printf("< sending JOIN\n")
	c.ogmHandler <- NOM("", "JOIN", c.channel, "")
}

func (c *IRCClient) speak() {
	fmt.Printf("< speaking\n")
	c.ogmHandler <- NOM("", "PRIVMSG", c.channel, "OKAY")
}

func (c *IRCClient) postWeather() {
	fmt.Printf("< sending weather report\n")
	weatherReport := utils.GetWeather()
	if weatherReport == "" {
		c.ogmHandler <- NOM("", "PRIVMSG", c.channel, "Couldn't reach the weather service")
	}
	c.ogmHandler <- NOM("", "PRIVMSG", c.channel, weatherReport)
}

func (c *IRCClient) thankLeave(prefix string) {
	fmt.Printf("< thankful leaving\n")
	exc := strings.IndexAny(prefix, "!")
	person := ""
	if exc != -1 {
		person = prefix[0:exc]
	} else {
		person = prefix
	}
	if utils.RandInt(0, 1) == 0 {
		choice := utils.RandInt(0, len(ops_thankLeave()))
		reply := fmt.Sprintf(ops_thankLeave()[choice], person)
		c.ogmHandler <- NOM("", "PRIVMSG", c.channel, reply)
	}
}

func (c *IRCClient) doWelcome(prefix string) {
	fmt.Printf("< welcoming\n")
	exc := strings.IndexAny(prefix, "!")
	person := ""
	if exc != -1 {
		person = prefix[0:exc]
	} else {
		person = prefix
	}
	c.ogmHandler <- NOM("", "PRIVMSG", c.channel, "welome, "+person+"!")
}

func ops_8ball() []string {
	return []string{"yes", "no", "maybe", "i dunno", "perhaps", "unlikely", "you wish",
		"absolutely", "lol", "oh yea shake me harder baby", "dude leave me alone",
		"do i look like an 8 ball to you?", "hell no", "hell yea!", "don't waste my time",
		"you really need to ask?", "my magic 8 ball says yes", "who cares?",
	}
}

func ops_thankLeave() []string {
	return []string{"ugh finally, %s has left", "laterz, %s", "alright, %s is gone, we can party now",
		"bye %s", "time to move %s's desk to the patio", "%s has left the building",
	}
}
