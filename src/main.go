package main

import "fmt"
import "io/ioutil"
import "strconv"
import "strings"

import "bot"
import "irc"

// FreeNode ports: 6665, 6666, 6667, 8000, 8001, 8002
// FreeNode SSL ports: 6697  7000 7070  

func main() {
    fName := "config"
    fmt.Printf("CoffeeBot v0.0 2011\n")
    ircbot := bot.NewBot(ReadConfig(fName))
    ircbot.SayHi()
    Scratch()
}

func Scratch() {
    port := uint16(8000)
    host := "wolfe.freenode.net"
    nick := "coffeebot"
    name := "coffeebot"
    ident := "coffeebot"
    realname := "Coffee Bot"
    owner := "Seth Hoenig"
    ircc := irc.NewIRCClient(port, host, nick, name, ident, realname, owner, "#test")
    fmt.Printf("%v\n", ircc)
    ircc.PokeInternet()
    ircc.MainLoop()
}

//TODO: make this less awful 
func ReadConfig(fName string) (uint16, string, string, string, string, string, string, string) {
    content, err := ioutil.ReadFile(fName)
    if err != nil { panic("Error Reading File") }
    values := strings.Fields(string(content))
    port := uint16(0)
    host := ""
    nick := ""
    name := ""
    ident := ""
    realname := ""
    owner := ""
    channel := ""
    for i:=0; i<len(values); {
        switch(values[i]) {
            case "PORT":
                temp, _ := strconv.Atoi(values[i+1])
                port = uint16(temp)
            case "HOST": host = values[i+1]
            case "NICK": nick = values[i+1]
            case "NAME": name = values[i+1]
            case "IDENT": ident = values[i+1]
            case "REALNAME": realname = values[i+1]
            case "OWNER": owner = values[i+1]
            case "CHANNEL": channel = values[i+1]
            default:
                fmt.Printf("option: %s\n", values[i])
                panic("Invalid Option")
        }
        i+=2
    }
    return port, host, nick, name, ident, realname, owner, channel
}
