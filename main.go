package main

import (
	"fmt"

	"github.com/ipoluianov/xchg-network/cmd"
)

func main() {
	cmd.CmdCreateKey("12345")
	cmd.CmdCreateNetworkContainer("12345")
	fmt.Println("started")
}
