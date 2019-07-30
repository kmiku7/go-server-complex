package config

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tomlRaw := `[core]
	listen_address = ":8066"

[log]
	log_dir = "../log"
`

	tempFile, err := ioutil.TempFile("", "tmp-config")
	if err != nil {
		t.Fatalf("create temporary config file failed, err: %v.", err)
	}

	writeLen, err := tempFile.Write([]byte(tomlRaw))
	if err != nil {
		t.Fatalf("write config file content failed, err: %v.", err)
	}
	if writeLen != len(tomlRaw) {
		t.Fatalf("write config file content incomplete, expect: %v, actual: %v",
			len(tomlRaw), writeLen)
	}

	globalConfig := GetGlobalConfig()
	if globalConfig != nil {
		t.Errorf("global config instance should be nil now.")
	}
	err = InitGlobalConfig(tempFile.Name())
	if err != nil {
		t.Fatalf("initialize global config failed, err: %v.", err)
	}
	globalConfig = GetGlobalConfig()
	if globalConfig == nil {
		t.Fatalf("global config instance should not be nil now.")
	}

	buffer := new(bytes.Buffer)
	encoder := toml.NewEncoder(buffer)
	encoder.Indent = "\t"
	err = encoder.Encode(globalConfig)
	if err != nil {
		t.Fatalf("serialize global config instance failed, err: %v.", err)
	}
	if buffer.String() != tomlRaw {
		t.Errorf("Config content not match, expect: %v, actual: %v.",
			tomlRaw, buffer.String())
	}
}
