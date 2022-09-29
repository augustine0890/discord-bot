package config

import (
	"encoding/json"
	"io/ioutil"
)

var (
	Prefix            string
	Token             string
	AttendanceChannel string
	Database          string
	conf              *Config
)

type Config struct {
	Prefix            string `json:"prefix"`
	Token             string `json:"token"`
	AttendanceChannel string `json:"attendance_channel"`
	Database          string `json:"database"`
}

func ReadConfig() (*Config, error) {
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(bytes, &conf)

	// Prefix = conf.Prefix
	// Token = conf.Token
	// AttendanceChannel = conf.AttendanceChannel
	Database = conf.Database

	return conf, nil
}
