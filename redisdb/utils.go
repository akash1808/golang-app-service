
package redisdb

import (
	"github.com/BurntSushi/toml"
	"fmt"
)

const configFilePath = "/home/akash/config.toml"

func (rDB *RedisDB) initConfig() {

	var conf RedisDB
	if _, err := toml.DecodeFile(configFilePath, &conf); err != nil {
		fmt.Println("Error in redisDB: Init: toml.DecodeFile: ", err)
	}
	fmt.Println("Successfully fetched redisDB config:", conf.Config)
	rDB.Config = conf.Config
}

