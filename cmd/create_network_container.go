package cmd

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/ipoluianov/xchg/xchg"
)

func CmdCreateNetworkContainer(password string) {
	network := xchg.NewNetworkDefault()

	zipFile, err := xchg.NetworkContainerMake(network, xchg.NetworkContainerEncryptedPrivateKey, password)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	err = ioutil.WriteFile(CurrentExePath()+"/network.zip", zipFile, 0666)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	zipFileBase64 := base64.StdEncoding.EncodeToString(zipFile)

	err = ioutil.WriteFile(CurrentExePath()+"/network.zip.base64", []byte(zipFileBase64), 0666)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
}
