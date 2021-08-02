package boot

import (
	"github.com/pkg/errors"
	"hello/model"
	"log"

	"github.com/BurntSushi/toml"
)

func loadConfig(path string, recv interface{}) error {
	_, err := toml.DecodeFile(path, recv)
	if nil != err {
		return err
	}

	log.Printf("load configure file got: %+v", err)
	return nil
}

func bootConfig(configPath string) error {
	if err := loadConfig(configPath, model.GetConfig()); nil != err {
		return errors.Wrap(err, "loadConfig")
	}

	return nil
}
