package cli

import "testing"

func TestRunAdd(t *testing.T) {
	t.Run("valid url", func(t *testing.T) {
		err := RunAdd([]string{"https://example.com"})
		if err != nil {
			t.Errorf("want nil, got %v", err)
		}
	})

	t.Run("missing url", func(t *testing.T) {
		err := RunAdd([]string{})
		if err == nil {
			t.Error("want error, got nil")
		}
	})

	t.Run("too many positionals", func(t *testing.T) {
		err := RunAdd([]string{"https://a.com", "https://b.com"})
		if err == nil {
			t.Error("want error, got nil")
		}
	})

	t.Run("with extractor flag", func(t *testing.T) {
		err := RunAdd([]string{"--extractor", "greenhouse", "https://example.com"})
		if err != nil {
			t.Errorf("want nil, got %v", err)
		}
	})
}
