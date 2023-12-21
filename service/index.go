package service

import (
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// GetUserList
// @Tags 首页
// @Success 200 {string} json{"code","message"}
// PingExample godoc
// @Summary 首页
// @Schemes
// @Description do ping
// @Accept json
// @Produce json
// @Router /user/getUseList [get]

func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome!!",
	})

}
