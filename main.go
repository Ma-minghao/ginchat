package main

import (
	rt "ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()

	r := rt.Router()
	r.Run("localhost:8081")
}
