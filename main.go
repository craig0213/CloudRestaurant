package main

import (
	"CloudRestaurant/controller"
	"CloudRestaurant/tool"

	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := tool.ParseConfig("./config/app.json")

	if err != nil {
		panic(err)
	}

	app := gin.Default()
	registerRouter(app)
	// app.Run(fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort))
	app.Run(fmt.Sprintf(":%s", cfg.AppPort))

}

func registerRouter(r *gin.Engine) {
	new(controller.HelloController).Router(r)

}
