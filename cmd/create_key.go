package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ipoluianov/xchg/xchg"
)

func CurrentExePath() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

func CmdCreateKey(password string) {
	encryptedPrivateKeyBase64, publicKeyBase64, err := xchg.NetworkContainerCreateKey(password)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	err = ioutil.WriteFile(CurrentExePath()+"/xchgr_private_key.base64", []byte(encryptedPrivateKeyBase64), 0666)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	err = ioutil.WriteFile(CurrentExePath()+"/xchgr_public_key.base64", []byte(publicKeyBase64), 0666)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
}
