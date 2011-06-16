package main

import "fmt"
import "irc"

func main() {
    fmt.Printf("Hello World\n")
    ircc := irc.NewIRCClient(8093, "coffeebot", "cbot", "ident", "Coffee Bot", "Seth Hoenig", "irc.freenode.net")
    fmt.Printf("%v\n", ircc)
    ircc.PokeInternet()
}

