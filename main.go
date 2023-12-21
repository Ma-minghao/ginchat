package main

import (
	"ginchat/router1"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()

	r := router1.Router()
	r.Run("localhost:8081")
}
