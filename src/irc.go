package irc

import "fmt"
import "net"

const (
    ADMIN = iota
    AWAY
    CONNECT
    DIE
    EROR
    INFO
    INVITE
    ISON
    JOIN
    KICK
    KILL
    LINKS
    LIST
    LUSERS
    MODE
    MOTD
    NAMES
    NICK
    NOTICE
    OPER
    PART
    PASS
    PING
    PONG
    PRIVMSG
    QUIT
    REHASH
    RESTART
    SERVICE
    SERVLIST
    SSERVER
    SQUERY
    SQUIT
    STATS
    SUMMON
    TIME
    TOPIC
    TRACE
    USER
    USERHOST
    USERS
    VERSION
    WALLOPS
    WHO
    WHOIS
    WHOWAS
)

type IRCClient struct {
    port uint16
    nick string
    name string
    ident string
    realname string
    owner string
    channel string
    recording []string
}

func NewIRCClient(port uint16, nick, name, ident, realname, owner, channel string) *IRCClient {
    var c IRCClient
    c.SetPort(port)
    c.SetNick(nick)
    c.SetName(name)
    c.SetIdent(ident)
    c.SetRealName(realname)
    c.SetOwner(owner)
    c.SetChannel(channel)
    return &c
}

func (c *IRCClient) String() string {
    return fmt.Sprintf("(%d, %s, %s, %s, %s, %s, %s)",
                        c.port, c.nick, c.name, c.ident, c.realname, c.owner, c.channel)
}

func (c *IRCClient) Port() uint16 { return c.port }
func (c *IRCClient) SetPort(port uint16) {
    if port < 1024 { panic("Invalid Port") }
    c.port = port
    fmt.Printf("c.port: %d\n", c.port)
}

func (c *IRCClient) Nick() string { return c.nick }
func (c *IRCClient) SetNick(nick string) {
    if nick == "" { panic("Invalid Nick") }
    c.nick = nick
}

func (c *IRCClient) Name() string { return c.name }
func (c *IRCClient) SetName(name string) {
    if name == "" { panic("Invalid Name") }
    c.name = name
}

func (c *IRCClient) Ident() string { return c.ident}
func (c *IRCClient) SetIdent(ident string) {
    if ident == "" { panic("Invalid Ident") }
    c.ident = ident
}

func (c *IRCClient) RealName() string { return c.realname }
func (c *IRCClient) SetRealName(realname string) {
    if realname == "" { panic("Invalid RealName") }
    c.realname = realname
}

func (c *IRCClient) Owner() string { return c.owner }
func (c *IRCClient) SetOwner(owner string) {
    if owner == "" { panic("Invalid Owner") }
    c.owner = owner
}

func (c *IRCClient) Channel() string { return c.channel }
func (c *IRCClient) SetChannel(channel string) {
    if channel == "" { panic("Invalid Channel") }
    c.channel = channel
}

func (c *IRCClient) PokeInternet() {
    _, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "google.com", 80))
    if err != nil {
        fmt.Printf("Crash and BURN!\n")
        fmt.Printf("%v\n", err)
    } else {
        fmt.Printf("Success!\n")
    }
}
