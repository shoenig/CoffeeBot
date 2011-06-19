package irc

import "testing"

type NewOutgoingMessageTest struct {
	in_prefix   string
	in_command  string
	in_argument string
	out         []byte
}

var NewOutgoingMessageTests = []NewOutgoingMessageTest{
	NewOutgoingMessageTest{"", "PONG", "wolfe.irc.net", []byte("PONG :wolfe.irc.net\r\n")},
	NewOutgoingMessageTest{"", "JOIN #botwar", "", []byte("JOIN #botwar\r\n")},
	NewOutgoingMessageTest{"Alpha", "PRIVMSG #botwar", "hey there",
		[]byte(":Alpha PRIVMSG #botwar :hey there\r\n")},
}

func TestNewOutgoingMessage(t *testing.T) {
	for _, nomt := range NewOutgoingMessageTests {
		result := NewOutgoingMessage(nomt.in_prefix, nomt.in_command, nomt.in_argument)
		if string(result) != string(nomt.out) {
			t.Errorf("NOMT failed, exp: %q, got: %q", string(nomt.out), string(result))
		}
	}
}
