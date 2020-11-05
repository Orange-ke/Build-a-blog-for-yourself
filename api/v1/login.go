package v1

import (
	"github.com/gin-gonic/gin"
	"myblog/middleware"
	"myblog/model"
	"myblog/utils/errmsg"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	var token string
	code := model.CheckLogin(data.Username, data.Password)

	if code == errmsg.SUCCESS {
		token, code = middleware.GenerateToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
		"token": token,
	})
}
