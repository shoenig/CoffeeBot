package bot

import "fmt"
import "testing"
import "os"

import "irc"

type NewConfigReaderTest struct {
	port     uint16
	host     string
	nick     string
	ident    string
	realname string
	owner    string
	channel  string
	password string
	out      *irc.IRCConfig
}

var NewConfigReaderTests = []NewConfigReaderTest{
	NewConfigReaderTest{6667, "wolfe.irc.net", "mynick", "myident", "myrealname", "myowner", "mychannel", "mypass",
		&irc.IRCConfig{6667, "wolfe.irc.net", "mynick", "myident", "myrealname", "myowner", "mychannel", "mypass"}},
	NewConfigReaderTest{6697, "wolfe.irc.net", "mynick", "myident", "myrealname", "myowner", "mychannel", "mypass",
		&irc.IRCConfig{6697, "wolfe.irc.net", "mynick", "myident", "myrealname", "myowner", "mychannel", "mypass"}},
	NewConfigReaderTest{6697, "chat.irc.net", "mynick", "myident", "myrealname", "myowner", "mychannel", "mypass",
		&irc.IRCConfig{6697, "chat.irc.net", "mynick", "myident", "myrealname", "myowner", "mychannel", "mypass"}},
	NewConfigReaderTest{6697, "chat.irc.net", "mynick", "myident", "myrealname", "myowner", "mychannel", "",
		&irc.IRCConfig{6697, "chat.irc.net", "mynick", "myident", "myrealname", "myowner", "mychannel", ""}},
	NewConfigReaderTest{6697, "chat.irc.net", "mynick", "", "myrealname", "myowner", "mychannel", "",
		&irc.IRCConfig{6697, "chat.irc.net", "mynick", "", "myrealname", "myowner", "mychannel", ""}},
}

func TestConfigReader(t *testing.T) {
	for _, crt := range NewConfigReaderTests {

		tmpFile, terr := os.Create("/tmp/testconfig.dat")
		if terr != nil {
			panic("Error")
		}

		dictStr := fmt.Sprintf("{ \"Port\":\"%d\", \"Host\":\"%s\", \"Nick\":\"%s\", \"Ident\":\"%s\", \"Realname\":\"%s\", \"Owner\":\"%s\", \"Channel\":\"%s\", \"Password\":\"%s\"}",
			crt.port, crt.host, crt.nick, crt.ident, crt.realname, crt.owner, crt.channel, crt.password)

		_, werr := tmpFile.WriteString(dictStr)

		if werr != nil {
			panic("Werror")
		}

		result := ReadJSONConfig("/tmp/testconfig.dat")
		if result.Port != crt.out.Port {
			t.Errorf("Invalid Port, expected: %v, got: %v", crt.out.Port, result.Port)
		}
		if result.Host != crt.out.Host {
			t.Errorf("Invalid Host, expected: %v, got: %v", crt.out.Host, result.Host)
		}
		if result.Nick != crt.out.Nick {
			t.Errorf("Invalid Nick, expected: %v, got: %v", crt.out.Nick, result.Nick)
		}
		if result.Ident != crt.out.Ident {
			t.Errorf("Invalid Ident, expected: %v, got: %v", crt.out.Ident, result.Ident)
		}
		if result.Realname != crt.out.Realname {
			t.Errorf("Invalid Realname, expected: %v, got: %v", crt.out.Realname, result.Realname)
		}
		if result.Owner != crt.out.Owner {
			t.Errorf("Invalid Owner, expected: %v, got: %v", crt.out.Owner, result.Owner)
		}
		if result.Channel != crt.out.Channel {
			t.Errorf("Invalid Channel, expected: %v, got: %v", crt.out.Channel, result.Channel)
		}
		if result.Password != crt.out.Password {
			t.Errorf("Invalid Password, expected: %v, got: %v", crt.out.Password, result.Password)
		}
	}
}
