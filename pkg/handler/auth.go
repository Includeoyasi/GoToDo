package handler

import (
	"net/http"

	"github.com/Includeoyasi/todo-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) singUp(gctx *gin.Context) {
	var input todo.User

	if err := gctx.BindJSON(&input); err != nil {
		NewErrorResponse(gctx, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) singIn(ctx *gin.Context) {

}
