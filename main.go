package main

import (
	"fmt"
	"github.com/steven1213/go-douyin/common"
	"os"
)

func main() {
	fmt.Println("Gin-Web Douyin API Start...")

	env := os.Getenv("GIN_ENV")
	if env == "" {
		env = "local"
	}

	switch env {
	case "local":
		common.ConfigFile = "config/config.local.yaml"
	case "dev":
		common.ConfigFile = "config/config.dev.yaml"
	case "pre":
		common.ConfigFile = "config/config.pre.yaml"
	case "prod":
		common.ConfigFile = "config/config.prod.yaml"
	default:
		common.ConfigFile = "config/config.local.yaml"
	}
	common.Loading()
}
