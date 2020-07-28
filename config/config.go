package config

import (
	"flag"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
)

var strpath string

func init() {
	flag.StringVar(&strpath, "conf", "./conf.toml", "set config file path")
}

//TomlConfig config
type TomlConfig struct {
	Title   string
	Owner   ownerInfo
	DB      database `toml:"database"`
	Servers map[string]server
	Clients clients
}

type ownerInfo struct {
	Name string
	Org  string `toml:"organization"`
	Bio  string
	DOB  time.Time
}

type database struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type server struct {
	IP string
	DC string
}

type clients struct {
	Data  [][]interface{}
	Hosts []string
}

var cfg *TomlConfig

//LoadConfig load config
func LoadConfig() (*TomlConfig, error) {
	filePath, err := filepath.Abs(strpath)
	if err != nil {
		return nil, err
	}
	config := new(TomlConfig)
	if _, err := toml.DecodeFile(filePath, config); err != nil {
		return nil, err
	}

	cfg = config
	return cfg, nil
}

//GetConfig get load config
func GetConfig() *TomlConfig {
	return cfg
}
