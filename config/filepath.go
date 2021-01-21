package config

import (
	"flag"
	"github.com/twkun/util/process"
)

const (
	DefaultConfigFile = "config.toml"
)
func GetConfigFile() string {
	conf_file_path := flag.String("config", "./config.toml", "path to config file")
	flag.Parse()
	var ConfigFilePath = *conf_file_path
	var cf_file string
	if ConfigFilePath == "" {
		cf_file = process.GetRunPath() + "/" + DefaultConfigFile
	} else {
		cf_file = ConfigFilePath
	}
	return cf_file
}
