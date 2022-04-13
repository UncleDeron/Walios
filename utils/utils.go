package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ServerConfig struct {
	Host             string
	Port             int
	MaxContentLength uint32
	Redis            RedisConfig
}

type RedisConfig struct {
	Host        string
	Port        int
	Password    string
	DB          int
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

func LoadConfig(configFilePath string, config *ServerConfig) error {
	// Load configuration from file
	if confFileExists, error := PathExists(configFilePath); confFileExists != true {
		fmt.Println("Config File ", configFilePath, " is not exist!!")
		return error
	}
	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}

	//将json数据解析到struct中
	err = json.Unmarshal(data, config)
	if err != nil {
		panic(err)
	}
	return nil
}

//PathExists 判断一个文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
