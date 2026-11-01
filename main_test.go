package main

import (
	"strings"
	"testing"
)

func TestBackgroundUpdateIsRejected(t *testing.T) {
	t.Setenv("GHX_CONFIG", t.TempDir()+"/config.yaml")
	err := run([]string{backgroundUpdateFlag})
	if err == nil || !strings.Contains(err.Error(), "background updates are disabled") {
		t.Fatalf("background invocation must be rejected: %v", err)
	}
}
