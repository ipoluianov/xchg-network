package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"golang.org/x/term"

	"github.com/ipoluianov/xchg/xchg"
)

func CurrentExePath() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

func EnterPassword() (string, error) {
	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	password := string(bytePassword)
	return strings.TrimSpace(password), nil
}

// Generate RSA key pair.
// AES-GSM encrypted private key (PKCS8) is stored in file xchgr_private_key.base64 in base64 format
// RSA public key is stored in xchgr_public_key.base64 in base64 format
func CmdCreateKey(password string) {
	encryptedPrivateKeyBase64, publicKeyBase64, err := xchg.NetworkContainerCreateKey(password)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	err = os.MkdirAll(CurrentExePath()+"/xchg-network-result", 0777)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	err = ioutil.WriteFile(CurrentExePath()+"/xchg-network-result/xchgr_private_key.base64", []byte(encryptedPrivateKeyBase64), 0666)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	err = ioutil.WriteFile(CurrentExePath()+"/xchg-network-result/xchgr_public_key.base64", []byte(publicKeyBase64), 0666)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("SUCCESS")
}

// Make network container.
// Result: zip-file in base64 format. It is stored in file network.zip.base64
func CmdCreateNetworkContainer(network *xchg.Network, password string) {

	zipFile, err := xchg.NetworkContainerMake(network, xchg.NetworkContainerEncryptedPrivateKey, password)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	err = os.MkdirAll(CurrentExePath()+"/xchg-network-result", 0777)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	err = ioutil.WriteFile(CurrentExePath()+"/xchg-network-result/network.zip", zipFile, 0666)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	zipFileBase64 := base64.StdEncoding.EncodeToString(zipFile)

	err = ioutil.WriteFile(CurrentExePath()+"/xchg-network-result/network.zip.base64", []byte(zipFileBase64), 0666)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println("SUCCESS")
}

func makeNetwork() *xchg.Network {
	network := xchg.NewNetwork()
	network.Name = "MainNet-23-09-18-001"
	network.Timestamp = time.Now().Unix()

	network.InitialPoints = make([]string, 0)
	network.InitialPoints = append(network.InitialPoints, "http://xchgx.net/network.zip.base64")
	network.InitialPoints = append(network.InitialPoints, "http://xchg.gazer.cloud/network.zip.base64")

	s1 := "x03.gazer.cloud:8084"
	s2 := "x04.gazer.cloud:8084"

	for r := 0; r < 16; r++ {
		rangePrefix := fmt.Sprintf("%X", r)
		network.AddHostToRange(rangePrefix, s1)
		network.AddHostToRange(rangePrefix, s2)
	}

	return network
}

func main() {
	fmt.Println("xchg-network")

	for {
		fmt.Print("cmd>")
		var cmd string
		fmt.Scanln(&cmd)
		if cmd == "quit" || cmd == "exit" {
			break
		}

		if cmd == "generate" {
			password, _ := EnterPassword()
			CmdCreateKey(password)
		}

		if cmd == "network" {
			password, _ := EnterPassword()
			CmdCreateNetworkContainer(makeNetwork(), password)
		}
	}
}
