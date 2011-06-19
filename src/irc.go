package irc

import "bufio"
import "fmt"
import "net"
import "strings"

import "utils"

// FreeNode ports: 6665, 6666, 6667, 8000, 8001, 8002
// FreeNode SSL ports: 6697  7000 7070

//    ADMIN    AWAY   CONNECT DIE     ERROR  INFO   INVITE  ISON    JOIN    KICK
//    KILL     LINKS  LIST    LUSERS  MODE   MOTD   NAMES   NICK    NOTICE  OPER PART
//    PASS     PING   PONG    PRIVMSG QUIT   REHASH RESTART SERVICE SERVLIST
//    SSERVER  SQUERY SQUIT   STATS   SUMMON TIME   TOPIC   TRACE   USER
//    USERHOST USERS  VERSION WALLOPS WHO    WHOIS  WHOWAS

type IRCClient struct {
	port               uint16
	host               string
	nick               string
	name               string
	ident              string
	realname           string
	owner              string
	channel            string
	recording          []string
	conn               net.Conn
	messageHandlerChan chan string
}

func NewIRCClient(port uint16, host, nick, name, ident, realname, owner, channel string) *IRCClient {
	var c IRCClient
	c.setPort(port)
	c.setHost(host)
	c.setNick(nick)
	c.setName(name)
	c.setIdent(ident)
	c.setRealName(realname)
	c.setOwner(owner)
	c.setChannel(channel)
	c.messageHandlerChan = make(chan string)
	return &c
}

func (c *IRCClient) String() string {
	return fmt.Sprintf("(%d, %s, %s, %s, %s, %s, %s)",
		c.port, c.nick, c.name, c.ident, c.realname, c.owner, c.channel)
}

func (c *IRCClient) Port() uint16 { return c.port }
func (c *IRCClient) setPort(port uint16) {
	if port < 1024 {
		panic("Invalid Port")
	}
	c.port = port
}

func (c *IRCClient) Host() string { return c.host }
func (c *IRCClient) setHost(host string) {
	if host == "" {
		panic("Invalid Host")
	}
	c.host = host
}

func (c *IRCClient) Nick() string { return c.nick }
func (c *IRCClient) setNick(nick string) {
	if len(nick) == 0 || len(nick) > 9 {
		panic("Invalid Nick")
	}
	c.nick = nick
}

func (c *IRCClient) Name() string { return c.name }
func (c *IRCClient) setName(name string) {
	if name == "" {
		panic("Invalid Name")
	}
	c.name = name
}

func (c *IRCClient) Ident() string { return c.ident }
func (c *IRCClient) setIdent(ident string) {
	if ident == "" {
		panic("Invalid Ident")
	}
	c.ident = ident
}

func (c *IRCClient) RealName() string { return c.realname }
func (c *IRCClient) setRealName(realname string) {
	if realname == "" {
		panic("Invalid RealName")
	}
	c.realname = realname
}

func (c *IRCClient) Owner() string { return c.owner }
func (c *IRCClient) setOwner(owner string) {
	if owner == "" {
		panic("Invalid Owner")
	}
	c.owner = owner
}

func (c *IRCClient) Channel() string { return c.channel }
func (c *IRCClient) setChannel(channel string) {
	if len(channel) < 2 || channel[0] != '#' || len(channel) > 200 {
		panic("Invalid Channel")
	}
	c.channel = channel
}

func (c *IRCClient) MainLoop() {
	c.initializeConnection()
	buffr := bufio.NewReader(c.conn)
	for {
		line, _, rerr := buffr.ReadLine()
		if rerr != nil {
			fmt.Printf("rerr: %v\n", rerr)
			panic("ERROR")
		}
		c.messageHandlerChan <- string(line)
	}
}

func (c *IRCClient) handleMessages() {
	for {
		line := <-c.messageHandlerChan
		fmt.Printf("%s\n", line)
		if strings.Contains(line, "PING") {
			c.sendPong(line)
		} else if strings.Contains(line, "/MOTD") {
			c.sendJoin()
		} else if strings.Contains(line, "speak") {
			c.speak()
		}
	}
}

func (c *IRCClient) sendPong(line string) {
	fmt.Printf("> sending PONG\n")
	c.conn.Write([]byte("PONG " + c.host + "\r\n"))
}

func (c *IRCClient) sendJoin() {
	fmt.Printf("> sending JOIN\n")
	c.conn.Write([]byte("JOIN " + c.channel + "\r\n"))
}

func (c *IRCClient) speak() {
	fmt.Printf("> speaking\n")
	c.conn.Write([]byte("PRIVMSG " + c.channel + " :yes master\r\n"))
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
	if nickerr != nil {
		panic(fmt.Sprintf("NICK message error: %s", nickerr))
	}
	user_mess := []uint8("USER " + c.nick + " 0 * :" + c.realname + "\r\n")
	_, usererr := c.conn.Write(user_mess)
	if usererr != nil {
		panic(fmt.Sprintf("USER message err: %s", usererr))
	}
	c.conn.SetTimeout(utils.SecsToNSecs(600))
	go c.handleMessages()
}
