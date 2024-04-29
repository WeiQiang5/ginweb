package api

import (
	"ginWeb/models"
	"ginWeb/pkg/e"
	"ginWeb/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

type auth struct {
	Id       int    `valid:"Required"; MaxSize(50)`
	Username string `valid:"Required"; MaxSize(50)`
}

func GetAuth(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()
	username := c.Query("username")

	valid := validation.Validation{}
	a := auth{Id: id, Username: username}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if ok {
		isExist := models.CheckAuth(id, username)
		if isExist {
			token, err := util.GenerateToken(id, username)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  e.GetMsg(code),
	})
}
