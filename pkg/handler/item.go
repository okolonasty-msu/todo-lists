package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createItem(c *gin.Context) {

	c.String(http.StatusOK, "hello item")
}

func (h *Handler) getAllItems(c *gin.Context) {

}

func (h *Handler) getItemById(c *gin.Context) {

}

func (h *Handler) updateItem(c *gin.Context) {

}

func (h *Handler) deleteItem(c *gin.Context) {

}
