/* config.go*/
/*
modification history
--------------------
2017/07/18, by linxiongmin, create
*/
/*
DESCRIPTION
*/

package mini_spider_config

import (
	"fmt"
)

import (
	"code.google.com/p/gcfg"
)

type MiniSpiderConf struct {
	Basic  BasicConfig
}

// load config and check for validation
func LoadConfig(confPath string) (MiniSpiderConf, error) {
	conf, err := loadConfig(confPath)
	if err != nil {
		return conf, fmt.Errorf("configLoad(): %s\n", err.Error())
	}

	// check for validation
	if err = conf.Check(); err != nil {
		return conf, fmt.Errorf("conf.check(): %s\n", err.Error())
	}

	return conf, nil
}

func loadConfig(confPath string) (MiniSpiderConf, error) {
	var conf MiniSpiderConf
	var err error

	//load conf file to conf struct
	err = gcfg.ReadFileInto(&conf, confPath)
	if err != nil {
		return conf, fmt.Errorf("gcfg.ReadFileInto(): %s\n", err.Error())
	}

	return conf, nil
}

// check config validation
func (conf *MiniSpiderConf) Check() error {
	var err error

    // check conf data
	if err = conf.Basic.Check(); err != nil {
		return fmt.Errorf("Basic.Check(): %s", err.Error())
	}

	return nil
}
