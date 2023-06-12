package main

import (
	"fmt"

	"github.com/katelynn620/gin-restful/pkg/api"
	"github.com/katelynn620/gin-restful/pkg/config"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	fmt.Printf("port: %v\n", c.Port)
	fmt.Printf("db_url: %v\n", c.DBUrl)

	r := api.Init()

	port := fmt.Sprintf(":%v", c.Port)
	r.Run(port)
}
