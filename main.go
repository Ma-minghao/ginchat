package main

import (
	"ginchat/router1"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	r := router1.Router()
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
