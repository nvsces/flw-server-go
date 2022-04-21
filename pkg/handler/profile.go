package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler) info(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.services.Profile.Info(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}