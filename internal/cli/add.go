package cli

import "fmt"

func runAdd(args []string) error {
	fmt.Println("add called with: ", args)
	return nil
}
