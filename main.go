package main

import (
	"library/config"
	"library/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	r := gin.Default()
	routes.Routes(r)
	r.Run()
}
