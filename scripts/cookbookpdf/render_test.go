package main

import (
	"strings"
	"testing"
)

func TestTrimCommandOutput(t *testing.T) {
	input := strings.Repeat("x", commandOutputLimit+10)
	output := trimCommandOutput([]byte(input))
	if len(output) <= commandOutputLimit {
		t.Fatalf("trimmed output length = %d, want greater than limit with suffix", len(output))
	}
	if !strings.Contains(output, "output truncated") {
		t.Fatalf("trimmed output missing suffix: %q", output)
	}
}
