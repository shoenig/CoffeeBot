package main

import "fmt"
import "io/ioutil"
import "os"
import "strconv"
import "strings"

import "bot"
import "irc"

// FreeNode ports: 6665, 6666, 6667, 8000, 8001, 8002
// FreeNode SSL ports: 6697  7000 7070  

//TODO: allow specify location of config
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

func ReadConfig(fName string) (uint16, string, string, string, string, string, string, string) {
    cfile, err := ioutil.ReadFile(fName)
    if err != nil {
        createDefaultConfig()
        fmt.Printf("Default config created in ./config Please update info\n")
        os.Exit(1)
    }
    port := uint16(0)
    host := ""
    nick := ""
    name := ""
    ident := ""
    realname := ""
    owner := ""
    channel := ""

    bylines := strings.Split(string(cfile), "\n", -1)
    for lnum, line := range bylines {
        line = strings.TrimSpace(line)
        if len(line) == 0 || line[0] == '#' { continue } // ignore comments
        lsplit := strings.Split(line, " ", -1)
        if len(lsplit) < 2 { panic(fmt.Sprintf("Invalid config file, see line %d\n", lnum)) }
        switch(lsplit[0]) {
            case "PORT":
                temp, err2 := strconv.Atoi(lsplit[1])
                if err2 != nil { panic(fmt.Sprintf("Invalid config (invalid port) line %d\n", lnum)) }
                port = uint16(temp)
            case "HOST": host = lsplit[1]
            case "NICK": nick = lsplit[1]
            case "NAME": name = lsplit[1]
            case "IDENT": ident = lsplit[1]
            case "REALNAME": realname = lsplit[1]
            case "OWNER": owner = lsplit[1]
            case "CHANNEL": channel = lsplit[1]
            default:
                panic(fmt.Sprintf("Invalid config option, line %d, %s\n", lnum, lsplit[0]))
        }
    }
    return port, host, nick, name, ident, realname, owner, channel
}

func createDefaultConfig() {
    oFile, err := os.Create("config")
    if err != nil { panic("Could not create config file") }
    defer oFile.Close()
    _, err = oFile.Write([]byte("PORT\nHOST\nNICK\nNAME\nIDENT\nREALNAME\nOWNER\nCHANNEL"))
    if err != nil { panic("Courld not write to config file") }
}
