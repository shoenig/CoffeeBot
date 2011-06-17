package bot

import "fmt"
import "io/ioutil"
import "strings"

type Bot struct{
}

func NewBot(config string) *Bot {
    readConfig(config)
    var b Bot
    return &b
}

//TODO: make this less awful 
func readConfig(fName string) (string, string, string, string, string, string, string, string) {
    content, err := ioutil.ReadFile(fName)
    if err != nil { panic("Error Reading File") }
    values := strings.Fields(string(content))
    port := ""
    host := ""
    nick := ""
    name := ""
    ident := ""
    realname := ""
    owner := ""
    channel := ""
    for i:=0; i<len(values); {
        switch(values[i]) {
            case "PORT": port = values[i+1]
            case "HOST": host = values[i+1]
            case "NICK": nick = values[i+1]
            case "IDENT": ident = values[i+1]
            case "REALNAME": realname = values[i+1]
            case "OWNER": owner = values[i+1]
            case "CHANNEL": channel = values[i+1]
            default: panic("Invalid Option")
        }
        i+=2
    }
    return port, host, nick, name, ident, realname, owner, channel
}

func (b *Bot) SayHi() {
    fmt.Printf("Hi\n")
}


