package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
)

func runList(args []string) error {
	fs := flag.NewFlagSet("list", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	format := fs.String("format", "table", "output format: table or json")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("parsing flags: %w", err)
	}

	if fs.NArg() > 0 {
		return errors.New("usage: wache list [--format table|json]")
	}

	switch *format {
	case "table", "json":
		// valid
	default:
		return fmt.Errorf("invalid format %q: must be 'table' or 'json'", *format)
	}

	fmt.Printf("list: format=%s\n", *format)
	return nil
}
