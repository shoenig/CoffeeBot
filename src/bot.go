:package bot

import "fmt"
import "json"
import "strconv"
import "os"

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

func ReadJSONConfig(fName string) *irc.IRCConfig {
	f, ferr := os.Open(fName)
	if ferr != nil {
		panic("File Open Error")
	}
	d := json.NewDecoder(f)
	m := make(map[string]string)
	err := d.Decode(&m)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic("JSON Decode Error")
	}

	tp, _ := strconv.Atoi(m["Port"])
	port := uint16(tp)
	host := m["Host"]
	nick := m["Nick"]
	ident := m["Ident"]
	realname := m["Realname"]
	owner := m["Owner"]
	channel := m["Channel"]
	password := m["Password"]
	return &irc.IRCConfig{Port: port, Host: host, Nick: nick, Ident: ident,
		Realname: realname, Owner: owner, Channel: channel, Password: password}
}

func CreateDefaultConfig() {
	oFile, err := os.Create("config")
	if err != nil {
		panic("Could not create config file")
	}
	defer oFile.Close()
	_, err = oFile.Write([]byte("{\"Port\":\"6667\", \"Host\":\"chat.freenode.net\", \"Nick\":\"cbot\", \"Ident\":\"cbot\", \"Realname\":\"CBot\", \"Owner\":\"Your Name\", \"Channel\":\"#botwar\"}"))
	if err != nil {
		panic("Courld not write to config file")
	}
	fmt.Printf("Created default config file in ./config, please fill it in and restart bot\n")
	os.Exit(0)
}
