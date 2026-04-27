package cli

import "testing"

func TestRunAdd(t *testing.T) {
	args := []string{"foo", "bar"}
	got := runAdd(args)

	if got != nil {
		t.Errorf("want nil got %v", got)
	}
}
