package main

import (
	"fmt"

	"github.com/ipoluianov/xchg-network/cmd"
)

func main() {
	cmd.CmdCreateKey("123")
	fmt.Println("started")
}
