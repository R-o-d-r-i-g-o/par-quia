package controller

import (
	"scheduler/module/home/user/repository"
	"scheduler/util"

	"github.com/gin-gonic/gin"
)

func Find(ctx *gin.Context) {
	var registred repository.UserReferences

	if err := ctx.BindJSON(&new.User); err == nil && util.IsEmailValid(new.User.Email) {
		util.AcceptedOrNotStatusReturn(new.Create(), ctx, new.User)
	}
}

func Create(ctx *gin.Context) {
	var new repository.UserReferences

	if err := ctx.BindJSON(&new.User); err == nil && util.IsEmailValid(new.User.Email) {
		util.AcceptedOrNotStatusReturn(new.Create(), ctx, new.User)
	}
}
