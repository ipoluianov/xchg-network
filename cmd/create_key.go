package cmd

import (
	"fmt"

	"github.com/ipoluianov/xchg/xchg"
)

func CmdCreateKey(password string) {
	encryptedPrivateKeyBase64, publicKeyBase64, err := xchg.NetworkContainerCreateKey()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
}
