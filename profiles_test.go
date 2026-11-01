package main

import (
	"github.com/paradoxon-tools/ghx/internal/config"
	"testing"
)

func TestConfiguredProfiles(t *testing.T) {
	profiles := configuredProfiles(config.Config{Profiles: map[string]config.Profile{
		"work":     {User: "worker", Host: "github.example.com"},
		"personal": {User: "person"},
	}})
	if len(profiles) != 2 || profiles[0].ID != "personal" || profiles[0].Host != "github.com" || profiles[1].User != "worker" {
		t.Fatalf("unexpected profiles: %+v", profiles)
	}
	if err := runProfiles(config.Config{}, []string{"token", "missing"}); err == nil {
		t.Fatal("unknown profile must fail")
	}
}
