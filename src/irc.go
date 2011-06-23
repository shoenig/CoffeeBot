package irc

import "bufio"
import "fmt"
import "net"
import "rand"
import "strings"
import "time"

import "utils"

// FreeNode ports: 6665, 6666, 6667, 8000, 8001, 8002
// FreeNode SSL ports: 6697  7000 7070


type IRCClient struct {
	port       uint16
	host       string
	nick       string
	name       string
	ident      string
	realname   string
	owner      string
	channel    string
	recording  []string
	conn       net.Conn
	ogmHandler chan []byte
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
    rand.Seed(time.Nanoseconds())
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
		c.handleMessage(string(line))
	}
}

//TODO: rearchitect this crap. Handle messages in thread 0. Handle sending/outgoing messages in thread 1.
// Other threads can be there too, counting pings or days or whatever. Everybody just
// sends NEW outgoing messages to thread 1

func (c *IRCClient) handleMessage(line string) {
	inmess := NewIncomingMessage(line)
	if inmess == nil { // happens with empty messages/ invalid cmds, etc
		if strings.Contains(line, "/MOTD") {
			println("ghetto join")
			c.sendJoin()
		}
		return
	}
	fmt.Printf(">%s", inmess)
	if inmess.PureCmd() == "PING" {
		c.sendPong()
	} else if strings.Contains(inmess.Arg(), "/MOTD") {
		c.sendJoin()
	} else if inmess.PureCmd() == "PART" || inmess.PureCmd() == "QUIT" {
		c.thankLeave(inmess.Prefix())
	} else if strings.Contains(inmess.Arg(), "speak") {
		c.speak()
	} else if inmess.PureCmd() == "KICK" {
		if strings.Contains(inmess.Arg(), c.nick) {
			c.sendJoin()
		}
	}
}

func (c *IRCClient) sendPong() {
	fmt.Printf("< sending PONG\n")
	c.ogmHandler <- NewOutgoingMessage("", "PONG", c.host)
	fmt.Printf("< sending message about pong\n")

    v := rand.Float64()
    str := fmt.Sprintf("%v", v)
    if v > .5 {
        str += " hello"
    } else {
        str += " hi"
    }
	c.ogmHandler <- NewOutgoingMessage("", "PRIVMSG "+c.channel, str)
}

func (c *IRCClient) sendJoin() {
	fmt.Printf("< sending JOIN\n")
	c.ogmHandler <- NewOutgoingMessage("", "JOIN "+c.channel, "")
}

func (c *IRCClient) speak() {
	fmt.Printf("< speaking\n")
	c.ogmHandler <- NewOutgoingMessage("", "PRIVMSG "+c.channel, "OKAY")
}

func (c *IRCClient) thankLeave(prefix string) {
	fmt.Printf("< thankful leaving\n")
	exc := strings.IndexAny(prefix, "!")
	person := ""
	if exc != -1 {
		person = prefix[0:exc]
	} else {
		person = prefix
	}
	c.ogmHandler <- NewOutgoingMessage("", "PRIVMSG "+c.channel,
		"YAY! "+person+"  has left")
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
	c.ogmHandler = make(chan []byte)
	c.conn.SetTimeout(utils.SecsToNSecs(600))
	go handleOutgoingMessages(c.conn, c.ogmHandler)
}

func handleOutgoingMessages(conn net.Conn, input chan []byte) {
	for {
		ogm := <-input
		conn.Write(ogm)
	}
}
