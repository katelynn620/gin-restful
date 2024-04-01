package main

import (
	"fmt"

	"github.com/katelynn620/gin-restful/pkg/api"
	"github.com/katelynn620/gin-restful/pkg/config"
	"github.com/katelynn620/gin-restful/pkg/util"
	"go.uber.org/zap"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	util.InitLogger(c.Debug)
	logger := zap.L().Sugar()
	defer logger.Sync()

	r := api.Init()

	port := fmt.Sprintf(":%v", c.Port)
	logger.Info("Server is running on port", port)
	r.Run(port)
}
