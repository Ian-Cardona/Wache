package cli

import "testing"

func TestRunList(t *testing.T) {
	t.Run("default format", func(t *testing.T) {
		if err := RunList([]string{}); err != nil {
			t.Errorf("want nil, got %v", err)
		}
	})

	t.Run("explicit table format", func(t *testing.T) {
		if err := RunList([]string{"--format", "table"}); err != nil {
			t.Errorf("want nil, got %v", err)
		}
	})

	t.Run("json format", func(t *testing.T) {
		if err := RunList([]string{"--format", "json"}); err != nil {
			t.Errorf("want nil, got %v", err)
		}
	})

	t.Run("invalid format", func(t *testing.T) {
		if err := RunList([]string{"--format", "xml"}); err == nil {
			t.Error("want error, got nil")
		}
	})

	t.Run("rejects positional args", func(t *testing.T) {
		if err := RunList([]string{"foo"}); err == nil {
			t.Error("want error, got nil")
		}
	})
}
