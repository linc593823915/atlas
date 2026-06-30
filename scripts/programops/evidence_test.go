package main

import "testing"

func TestCountSectionFieldMissing(t *testing.T) {
	body := `- current progress: done
- next action: review`

	if countSectionFieldMissing(body, "current progress") != 0 {
		t.Fatal("current progress should be present")
	}
	if countSectionFieldMissing(body, "missing") != 1 {
		t.Fatal("missing field should be counted")
	}
}
