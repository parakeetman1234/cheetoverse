package main

import (
	"fmt"
	"encoding/json"
	"strings"

	"gopkg.in/ini.v1"
)

type Config struct {
	DefaultAdmins []string
	DatabaseLocation string
	RollInterval int
	PollInterval int
	BatchSize int
	InstanceUrl string // credentials -> Instance
	Token string
}

func SectionMissingErr(section string) error {
	return fmt.Errorf("%s section missing")
}

func KeyMissingErr(key, section string) error {
	return fmt.Errorf("%s key in %s section missing")
}

// this function *painstakingly* loads every single element into a Config struct and returns the struct. LET ME KNOW IF THERES A BETTER WAY OF DOING THIS!!!!
func LoadConfig(ininame string) (Config, error) {
	var (
		cfg Config
		section *ini.Section
		key *ini.Key
		templist []string
		tempint int
	)
	rawini, err := ini.Load(ininame)
	if err != nil {
		return Config{}, err
	}

	// application section
	section = rawini.Section("application")
	if section == nil {
		return Config{}, SectionMissingErr("application")
	}
	key = section.Key("DefaultAdmins")
	if key == nil {
		return Config{}, KeyMissingErr("DefaultAdmins", "application") 
	}
	tempstr := key.String()
	tempstr = strings.ReplaceAll(tempstr, "'", "\"") // LMK if theres a better way of making sure a string is safe to pass into the unmarshaler.
	err = json.Unmarshal([]byte(tempstr), &templist)
	if err != nil {
		return Config{}, err
	}
	cfg.DefaultAdmins = templist
	key = section.Key("DatabaseLocation")
	if key == nil {
		return Config{}, KeyMissingErr("DatabaseLocation", "application")  
	}
	cfg.DatabaseLocation = key.String()

	// gacha section 
	section = rawini.Section("gacha")
	if section == nil { 
		return Config{}, SectionMissingErr("gacha")
	}
	key = section.Key("RollInterval")
	if key == nil {
		return Config{}, KeyMissingErr("RollInterval", "gacha") 
	}
	tempint, err = key.Int()
	if err != nil {
		return Config{}, err
	}
	cfg.RollInterval = tempint
	
	// notification section
	section = rawini.Section("notification")
	if section == nil {
		return Config{}, SectionMissingErr("section")
	}
	key = section.Key("PollInterval")
	if key == nil {
		return Config{}, KeyMissingErr("PollInterval", "notification")
	}
	tempint, err = key.Int()
	if err != nil {
		return Config{}, err
	}
	cfg.PollInterval = tempint
	key = section.Key("BatchSize")
	if key == nil {
		return Config{}, KeyMissingErr("PollInterval", "notification")
	}
	tempint, err = key.Int()
	if err != nil {
		return Config{}, err
	}
	cfg.BatchSize = tempint
	
	// credentials section
	section = rawini.Section("credentials")
	if section == nil {
		return Config{}, SectionMissingErr("credentials")
	}
	key = section.Key("Instance")
	if key == nil {
		return Config{}, KeyMissingErr("Instance", "credentials")
	}
	cfg.InstanceUrl = key.String()
	key = section.Key("Token")
	if key == nil {
		return Config{}, KeyMissingErr("Token", "credentials")
	}
	cfg.Token = key.String()

	return cfg, nil
}
