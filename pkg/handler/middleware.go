package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) userIdentity(gctx *gin.Context) {
	header := gctx.GetHeader("Authorization")
	if header == "" {
		NewErrorResponse(gctx, http.StatusUnauthorized, "header is empty")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(gctx, http.StatusUnauthorized, "header is broken")
		return
	}

	token := headerParts[1]
	userId, err := h.service.ParseToken(token)
	if err != nil {
		NewErrorResponse(gctx, http.StatusUnauthorized, err.Error())
		return
	}

	gctx.Set("UserCtx", userId)
}

func getUserId(gctx *gin.Context) (int, error) {
	id, ok := gctx.Get("UserCtx")
	if !ok {
		NewErrorResponse(gctx, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int)
	if !ok {
		NewErrorResponse(gctx, http.StatusInternalServerError, "userId is not Int")
		return 0, errors.New("userId is not Int")
	}
	return idInt, nil
}
