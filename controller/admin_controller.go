package controller

import (
	"github.com/gin-gonic/gin"
)

type TableData struct {
	TableColumnNum int
}

func HandleAdmin(ctx *gin.Context) {
	// store := ginsession.FromContext(ctx)
	// _, ok := store.Get("user")
	// if !ok {
	// 	// User is not logged in, redirect them to the login page
	// 	ctx.Redirect(http.StatusTemporaryRedirect, "/login")
	// 	return
	// }

	ctx.HTML(200, "admin.html", nil)
}