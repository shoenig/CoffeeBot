package irc

import "bufio"
import "crypto/tls"
import crand "crypto/rand"
import "fmt"
import "net"
import "rand"
import "strings"
import "time"

import "utils"

// FreeNode ports: 6665, 6666, 6667, 8000, 8001, 8002
// FreeNode SSL ports: 6697  7000 7070

type IRCConfig struct {
	Port     uint16
	Host     string
	Nick     string
	Ident    string
	Realname string
	Owner    string
	Channel  string
	Password string
}

type IRCClient struct {
	port         uint16
	host         string
	nick         string
	ident        string
	realname     string
	owner        string
	channel      string
	password     string
	recording    []string
	tlsc         *tls.Conn
	t0           int64
	pushNickList bool
	ogmHandler   chan []byte
}

func NewIRCClient(ircConfig *IRCConfig) *IRCClient {
	var c IRCClient
	c.setPort(ircConfig.Port)
	c.setHost(ircConfig.Host)
	c.setNick(ircConfig.Nick)
	c.setIdent(ircConfig.Ident)
	c.setRealName(ircConfig.Realname)
	c.setOwner(ircConfig.Owner)
	c.setChannel(ircConfig.Channel)
	c.setPassword(ircConfig.Password)
	rand.Seed(time.Nanoseconds())
	c.pushNickList = false
	return &c
}

func (c *IRCClient) String() string {
	return fmt.Sprintf("(%d, %s, %s, %s, %s, %s)",
		c.port, c.nick, c.ident, c.realname, c.owner, c.channel)
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

func (c *IRCClient) setPassword(password string) {
	c.password = password
}

func (c *IRCClient) MainLoop() {
	c.initializeConnection()
	buffr := bufio.NewReader(c.tlsc)
	for {
		line, _, rerr := buffr.ReadLine()
		if rerr != nil {
			fmt.Printf("rerr: %v\n", rerr)
			panic("ERROR")
		}
		c.handleMessage(string(line))
	}
}

func (c *IRCClient) handleMessage(line string) {
	inmess := NewIncomingMessage(line)

	if inmess == nil { // happens with empty messages/ invalid cmds, etc
		if strings.Contains(line, "/MOTD") {
			c.sendJoin()
		}
		return
	}
	fmt.Printf(">%s", inmess)
	if inmess.PureCmd() == "PING" {
		c.sendPong()
	} else if inmess.PureCmd() == "JOIN" {
		if strings.Contains(inmess.Arg(), "#yelp") {
            c.fixChannel()
        }
	} else if inmess.PureCmd() == "PART" || inmess.PureCmd() == "QUIT" {
		//c.thankLeave(inmess.Prefix())
	} else if inmess.PureCmd() == "KICK" {
		if strings.Contains(inmess.Arg(), c.nick) {
			c.sendJoin()
		}
	} else if inmess.PureCmd() == "353" {
		c.doCoffeePSA(inmess.Arg())
	} else if inmess.PureCmd() == "PRIVMSG" {
		/* most of everything will live in here */
		//TODO turn into switch statement
		if strings.Contains(inmess.Arg(), "!8ball") {
			c.do8Ball()
		} else if strings.Contains(inmess.Arg(), "!uptime") {
			c.sendUptime()
		} else if strings.Contains(inmess.Arg(), "!weather") {
			c.postWeather()
		} else if strings.Contains(inmess.Arg(), "!help") {
			c.showHelp()
		} else if strings.Contains(inmess.Arg(), "!coffee") || strings.Contains(inmess.Arg(), "!COFFEE") {
			c.coffeeTime()
		} else if strings.Contains(inmess.Arg(), "!about") {
			c.showAbout()
		} else if strings.Contains(inmess.Arg(), "!wiki") {
			c.searchWiki(inmess.Arg())
		}
	}
}

func (c *IRCClient) sendPong() {
	fmt.Printf("< sending PONG\n")
	c.ogmHandler <- NewOutgoingMessage("", "PONG", "", c.host)
}

func (c *IRCClient) fixChannel() {
    fmt.Printf("< leaving #yelp, moving to #" + c.channel)
    time.Sleep(utils.SecsToNSecs(2))
    c.ogmHandler <- NewOutgoingMessage("", "JOIN", c.channel, "")
    c.ogmHandler <- NewOutgoingMessage("", "PART", "#yelp", "Quit: Leaving.")
    time.Sleep(utils.SecsToNSecs(2))
    c.ogmHandler <- NewOutgoingMessage("", "NICK", c.nick, "")
}

func (c *IRCClient) initializeConnection() {
	c.t0 = time.Seconds()
	tconn, cerr := net.Dial("tcp", fmt.Sprintf("%s:%d", c.host, c.port))
	if cerr != nil {
		fmt.Printf("cerr: %v\n", cerr)
		panic("Error Connecting!!")
	}
	cf := &tls.Config{Rand: crand.Reader, Time: time.Nanoseconds}
	c.tlsc = tls.Client(tconn, cf)
	if c.password != "" {
		fmt.Printf("Sending Password: %s\n", c.password)
		pass_mes := []byte("PASS " + c.password + "\r\n")
		_, passerr := c.tlsc.Write(pass_mes)
		if passerr != nil {
			panic(fmt.Sprintf("PASS message error: %s", passerr))
		}
	}

	nick_mess := []byte("NICK " + c.nick + "\r\n")
	_, nickerr := c.tlsc.Write(nick_mess)
	if nickerr != nil {
		panic(fmt.Sprintf("NICK message error: %s", nickerr))
	}

	user_mess := []byte("USER " + c.nick + " 0 * :" + c.realname + "\r\n")
	_, usererr := c.tlsc.Write(user_mess)
	if usererr != nil {
		panic(fmt.Sprintf("USER message err: %s", usererr))
	}

	c.ogmHandler = make(chan []byte)
	tconn.SetTimeout(utils.SecsToNSecs(600))
	go handleOutgoingMessages(c.tlsc, c.ogmHandler)
}


/* Long Running Go-routines that handle long standing tasks */
func (c *IRCClient) randomHelloSender() {
	for {
		stime := (rand.Int() % 100) + 25
		time.Sleep(utils.SecsToNSecs(int64(stime)))
		v := rand.Float64()
		str := ""
		if v > .5 {
			str += " hello"
		} else {
			str += " hi"
		}
		c.ogmHandler <- NewOutgoingMessage("", "PRIVMSG", c.channel, str)
	}
}

func handleOutgoingMessages(conn net.Conn, input chan []byte) {
	for {
		ogm := <-input
		conn.Write(ogm)
	}
}
