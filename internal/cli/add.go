package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
)

func runAdd(args []string) error {
	fs := flag.NewFlagSet("add", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	extractor := fs.String("extractor", "", "extractor name (default: auto-detect)")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("parsing flags: %w", err)
	}

	positionals := fs.Args()
	if len(positionals) != 1 {
		return errors.New("usage: wache add [--extractor NAME] <url>")
	}

	url := positionals[0]
	fmt.Printf("add: url=%s extractor=%s\n", url, *extractor)
	return nil
}
