package main

import (
	"github.com/eduardonakaidev/go-tasks/config"
	"github.com/eduardonakaidev/go-tasks/router"
	"github.com/eduardonakaidev/go-tasks/util"
)

func init () {
	err := util.LoadEnv()
	if err != nil {
		panic(err)
	}
}
func main() {
	config.ConnectPostgresDB()
	router.HandlerRequest()
}