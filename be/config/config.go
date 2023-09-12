package config

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/MixinNetwork/trusted-group/mtg"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/mvg-fi/common/messenger"
	"github.com/mvg-fi/common/web"
	"github.com/pelletier/go-toml"
)

type Proxy struct {
	ProxyPIN        string `toml:"proxy-pin"`
	ProxyUserSecret string `toml:"proxy-user-secret"`
}

type Configuration struct {
	MTG       *mtg.Configuration            `toml:"mtg"`
	Messenger *messenger.MixinConfiguration `toml:"messenger"`
	API       *web.Configuration            `toml:"api"`
	ProxyRoot *mixin.Keystore               `toml:"proxyroot"`
	Proxy     *Proxy                        `toml:"proxy"`
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
