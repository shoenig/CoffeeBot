package irc

import "fmt"
import "strings"
import "strconv"
import "time"

import "utils"

func (c *IRCClient) showHelp() {
	c.ogmHandler <- NOM("", "PRIVMSG", c.channel, "cmds: !coffee, !help, !about, !8ball, !weather, !uptime, !wiki")
}

func (c *IRCClient) showAbout() {
	mess1 := fmt.Sprintf("CoffeeBot v%s, A Bot for Coffee", utils.VERSION)
	mess2 := "Seth Hoenig, June 2011"
	mess3 := "Source code: https://github.com/Queue29/CoffeeBot"
	c.ogmHandler <- NOM("", "PRIVMSG", c.channel, mess1)
	c.ogmHandler <- NOM("", "PRIVMSG", c.channel, mess2)
	c.ogmHandler <- NOM("", "PRIVMSG", c.channel, mess3)
}

func (c *IRCClient) showTitle(arg string) {
	splitted := strings.Fields(arg)
	for _, w := range splitted {
		if strings.Contains(w, "http://") || strings.Contains(w, "www.") {
			title := utils.GetTitle(c.logger, w)
			if title != "" {
				c.ogmHandler <- NOM("", "PRIVMSG", c.channel, title)
			}
		}
	}
}

func (c *IRCClient) searchWiki(term string) {
	mess := utils.Wikify(term)
	c.ogmHandler <- NOM("", "PRIVMSG", c.channel, mess)
}

func (c *IRCClient) do8Ball() {
	opts := ops_8ball()
	choice := utils.RandInt(0, len(opts))
	reply := opts[choice]
	c.ogmHandler <- NOM("", "PRIVMSG", c.channel, reply)
}

func (c *IRCClient) coffeeTime() {
	c.logger.Printf("< it's coffee time\n")
	c.pushNickList = true
	st := utils.SimpleTime()
	cfe := "\u2615\u2615\u2615\u2615\u2615"
	c.ogmHandler <- NOM("", "PRIVMSG", c.channel, cfe+" COFFEE TIME! "+cfe+" "+st)
	c.ogmHandler <- NOM("", "NAMES", c.channel, "") //this will trip the icmh
}

func (c *IRCClient) doCoffeePSA(arg string) {
	if !c.pushNickList {
		return
	}
	c.logger.Printf("< Coffee PSA\n")
	c.logger.Printf("<< arg: %s\n", arg)
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
	tP := time.Seconds()
	uptimeSecs := tP - c.t0
	mess := utils.SecsToTime(uptimeSecs)
	c.ogmHandler <- NOM("", "PRIVMSG", c.channel, mess)
}

func (c *IRCClient) sendJoin() {
	c.logger.Printf("< sending JOIN\n")
	c.ogmHandler <- NOM("", "JOIN", c.channel, "")
}

func (c *IRCClient) postWeather(zipcode string) {
	_, err := strconv.Atoi(zipcode)
	if err != nil {
		c.logger.Printf("\tinvalid zipcode: %s\n", zipcode)
		c.ogmHandler <- NOM("", "PRIVMSG", c.channel, "usage: !weather <zipcode>")
	} else {
		weatherReport := utils.GetWeather(c.logger, zipcode)
		if weatherReport == "" {
			c.ogmHandler <- NOM("", "PRIVMSG", c.channel, "Couldn't reach the weather service")
		} else {
			c.ogmHandler <- NOM("", "PRIVMSG", c.channel, weatherReport)
		}
	}
}

func ops_8ball() []string {
	return []string{"yes", "no", "maybe", "i dunno", "perhaps", "unlikely", "you wish",
		"absolutely", "lol", "oh yea shake me harder baby", "dude leave me alone",
		"do i look like an 8 ball to you?", "hell no", "hell yea!", "don't waste my time",
		"you really need to ask?", "my magic 8 ball says yes", "who cares?",
	}
}
