package main

import (
	"fmt"
	"os"
)

func main() {
	filename := ParseArgs()
}

func ParseArgs() string {
	args := os.Args
	argc := len(args)

	cmdName := args[0]

	if argc < 2 {
		fmt.Printf("usage: %s Prog.asm\n", cmdName)
	}

	filename := args[1]

	return filename
}
