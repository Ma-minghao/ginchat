package service

import (
	"ginchat/models"
	"github.com/gin-gonic/gin"
)

// GetUserList
//	@Summary	所有用户
// @Schemes
// @Description do ping
//	@Tags		首页
// @Accept json
// @Produce json
//	@Success	200	{string}	json{"code","message"}
//	@Router		/user/getUserList [get]

func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
//	@Summary	新增用户
//	@Tags		用户模块
//	@Param		name query string false "用户名"
//	@Param		password query string false "密码"
//	@Param		repassword query string false "确认密码"
//	@Success	200	{string}	json{"code","message"}
//	@Router		/user/getUserList [get]

func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name =c.Query( "name")
	password := c.Query( "password")
	repassword := c.Query( "repassword")
	if password !=repassword {
		c.JSON(-1,gin.H{
			"message":"两次密码不一致",
		})
	}
	
	user.Password = password
	models.CreateUser(user)
	c.JSON(200,gin.H{
		"message":"新增用户成功",
	})
	
}