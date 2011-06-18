package irc

import "fmt"
import "net"
import "strings"

import "utils"

// FreeNode ports: 6665, 6666, 6667, 8000, 8001, 8002
// FreeNode SSL ports: 6697  7000 7070

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
    conn net.Conn
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

func (c *IRCClient) MainLoop() {
    c.initializeConnection()
    var rbuff string
    for {
        var buff = []byte("☈") // one byte at a time b/c i don't know what i am doing
        _, rerr := c.conn.Read(buff)
        rbuff += strings.Replace(string(buff), "☈", "", -1)
        if rerr != nil {
            fmt.Printf("rerr: %v\n", rerr)
            panic("ERROR")
        }
        temp := strings.Split(strings.TrimSpace(rbuff), "\n", -1)
        rbuff = temp[len(temp)-1]
        temp = temp[0:len(temp)-1] // pop
        for i:=0; i<len(temp); i++ {
            line := strings.TrimSpace(temp[i])
            if line == "" { continue }
            fmt.Printf("%s\n", line)
            sp := strings.Fields(line)
            for _, strn := range sp {
                if strings.Contains(strn, "PING") {
                    fmt.Printf("sending PONG...\n")
                    pongmess := []byte("PONG wolfe.freenode.net\r\n")
                    c.conn.Write(pongmess)
                }
            }
        }
    }
}

func (c *IRCClient) initializeConnection() {
    tconn, cerr := net.Dial("tcp", fmt.Sprintf("%s:%d", c.host, c.port))
    if cerr != nil {
        fmt.Printf("cerr: %v\n", cerr)
        panic("Error Connecting!!")
    }
    c.conn = tconn
    nick_mess := []uint8("NICK " + c.nick + "\r\n")
    _, nickerr := c.conn.Write(nick_mess)
    if nickerr != nil { panic(fmt.Sprintf("NICK message error: %s", nickerr)) }
    user_mess := []uint8("USER " + c.nick + " 0 * :" + c.realname + "\r\n")
    _, usererr := c.conn.Write(user_mess)
    if usererr != nil { panic(fmt.Sprintf("USER message err: %s", usererr)) }
    c.conn.SetTimeout(utils.SecsToNSecs(600))
}
