package main

import "fmt"
import "irc"

func main() {
    fmt.Printf("Hello World\n")
    ircc := irc.NewIRCClient(8000, "chat.freenode.net", "coffeebot", "cbot", "ident", "Coffee Bot", "Seth Hoenig", "#help")
    fmt.Printf("%v\n", ircc)
    ircc.PokeInternet()
}

