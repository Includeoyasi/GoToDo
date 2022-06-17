package handler

import (
	"net/http"
	"strconv"

	"github.com/Includeoyasi/todo-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(gctx *gin.Context) {
	userId, err := getUserId(gctx)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(gctx.Param("id")) //getting params from uri
	if err != nil {
		NewErrorResponse(gctx, http.StatusBadRequest, "invalid id parametr")
		return
	}

	var input todo.TodoItem
	if err := gctx.BindJSON(&input); err != nil {
		NewErrorResponse(gctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.TodoItem.Create(userId, listId, input)
	if err != nil {
		NewErrorResponse(gctx, http.StatusInternalServerError, err.Error())
		return
	}

	gctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type getAllItemResponse struct {
	Data []todo.TodoItem `json:"data"`
}

func (h *Handler) getAllItems(gctx *gin.Context) {
	userId, err := getUserId(gctx)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(gctx.Param("id"))
	if err != nil {
		NewErrorResponse(gctx, http.StatusBadRequest, "invalid id parametr")
		return
	}

	items, err := h.service.TodoItem.GetAll(userId, listId)
	if err != nil {
		NewErrorResponse(gctx, http.StatusInternalServerError, err.Error())
		return
	}

	gctx.JSON(http.StatusOK, getAllItemResponse{
		Data: items,
	})

}

func (h *Handler) getItemById(gctx *gin.Context) {
	userId, err := getUserId(gctx)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(gctx.Param("id"))
	if err != nil {
		NewErrorResponse(gctx, http.StatusBadRequest, "invalid id parametr")
		return
	}

	item, err := h.service.TodoItem.GetById(userId, itemId)
	if err != nil {
		NewErrorResponse(gctx, http.StatusInternalServerError, err.Error())
		return
	}

	gctx.JSON(http.StatusOK, item)

}

func (h *Handler) updateItem(gctx *gin.Context) {
	userId, err := getUserId(gctx)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(gctx.Param("id"))
	if err != nil {
		NewErrorResponse(gctx, http.StatusBadRequest, "invalid id parametr")
		return
	}

	var input todo.UpdateTodoItemInput
	if err := gctx.BindJSON(&input); err != nil {
		NewErrorResponse(gctx, http.StatusBadRequest, "invalid params to update")
		return
	}

	err = h.service.TodoItem.Update(userId, itemId, input)
	if err != nil {
		NewErrorResponse(gctx, http.StatusInternalServerError, err.Error())
		return
	}

	gctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) deleteItem(gctx *gin.Context) {
	userId, err := getUserId(gctx)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(gctx.Param("id"))
	if err != nil {
		NewErrorResponse(gctx, http.StatusBadRequest, "invalid id parametr")
		return
	}

	err = h.service.TodoItem.Delete(userId, itemId)
	if err != nil {
		NewErrorResponse(gctx, http.StatusInternalServerError, err.Error())
		return
	}

	gctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
