package bot

import "fmt"

import "irc"

type Bot struct {
	ircconf *irc.IRCConfig
	ircc    irc.IRCClient
}

func NewBot(conf *irc.IRCConfig) *Bot {
	var b Bot
	b.ircconf = conf
	return &b
}

func (b *Bot) Run() {
	b.connect()
}

func (b *Bot) String() string {
	return fmt.Sprintf("(%s:%d %s %s %s %s %s)",
		b.ircconf.Host,
		b.ircconf.Port,
		b.ircconf.Nick,
		b.ircconf.Ident,
		b.ircconf.Realname,
		b.ircconf.Owner,
		b.ircconf.Channel,
		b.ircconf.Password)
}

func (b *Bot) connect() {
	ircc := irc.NewIRCClient(b.ircconf)
	ircc.MainLoop()
}
