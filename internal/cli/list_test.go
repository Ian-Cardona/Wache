package cli

import "testing"

func TestRunList(t *testing.T) {
	args := []string{}
	got := runList(args)

	if got != nil {
		t.Errorf("want nil got %s", got)
	}
}
