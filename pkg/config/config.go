package config

import (
	"encoding/json"
	"os"

	"github.com/urfave/cli/v2"
)

type (
	// Config is a complete set of app configuration
	Config struct {
		Token       string
		Slug        string
		Changelog   string
		Ghr         bool
		Noci        bool
		Dry         bool
		Vf          bool
		Update      string
		GheHost     string
		Prerelease  bool
		TravisCom   bool
		BetaRelease BetaRelease
		Match       string
	}

	BetaRelease struct {
		MaintainedVersion string `json:"maintainedVersion"`
	}
)

// NewConfig returns a new Config instance
func NewConfig(c *cli.Context) *Config {
	conf := &Config{
		Token:      c.String("token"),
		Slug:       c.String("slug"),
		Changelog:  c.String("changelog"),
		Ghr:        c.Bool("ghr"),
		Noci:       c.Bool("noci"),
		Dry:        c.Bool("dry"),
		Vf:         c.Bool("vf"),
		Update:     c.String("update"),
		GheHost:    c.String("ghe-host"),
		Prerelease: c.Bool("prerelease"),
		TravisCom:  c.Bool("travis-com"),
		Match:      c.String("match"),
	}

	f, err := os.OpenFile(".semrelrc", os.O_RDONLY, 0)
	if err != nil {
		return conf
	}
	defer f.Close()

	src := &BetaRelease{}
	json.NewDecoder(f).Decode(src)
	conf.BetaRelease = *src

	return conf
}
