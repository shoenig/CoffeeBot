package main

import "fmt"

import "bot"
import "irc"

// FreeNode ports: 6665, 6666, 6667, 8000, 8001, 8002
// FreeNode SSL ports: 6697  7000 7070  

func main() {
    fmt.Printf("Hello World\n")

    ircbot := bot.NewBot("config")
    ircbot.SayHi()

    port := uint16(8000)
    host := "wolfe.freenode.net"
    nick := "coffeebot"
    name := "coffeebot"
    ident := "coffeebot"
    realname := "Coffee Bot"
    owner := "Seth Hoenig"
    ircc := irc.NewIRCClient(port, host, nick, name, ident, realname, owner, "#test")
    fmt.Printf("%v\n", ircc)
    ircc.PokeInternet()
}

