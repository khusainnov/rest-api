package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/khusainnov/rest-api/internal/user"
	"net/http"
)

func (h *Handler) SignUp(ctx *gin.Context) {
	var input user.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) SignIn(ctx *gin.Context) {

}
