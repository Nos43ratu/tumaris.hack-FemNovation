package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tumaris.hack-FemNovation/backend/internal/models"
)

func (h *Handler) SignOut(c *gin.Context) {
	accessToken, err := c.Cookie("access_token")
	if err != nil {
		h.logger.Error("access token doesn't exists in cookies", err)
		c.JSON(http.StatusUnauthorized, CreateResponse(models.StatusError, nil, models.ErrUnauthorized))
		c.Abort()
		return
	}

	if accessToken == "" {
		h.logger.Error("access token doesn't exists in cookies", err)
		c.JSON(http.StatusUnauthorized, CreateResponse(models.StatusError, nil, models.ErrUnauthorized))
		c.Abort()
		return
	}

	ok, err := h.service.Auth.VerifyToken(accessToken)
	if err != nil {
		h.logger.Errorf("[ERROR]: [AuthAdminMiddleware] service error: %s", err.Error())
		c.JSON(500, CreateResponse(models.StatusError, nil, models.ErrInternalServer))
		return
	}

	if !ok {
		h.logger.Error("access token is not valid")
		c.JSON(http.StatusUnauthorized, CreateResponse(models.StatusError, nil, models.ErrUnauthorized))
		c.Abort()
		return
	}

	_, err = h.service.Auth.ParseTokenWithClaims(accessToken)
	if err != nil {
		h.logger.Error("err", err)
		c.JSON(http.StatusUnauthorized, CreateResponse(models.StatusError, nil, models.ErrUnauthorized))
		c.Abort()
		return
	}

	if err := h.service.Auth.SignOut(accessToken); err != nil {
		h.logger.Error("err", err)
		c.JSON(500, CreateResponse(models.StatusError, nil, models.ErrInternalServer))
		return
	}

	c.SetCookie(
		"access_token",
		accessToken,
		0,
		"/",
		h.tokenInfo.Domain,
		h.tokenInfo.Secure,
		h.tokenInfo.HttpOnly,
	)

	h.logger.Info("[SUCCESS]: signed out user")
	c.JSON(200, CreateResponse(models.StatusOK, nil, nil))
}
