package controller

import (
	"scheduler/module/user/repository"
	"scheduler/util"

	"github.com/gin-gonic/gin"
)

func Find(ctx *gin.Context) {
	var registred repository.UserReferences

	util.AcceptedOrNotStatusReturn(registred.Find(), ctx, registred.Users)
}

func Create(ctx *gin.Context) {
	var new repository.UserReferences

	if err := ctx.BindJSON(&new.User); err == nil && util.IsEmailValid(new.User.Email) {
		util.AcceptedOrNotStatusReturn(new.Create(), ctx, new.User)
	}
}

func Update(ctx *gin.Context) {
	var registred repository.UserReferences
	// ctx.Param("id")

	if err := ctx.BindJSON(&registred.User); err == nil && util.IsEmailValid(registred.User.Email) {
		util.AcceptedOrNotStatusReturn(registred.Update(), ctx, registred.User)
	}
}

func Delete(ctx *gin.Context) {
	var registred repository.UserReferences
	// ctx.Param("id")

	if err := ctx.BindJSON(&registred.User); err == nil && util.IsEmailValid(registred.User.Email) {
		util.AcceptedOrNotStatusReturn(registred.Delete(), ctx, registred.User)
	}
}
