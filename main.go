package main

import (
	"backend-mono/cmd/config"
	"backend-mono/cmd/database/mysql"
	"backend-mono/cmd/handler"
	"backend-mono/cmd/svc"
	config2 "backend-mono/core/config"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	configType          = flag.String("config-type", "file", "Config type: file or remote")
	configFilePath      = flag.String("config-file-path", "configs", "Config file path: path to config dir")
	configRemoteAddress = flag.String("config-remote-address", "127.0.0.1:8500", "Address of remote config: ip:port")
	configRemoteKeys    = flag.String("config-remote-keys", "", "Keys on remote config: separate by ,")
)

func initConfig() {
	configSource := &config2.Viper{
		ConfigType:    *configType,
		FilePath:      *configFilePath,
		RemoteAddress: *configRemoteAddress,
		RemoteKeys:    *configRemoteKeys,
	}
	err := configSource.InitConfig()
	if err != nil {
		panic(err)
	}
}

func getBootstrapConfig() *config.Config {
	bc := &config.Config{}
	if err := viper.UnmarshalKey("http", &bc.Http); err != nil {
		panic(err)
	}
	if err := viper.UnmarshalKey("database", &bc.Database); err != nil {
		panic(err)
	}
	if err := viper.UnmarshalKey("redis", &bc.Redis); err != nil {
		panic(err)
	}
	return bc
}

func main() {
	flag.Parse()
	initConfig()
	c := getBootstrapConfig()

	userDb, err := mysql.NewUserDB()
	if err != nil {
		panic(err)
	}

	serverCtx := svc.NewServiceContext(*c, userDb)
	router := gin.Default()

	handler.RegisterHandlers(router, serverCtx)

	fmt.Printf("Starting server at %s:%s...\n", c.Http.Path, c.Http.Port)
	err = router.Run(fmt.Sprintf("%s:%s", c.Http.Path, c.Http.Port))
	if err != nil {
		panic(err)
	}
}
