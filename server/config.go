package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("yaml")
}

type config struct {
	MySQL MySQLInfo
}

type MySQLInfo struct {
	Host     string
	Port     int
	DB       string
	Username string
	Password string
}

func (m MySQLInfo) DataSource() string {
	opts := url.Values{}
	opts.Set("parseTime", "true")
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DB,
		opts.Encode(),
	)
}

func loadFile(path string) (config, error) {
	var cfg config
	file, err := os.Open(path)
	if err != nil {
		return cfg, err
	}
	defer file.Close()
	if err := viper.MergeConfig(file); err != nil {
		return cfg, err
	}
	err = viper.Unmarshal(&cfg)
	return cfg, err
}
