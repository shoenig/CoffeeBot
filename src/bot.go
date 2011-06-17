package bot

import "fmt"

import "irc"

type Bot struct{
    ircc irc.IRCClient
}

func NewBot(port uint16, host, nick, name, ident, realname, owner, channel string) *Bot {
    var b Bot
    //ircc := irc.NewIRCClient(port, host, nick, name, ident, realname, owner, channel)
    return &b
}

func (b *Bot) SayHi() {
    fmt.Printf("Hi\n")
}


