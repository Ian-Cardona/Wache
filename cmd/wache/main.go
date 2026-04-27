package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type CommandFunc func(args []string) error

func runAdd(args []string) error {
	for _, v := range args {
		fmt.Printf("Running %s...", v)
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		prog := filepath.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "%s: usage: %s <command> [args...]\n", prog, prog)
	}

	commands := map[string]func([]string) error{
		"add": runAdd,
	}

	handler, ok := commands[os.Args[1]]
	if !ok {
		os.Exit(1)
	}

	err := handler(os.Args[2:])
	if err != nil {
		fmt.Println("Exiting...")
	}
}
