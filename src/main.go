package main

import "fmt"
import "irc"

// FreeNode ports: 6665, 6666, 6667, 8000, 8001, 8002
// FreeNode SSL ports: 6697  7000 7070  

func main() {
    fmt.Printf("Hello World\n")
    ircc := irc.NewIRCClient(6666, "chat.freenode.net", "coffeebot", "cbot", "ident", "Coffee Bot", "Seth Hoenig", "#cbottestting")
    fmt.Printf("%v\n", ircc)
    ircc.PokeInternet()
}

