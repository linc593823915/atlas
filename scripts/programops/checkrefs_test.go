package main

import "testing"

func TestExtractLocalRefs(t *testing.T) {
	content := `[good](../docs/a.md)
[external](https://example.com)
[anchor](#section)
[json](program-state.json)`

	refs := extractLocalRefs(content)
	if len(refs) != 2 {
		t.Fatalf("len(refs) = %d, want 2", len(refs))
	}
	if refs[0] != "../docs/a.md" || refs[1] != "program-state.json" {
		t.Fatalf("unexpected refs: %#v", refs)
	}
}
