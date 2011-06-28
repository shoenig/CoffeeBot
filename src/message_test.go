package irc

import "testing"

type NewOutgoingMessageTest struct {
	in_prefix   string
	in_command  string
	in_cmd_arg  string
	in_argument string
	out         []byte
}

var NewOutgoingMessageTests = []NewOutgoingMessageTest{
	NewOutgoingMessageTest{"", "PONG", "", "wolfe.irc.net", []byte("PONG :wolfe.irc.net\r\n")},
	NewOutgoingMessageTest{"", "JOIN", "#botwar", "", []byte("JOIN #botwar\r\n")},
	NewOutgoingMessageTest{"Alpha", "PRIVMSG", "#botwar", "hey there",
		[]byte(":Alpha PRIVMSG #botwar :hey there\r\n")},
}

func TestNewOutgoingMessage(t *testing.T) {
	for _, nomt := range NewOutgoingMessageTests {
		result := NewOutgoingMessage(nomt.in_prefix, nomt.in_command, nomt.in_cmd_arg, nomt.in_argument)
		if string(result) != string(nomt.out) {
			t.Errorf("NOMT failed, exp: %q, got: %q", string(nomt.out), string(result))
		}
	}
}

type NewIncomingMessageTest struct {
	in_line string
	out     *IRCMessage
}

var NewIncomingMessageTests = []NewIncomingMessageTest{
	NewIncomingMessageTest{"PING :wolfe.freenode.net",
		&IRCMessage{"", "PING", "wolfe.freenode.net", nil, nil}},
	NewIncomingMessageTest{":niekie!~niek@CAcert/Assurer/niekie JOIN :#botwar",
		&IRCMessage{"niekie!~niek@CAcert/Assurer/niekie", "JOIN", "#botwar", nil, nil}},
	NewIncomingMessageTest{":kaffee!kaffee@debiancenter/admin/kaffee QUIT :*.net *.split",
		&IRCMessage{"kaffee!kaffee@debiancenter/admin/kaffee", "QUIT", "*.net *.split", nil, nil}},
	NewIncomingMessageTest{":ChanServ!ChanServ@services. MODE #botwar +o kaeffchen",
		&IRCMessage{"ChanServ!ChanServ@services.", "MODE #botwar +o kaeffchen", "", nil, nil}},
}

func TestNewIncomingMessage(t *testing.T) {
	for _, nimt := range NewIncomingMessageTests {
		result := NewIncomingMessage(nimt.in_line)
		if !result.Eq(nimt.out) {
			t.Errorf("\nNIMT failed, exp: %q, got: %q", nimt.out, result)
			t.Errorf(".%s.%s.%s.\n", nimt.out.prefix, nimt.out.command, nimt.out.argument)
			t.Errorf(".%s.%s.%s.\n", result.prefix, result.command, result.argument)
		}
	}
}
