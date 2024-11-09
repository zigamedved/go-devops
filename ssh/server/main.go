package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zigamedved/go-devops/ssh"
)

func main() {

	var err error

	authorizedKeysBytes, err := os.ReadFile("../keys/mykey.pub")
	if err != nil {
		log.Fatalf("Failed to load authorized public keys, err: %v", err)
	}
	privateKey, err := os.ReadFile("../keys/server.pem")
	if err != nil {
		log.Fatalf("Failed to load server.pem, err: %v", err)
	}

	if err = ssh.StartServer(privateKey, authorizedKeysBytes); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

}
