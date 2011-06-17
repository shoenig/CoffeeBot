package irc

import "fmt"
import "io/ioutil"
import "net"

import "utils"

const (
    ADMIN = iota;     AWAY;    CONNECT; DIE;    ERROR;  INFO;    INVITE;  ISON;   JOIN; KICK
    KILL;     LINKS;  LIST;    LUSERS;  MODE;   MOTD;   NAMES;   NICK;    NOTICE; OPER; PART
    PASS;     PING;   PONG;    PRIVMSG; QUIT;   REHASH; RESTART; SERVICE; SERVLIST
    SSERVER;  SQUERY; SQUIT;   STATS;   SUMMON; TIME;   TOPIC;   TRACE;   USER
    USERHOST; USERS;  VERSION; WALLOPS; WHO;    WHOIS;  WHOWAS
)

type IRCClient struct {
    port uint16
    host string
    nick string
    name string
    ident string
    realname string
    owner string
    channel string
    recording []string
}

func NewIRCClient(port uint16, host, nick, name, ident, realname, owner, channel string) *IRCClient {
    var c IRCClient
    c.SetPort(port)
    c.SetHost(host)
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
}

func (c *IRCClient) Host() string { return c.host }
func (c *IRCClient) SetHost(host string) {
    if host=="" { panic("Invalid Host") }
    c.host = host
}

func (c *IRCClient) Nick() string { return c.nick }
func (c *IRCClient) SetNick(nick string) {
    if len(nick) == 0 || len(nick) >9 { panic("Invalid Nick") }
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
    fmt.Printf("%s\n", channel)
    if len(channel) < 2 || channel[0] != '#' || len(channel) > 200 { panic("Invalid Channel") }
    c.channel = channel
}

func (c *IRCClient) PokeInternet() {
    conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d",c.host, c.port))
    if err != nil {
        fmt.Printf("Crash and BURN!\n")
        fmt.Printf("%v\n", err)
    } else {
        fmt.Printf("Success!\n")
    }
    b0 := []uint8("NICK " + c.nick + "\r\n")
    fmt.Printf("b0: %v\n", string(b0))
    i, err := conn.Write(b0)
    if err != nil { fmt.Printf("ERROR, err: %v\n", err) }
    fmt.Printf("i: %d\n", i)

    b1 := []uint8("USER coffeebot 0 * :Seth\r\n")
    fmt.Printf("b1: %v\n", string(b1))
    j, err2 := conn.Write(b1)
    if err2 != nil { fmt.Printf("Error, err2: %v\n", err2) }
    fmt.Printf("j: %d\n", j)

    et := conn.SetTimeout(utils.SecsToNSecs(60)) // 60s
    if et != nil { panic("Error setting timeout") }

    buff, rerr := ioutil.ReadAll(conn)
    if rerr != nil {
        fmt.Printf("buff:\n%v\n", string(buff))
        fmt.Printf("ERROR: %v\n", rerr)
        panic("Reader Error")
    }
    fmt.Printf("Success in reading from server\n")
    fmt.Printf("buff:\n%v\n", string(buff))
}
