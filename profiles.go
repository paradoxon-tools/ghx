package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/paradoxon-tools/ghx/internal/config"
	"github.com/paradoxon-tools/ghx/internal/proxy"
)

type profileInfo struct {
	ID   string `json:"id"`
	User string `json:"user"`
	Host string `json:"host"`
}

func configuredProfiles(cfg config.Config) []profileInfo {
	profiles := make([]profileInfo, 0, len(cfg.Profiles))
	for id, profile := range cfg.Profiles {
		profiles = append(profiles, profileInfo{ID: id, User: profile.User, Host: profile.HostOrDefault()})
	}
	sort.Slice(profiles, func(i, j int) bool { return profiles[i].ID < profiles[j].ID })
	return profiles
}

// Integration commands never resolve directories or apply repository Git identity.
func runProfiles(cfg config.Config, args []string) error {
	if len(args) == 2 && args[0] == "list" && args[1] == "--json" {
		return json.NewEncoder(os.Stdout).Encode(struct {
			Version  int           `json:"version"`
			Profiles []profileInfo `json:"profiles"`
		}{1, configuredProfiles(cfg)})
	}
	if len(args) == 2 && args[0] == "token" {
		profile, ok := cfg.Profiles[args[1]]
		if !ok {
			return fmt.Errorf("unknown profile %q", args[1])
		}
		ghPath, err := locateGh(cfg, newUpdater(cfg, config.StateDir()))
		if err != nil {
			return err
		}
		token, err := proxy.FetchToken(ghPath, profile)
		if err != nil {
			return err
		}
		fmt.Println(token)
		return nil
	}
	return fmt.Errorf("usage: ghx ghx profiles list --json | ghx ghx profiles token <id>")
}
