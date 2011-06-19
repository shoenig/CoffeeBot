package irc

import "fmt"
import "strings"

type IRCMessage struct {
	argument string
	command  string
	prefix   string
	cmds     map[string]bool
}

// returns nil for empty or invalid messages
func NewIncomingMessage(line string) *IRCMessage {
	var m IRCMessage
	m.cmds = makeCmdMap()
	line = strings.TrimSpace(line)
	if len(line) == 0 { // empty message
		return nil
	}
	m.prefix = ""
	if line[0] == ':' { // we have optional prefix
		fs := strings.IndexAny(line, " ")
		m.prefix = line[1:fs]
		line = line[fs+1:]
	}
	m.command = strings.Fields(line)[0]
	if !m.cmds[m.command] { // invalid command
		fmt.Printf("> Invalid command recieved: %s\n", m.command)
		return nil
	}
	m.argument = line[strings.IndexAny(line, ":")+1:]
	return &m
}

func NewOutgoingMessage(prefix, command, argument string) []byte {
	var outgoing string
	if argument != "" {
		argument = " :" + argument
	}
	if len(prefix) == 0 {
		outgoing = fmt.Sprintf("%s%s\r\n", command, argument)
	} else {
		outgoing = fmt.Sprintf(":%s %s%s\r\n", prefix, command, argument)
	}
	return []byte(outgoing)
}

func (m *IRCMessage) Arg() string {
	return m.argument
}

func (m *IRCMessage) Cmd() string {
return m.command
}

func (m *IRCMessage) Prefix() string {
	return m.prefix
}

func (m *IRCMessage) String() string {
	return string(NewOutgoingMessage(m.prefix, m.command, m.argument))
}

// map (as a set) of possible commands
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