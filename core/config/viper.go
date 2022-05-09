package config

import (
	"backend-mono/utils/constant"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"strings"
)

type Viper struct {
	ConfigType    string
	FilePath      string
	RemoteAddress string
	RemoteKeys    string
}

func (c *Viper) InitConfig() error {
	if c.ConfigType == constant.ConfigTypeFile {
		return c.LoadConfigFromFile()
	} else {
		return c.LoadConfigFromConsul()
	}
}

func (c *Viper) LoadConfigFromFile() error {
	viper.AddConfigPath(c.FilePath)
	viper.SetEnvPrefix("app")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Errorw("Failed while read config from file", "file-path", c.FilePath, "error", err.Error())
		return err
	}
	log.Infof("Config loaded from file")
	return nil
}

func (c *Viper) LoadConfigFromConsul() error {
	viper.SetEnvPrefix("app")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()
	configRemoteKeys := strings.Split(c.RemoteKeys, ",")
	if len(configRemoteKeys) > 0 {
		for _, remoteKey := range configRemoteKeys {
			err := viper.AddRemoteProvider("consul", c.RemoteAddress, remoteKey)
			if err != nil {
				log.Errorw("Failed while add remote key", "remote-key", remoteKey, "error", err.Error())
				continue
			}
			log.Infof("Add config remote key: %s", remoteKey)
		}
	}
	viper.SetConfigType("yaml")
	err := viper.ReadRemoteConfig()
	if err != nil {
		log.Errorw("Failed while read remote config", "error", err.Error())
		return err
	}
	log.Infof("Config loaded from remote")
	return nil
}
