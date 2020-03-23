package v1

import (
	"blogWithGin/models"
	"blogWithGin/pkg/errCode"
	"blogWithGin/pkg/logging"
	"blogWithGin/pkg/setting"
	"blogWithGin/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)" json:"username"`
	Password string `valid:"Required; MaxSize(50)" json:"password"`
}

func GetAuth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	data := make(map[string]interface{})
	code := errCode.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = errCode.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = errCode.SUCCESS
				c.SetCookie("token", token, setting.TokenExpireTime*3600, "", setting.Domain, false, true)
			}
		} else {
			code = errCode.ERROR_AUTH
		}

	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errCode.GetMsg(code),
		"data": data,
	})
}
