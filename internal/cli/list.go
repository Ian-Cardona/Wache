package cli

import "fmt"

func runList(args []string) error {
	fmt.Println("list called with: ", args)
	return nil
}
