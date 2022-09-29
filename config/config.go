package config

import (
	"encoding/json"
	"io/ioutil"
)

var (
	Token             string
	AttendanceChannel string
	Database          string
	conf              *Config
)

type Config struct {
	Token             string `json:"token"`
	AttendanceChannel string `json:"attendance_channel"`
	Database          string `json:"database"`
}

func ReadConfig() error {
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(bytes, &conf)

	Token = conf.Token
	AttendanceChannel = conf.AttendanceChannel
	Database = conf.Database

	return nil
}
