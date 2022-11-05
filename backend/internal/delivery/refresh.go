package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tumaris.hack-FemNovation/backend/internal/models"
)

func (h *Handler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		h.logger.Error("refresh token doesn't exists in cookies", err)
		c.JSON(http.StatusUnauthorized, CreateResponse(models.StatusError, nil, models.ErrUnauthorized))
		return
	}

	tokens, err := h.service.Auth.Refresh(refreshToken)
	if err != nil {
		h.logger.Error("err", err)
		c.JSON(500, CreateResponse(models.StatusError, nil, models.ErrInternalServer))
		return
	}

	c.SetCookie(
		"access_token",
		tokens.Access,
		int(h.tokenInfo.Access.TTL.Seconds()),
		"/",
		h.tokenInfo.Domain,
		h.tokenInfo.Secure,
		h.tokenInfo.HttpOnly,
	)

	c.SetCookie(
		"refresh_token",
		tokens.Refresh,
		int(h.tokenInfo.Refresh.TTL.Seconds()),
		"/refresh",
		h.tokenInfo.Domain,
		h.tokenInfo.Secure,
		h.tokenInfo.HttpOnly,
	)

	h.logger.Info("[SUCCESS]: refreshed user")
	c.JSON(200, CreateResponse(models.StatusOK, nil, nil))
}
