package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nubesFilius/bselling-go-oauth-api.git/src/domain/access_token"
	"github.com/nubesFilius/bselling-go-oauth-api.git/src/utils/errors"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	accessToken , err := h.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusNotImplemented, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var token access_token.AccessToken
	if err := c.ShouldBindJSON(&token); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	if err := h.service.Create(token); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, token)
}