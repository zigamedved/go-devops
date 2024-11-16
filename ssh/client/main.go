package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {

	var err error

	privateKey, err := os.ReadFile("../keys/mykey.pem")
	if err != nil {
		log.Fatalf("Failed to load mykey.pem, err: %v", err)
	}
	publicKey, err := os.ReadFile("../keys/server.pub")
	if err != nil {
		log.Fatalf("Failed to load server.pub, err: %v", err)
	}

	privateKeyParsed, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		log.Fatalf("Failed to parse private key, err: %v", err)
	}
	publicKeyParsed, _, _, _, err := ssh.ParseAuthorizedKey(publicKey)
	if err != nil {
		log.Fatalf("Failed to parse public key, err: %v", err)
	}

	config := &ssh.ClientConfig{
		User: "username",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(privateKeyParsed),
		},
		HostKeyCallback: ssh.FixedHostKey(publicKeyParsed),
	}
	client, err := ssh.Dial("tcp", "localhost:2022", config)
	if err != nil {
		log.Fatalf("Failed to dial, err: %v", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create new session, err: %v", err)
	}
	defer session.Close()

	output, err := session.Output("whoami") // uses exec request
	if err != nil {
		log.Fatalf("Failed to execute `whoami`, err: %v", err)
	}

	fmt.Printf("Output is %s\n", output)
}
