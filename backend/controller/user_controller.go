package controller

import (
	"backend/middleware"
	"backend/model"
	"backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (c *Controller) UserRegister(context *gin.Context) {
	var user model.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		return
	}
	c.dao.CreateUser(user)
	context.JSON(200, "register success")
}

type CodeRequest struct {
	Code string `json:"code"`
}

func (c *Controller) UserLogin(context *gin.Context) {
	var codeRequest CodeRequest
	err := context.BindJSON(&codeRequest)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(codeRequest.Code, "2")
	wxRes, err := utils.GetSession("wx8436dc91cc648b2c", "553edb39f39ed53429a18eaaebaaeff2", codeRequest.Code)
	if err != nil {
		fmt.Println(err)
	}
	tokenString, err := c.Jwt.CreateToken(middleware.NewClaims(wxRes.UnionID))
	if err != nil {
		fmt.Println(err)
	}
	context.JSON(200, gin.H{
		"token": tokenString,
	})
}

func (c *Controller) GetEvaluationsByUserID(context *gin.Context) {
	idString := context.Param("id")
	id := utils.StringToUint(idString)
	evaluations := c.dao.GetEvaluationByUserID(id)
	context.JSON(200, evaluations)
}
