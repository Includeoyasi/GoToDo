package handler

import (
	"net/http"
	"strconv"

	"github.com/Includeoyasi/todo-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(gctx *gin.Context) {
	UserId, err := getUserId(gctx)
	if err != nil {
		return
	}
	var input todo.TodoList

	if err := gctx.BindJSON(&input); err != nil {
		NewErrorResponse(gctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.TodoList.Create(UserId, input)
	if err != nil {
		NewErrorResponse(gctx, http.StatusInternalServerError, err.Error())
		return
	}
	gctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllList(gctx *gin.Context) {
	UserId, err := getUserId(gctx)
	if err != nil {
		return
	}

	lists, err := h.service.TodoList.GetAll(UserId)
	if err != nil {
		NewErrorResponse(gctx, http.StatusInternalServerError, err.Error())
		return
	}

	gctx.JSON(http.StatusOK, getAllListResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(gctx *gin.Context) {
	UserId, err := getUserId(gctx)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(gctx.Param("id")) //getting params from uri
	if err != nil {
		NewErrorResponse(gctx, http.StatusBadRequest, "invalid id parametr")
		return
	}

	list, err := h.service.TodoList.GetById(UserId, id)
	if err != nil {
		NewErrorResponse(gctx, http.StatusInternalServerError, err.Error())
		return
	}

	gctx.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(gctx *gin.Context) {
	UserId, err := getUserId(gctx)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(gctx.Param("id")) //getting params from uri
	if err != nil {
		NewErrorResponse(gctx, http.StatusBadRequest, "invalid id parametr")
		return
	}

	var input todo.UpdateTodoListInput
	if err := gctx.BindJSON(&input); err != nil {
		NewErrorResponse(gctx, http.StatusBadRequest, "invalid params to update")
		return
	}

	if err := h.service.TodoList.Update(UserId, id, input); err != nil {
		NewErrorResponse(gctx, http.StatusBadRequest, err.Error())
		return
	}

	gctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteList(gctx *gin.Context) {
	UserId, err := getUserId(gctx)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(gctx.Param("id"))
	if err != nil {
		NewErrorResponse(gctx, http.StatusBadRequest, "invalid id parametr")
		return
	}

	if err := h.service.TodoList.Delete(UserId, id); err != nil {
		NewErrorResponse(gctx, http.StatusInternalServerError, err.Error())
		return
	}

	gctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
