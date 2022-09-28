package main

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
	Discord discordConfig `json:"discord"`
}

type discordConfig struct {
	Token             string `json:"token"`
	AttendanceChannel string `json:"attendance_channel"`
}

var conf config

func init() {
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(bytes, &conf)
}
