package bot

import "fmt"

import "irc"

type Bot struct{
    port uint16
    host string
    nick string
    name string
    ident string
    realname string
    owner string
    channel string
    ircc irc.IRCClient
}

func NewBot(port uint16, host, nick, name, ident, realname, owner, channel string) *Bot {
    var b Bot
    b.port = port
    b.host = host
    b.nick = nick
    b.name = name
    b.ident = ident
    b.realname = realname
    b.owner = owner
    b.channel = channel
    return &b
}

func (b *Bot) Run() {
    b.connect()
}

func (b *Bot) String() string {
    return fmt.Sprintf("(%s:%d %s %s %s %s %s %s)", b.host, b.port, b.nick, b.name, b.ident, b.realname, b.owner, b.channel)
}

func (b *Bot) connect() {
    ircc := irc.NewIRCClient(b.port, b.host, b.nick, b.name, b.ident, b.realname, b.owner, b.channel)
    ircc.MainLoop()
}

