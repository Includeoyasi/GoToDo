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

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(gctx, http.StatusInternalServerError, err.Error())
	}

	gctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) singIn(ctx *gin.Context) {

}
