package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.LookupIP("Rohan.local"))
}
