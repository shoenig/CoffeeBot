package irc

import "fmt"
import "strings"

type IRCMessage struct {
	prefix   string
	command  string
	argument string
	cmds     map[string]bool
	ncmds    map[string]bool
}

// returns nil for empty or invalid messages
func NewIncomingMessage(line string) *IRCMessage {
	var m IRCMessage
	m.cmds = makeCmdMap()
	m.ncmds = makeNCmdMap()
	line = strings.TrimSpace(line)
	if len(line) == 0 { // empty message
		println("fuckitall")
		return nil
	}
	m.prefix = ""
	if line[0] == ':' { // we have optional prefix
		fs := strings.IndexAny(line, " ")
		m.prefix = line[1:fs]
		line = line[fs+1:]
	}
	//m.command = strings.Fields(line)[0]
	fc := strings.IndexAny(line, ":")
	if fc != -1 {
		m.command = strings.TrimSpace(line[0:fc])
		m.argument = line[fc+1:]
	} else {
		m.command = line
	}
	if !m.cmds[pureCmd(m.command)] { // not in literal cmds
		if !m.ncmds[pureCmd(m.command)] { // not in numerical cmds
			fmt.Printf("~>%s\n", m.command)
			fmt.Printf("\t%s\n", m.argument)
			return nil
		}
	}
	return &m
}

func NewOutgoingMessage(prefix, cmd, cmdarg, argument string) []byte {
	var outgoing string
    if cmdarg != "" {
        cmdarg = " " + cmdarg
    }
	if argument != "" {
		argument = " :" + argument
	}
	if len(prefix) == 0 {
		outgoing = fmt.Sprintf("%s%s%s\r\n", cmd, cmdarg, argument)
	} else {
		outgoing = fmt.Sprintf(":%s %s%s%s\r\n", prefix, cmd, cmdarg, argument)
	}
	return []byte(outgoing)
}

func (m *IRCMessage) Arg() string {
	return m.argument
}

func (m *IRCMessage) FullCmd() string {
	return m.command
}

func (m *IRCMessage) PureCmd() string {
	return pureCmd(m.FullCmd())
}

func (m *IRCMessage) Prefix() string {
	return m.prefix
}

func (m *IRCMessage) Eq(o *IRCMessage) bool {
	return (m.prefix == o.prefix &&
		m.command == o.command &&
		m.argument == o.argument)
}

func (m *IRCMessage) String() string {
	return string(NewOutgoingMessage(m.prefix, m.command, "", m.argument))
}

func pureCmd(cmd string) string {
	return strings.Fields(cmd)[0]
}


// map (as aset) of possible NUMERICAL commands
func makeNCmdMap() map[string]bool {
	var m = map[string]bool{
		"353": true,
	}
	return m
}

// map (as a set) of possible LITERAL commands
func makeCmdMap() map[string]bool {
	var m = map[string]bool{
		"ADMIN":    true,
		"AWAY":     true,
		"CONNECT":  true,
		"DIE":      true,
		"ERROR":    true,
		"INFO":     true,
		"INVITE":   true,
		"ISON":     true,
		"JOIN":     true,
		"KICK":     true,
		"KILL":     true,
		"LINKS":    true,
		"LIST":     true,
		"LUSERS":   true,
		"MODE":     true,
		"MOTD":     true,
		"NAMES":    true,
		"NICK":     true,
		"NOTICE":   true,
		"OPER":     true,
		"PART":     true,
		"PASS":     true,
		"PING":     true,
		"PONG":     true,
		"PRIVMSG":  true,
		"QUIT":     true,
		"REHASH":   true,
		"RESTART":  true,
		"SERVICE":  true,
		"SERVLIST": true,
		"SSERVER":  true,
		"SQUERY":   true,
		"SQUIT":    true,
		"STATS":    true,
		"SUMMON":   true,
		"TIME":     true,
		"TOPIC":    true,
		"TRACE":    true,
		"USER":     true,
		"USERHOST": true,
		"USERS":    true,
		"VERSION":  true,
		"WALLOPS":  true,
		"WHO":      true,
		"WHOIS":    true,
		"WHOWAS":   true,
	}
	return m
}
