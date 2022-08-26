package section

import (
	"scheduler/common/util"
	"scheduler/core/user"

	"github.com/gin-gonic/gin"
)

func Find(ctx *gin.Context) {
	var registred user.UserReferences

	if err := ctx.BindJSON(&registred.User); err == nil && util.IsEmailValid(registred.User.Email) {
		util.AcceptedOrNotStatusReturn(registred.Create(), ctx, registred.User)
	}
}

func Create(ctx *gin.Context) {
	var new user.UserReferences

	if err := ctx.BindJSON(&new.User); err == nil && util.IsEmailValid(new.User.Email) {
		util.AcceptedOrNotStatusReturn(new.Create(), ctx, new.User)
	}
}
