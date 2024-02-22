package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Printf("hello world\nArguments: %v\n", args[1:])
}

// run with: go run main.go arg1 arg2 ...
// echo $? = shows exit code of last executed program
