package conf

import (
    "github.com/zpatrick/go-config"
    "log"
)

const (
    CONFIG_FILE = "../config.ini"
)

func LoadConfig(file string) (*config.Config, error) {
    iniFile := config.NewINIFile(file)
    conf := config.NewConfig([]config.Provider{iniFile})
    if err := conf.Load(); err != nil {
        log.Fatal(err)
        return nil, err;
    }
    return conf, nil;
}
