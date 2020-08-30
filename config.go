package main

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/mitchellh/go-homedir"
	"log"
)

var k = koanf.New(".")

func loadConfig() bool {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	if err := k.Load(file.Provider(home + "/.ehmgr.json"), json.Parser()); err != nil {
		log.Fatalf("Failed to load config: %s", err)
		return false
	}

	ok, errs := checkKeySet([]string{
		"endpoint",
		"key",
		"ip",
		"domain",
		"package",
	})

	if !ok {
		log.Fatalf("Invalid configuration, keys not found: %v", errs)
	}

	return true
}

func checkKeySet(keys []string) (bool, []string) {
	var errs []string
	for _, key := range keys {
		if !k.Exists(key) {
			errs = append(errs, key)
		}
	}
	return len(errs) == 0,  errs
}