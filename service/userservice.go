package service

import (
	"ginchat/models"
	"github.com/gin-gonic/gin"
)
// @BasePath /api/v1
// GetUserList
// @Tags 首页
// @Success 200 {string} json{"code","message"}
// PingExample godoc
// @Schemes
// @Description do ping
// @Accept json
// @Produce json
// @Router /user/getUseList [get]

func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})

}
