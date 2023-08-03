package main

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/MixinNetwork/tip/messenger"
	"github.com/MixinNetwork/trusted-group/mtg"
	"github.com/pelletier/go-toml"
)

type Configuration struct {
	MTG       *mtg.Configuration            `toml:"mtg"`
	Messenger *messenger.MixinConfiguration `toml:"messenger"`
}

func ReadConfiguration(path string) (*Configuration, error) {
	if strings.HasPrefix(path, "~/") {
		usr, _ := user.Current()
		path = filepath.Join(usr.HomeDir, (path)[2:])
	}
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var conf Configuration
	err = toml.Unmarshal(f, &conf)
	return &conf, err
}
