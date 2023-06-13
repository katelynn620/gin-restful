// Documentation of Books API
//
//		   Simple Gin API
//
//	    Schemes: http
//	    BasePath: /v1
//	    Version: 1.0.0
//	    Contact: Test User <some_email@example.com> http://github.com/
//
//	    Consumes:
//	    - application/json
//
//	    Produces:
//	    - application/json
//
//	    Security:
//	    - basic
//
//	   SecurityDefinitions:
//	   basic:
//	     type: basic
//
// swagger:meta
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
