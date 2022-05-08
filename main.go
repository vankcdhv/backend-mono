package main

import (
	"backend-mono/database/mysql"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"

	"backend-mono/cmd/config"
	"backend-mono/cmd/handler"
	"backend-mono/cmd/svc"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/file-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	userDb, err := mysql.NewUserDB()
	if err != nil {
		panic(err)
	}

	serverCtx := svc.NewServiceContext(c, userDb)
	router := gin.Default()

	handler.RegisterHandlers(router, serverCtx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	err = router.Run(fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		panic(err)
	}
}
