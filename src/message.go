package irc

import "strings"

type IRCMessage struct {
	command string
	cmds    map[string]bool
}

// returns nil for empty or invalid messages
func NewMessage(line string) *IRCMessage {
	var m message
	m.cmds = makeCmdMap()
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return nil
	}
	if line[0] == ':' { // we have optional prefix
	} else {
		splitted := strings.Fields(line)
		m.command = splitted[0]
		if m.cmds[command] == nil { // invalid command
			fmt.Printf("> Invalid command recieved: %s\n", command)
			return nil
		}
	}
	return &m
}

func (m *IRCMessage) Command() bool {
	return m.command
}

func (m *IRCMessage) String() string {
}

// map (as a set) of possible commands
func makeCmdMap() map[string]bool {
	var m = map[string]bool{
		"ADMIN":    false,
		"AWAY":     false,
		"CONNECT":  false,
		"DIE":      false,
		"ERROR":    false,
		"INFO":     false,
		"INVITE":   false,
		"ISON":     false,
		"JOIN":     false,
		"KICK":     false,
		"KILL":     false,
		"LINKS":    false,
		"LIST":     false,
		"LUSERS":   false,
		"MODE":     false,
		"MOTD":     false,
		"NAMES":    false,
		"NICK":     false,
		"NOTICE":   false,
		"OPER":     false,
		"PART":     false,
		"PASS":     false,
		"PING":     false,
		"PONG":     false,
		"PRIVMSG":  false,
		"QUIT":     false,
		"REHASH":   false,
		"RESTART":  false,
		"SERVICE":  false,
		"SERVLIST": false,
		"SSERVER":  false,
		"SQUERY":   false,
		"SQUIT":    false,
		"STATS":    false,
		"SUMMON":   false,
		"TIME":     false,
		"TOPIC":    false,
		"TRACE":    false,
		"USER":     false,
		"USERHOST": false,
		"USERS":    false,
		"VERSION":  false,
		"WALLOPS":  false,
		"WHO":      false,
		"WHOIS":    false,
		"WHOWAS":   false,
	}
	return m
}
