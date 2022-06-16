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

type SingInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) singIn(gctx *gin.Context) {
	var input SingInInput

	if err := gctx.BindJSON(&input); err != nil {
		NewErrorResponse(gctx, http.StatusBadRequest, err.Error())
		return
	}

	toker, err := h.service.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		NewErrorResponse(gctx, http.StatusInternalServerError, err.Error())
	}

	gctx.JSON(http.StatusOK, map[string]interface{}{
		"toker": toker,
	})
}
