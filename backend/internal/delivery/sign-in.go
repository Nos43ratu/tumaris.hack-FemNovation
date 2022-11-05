package delivery

import (
	"errors"
	"log"
	"net/http"
	"net/mail"
	"unicode"

	"github.com/gin-gonic/gin"

	"tumaris.hack-FemNovation/backend/internal/models"
)

func (h *Handler) OK(c *gin.Context) {
	h.logger.Info("[SUCCESS]: options")
	c.JSON(200, CreateResponse(models.StatusOK, nil, nil))
}

func (h *Handler) SignIn(c *gin.Context) {
	input := &models.User{}
	if err := parseJSON(c, input); err != nil {
		h.logger.Errorf("[ERROR]: [%s] bad request error: %s", input.Email, err.Error())
		c.JSON(400, CreateResponse(models.StatusError, nil, models.ErrInvalidInput))
		return
	}

	if !invalidEmail(input.Email) {
		h.logger.Errorf("[ERROR]: bad request email: %s", input.Email)
		c.JSON(400, CreateResponse(models.StatusError, nil, models.ErrInvalidInput))
		return
	}

	tokens, err := h.service.Auth.SignIn(input)
	if err != nil {
		h.logger.Errorf("[ERROR]: [%s] error: %s", input.Email, err.Error())
		if errors.Is(err, models.ErrUnauthorized) {
			c.JSON(http.StatusUnauthorized, CreateResponse(models.StatusError, nil, models.ErrUnauthorized))
			return
		}
		c.JSON(500, CreateResponse(models.StatusError, nil, models.ErrInternalServer))
		return
	}

	user, err := h.service.Auth.GetUserByEmail(input.Email)
	if err != nil {
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

	h.logger.Infof("[SUCCESS]: signed in %s", input.Email)
	c.JSON(200, CreateResponse(models.StatusOK, user.Role, nil))
}

func invalidEmail(email string) bool {
	runes := []rune(email)
	if len(runes) < 4 || len(runes) > 320 {
		log.Println("len", len(runes))
		return false
	}

	for _, r := range runes {
		if !unicode.Is(unicode.Latin, r) && !unicode.IsDigit(r) && r != '@' && r != '_' && r != '-' && r != '.' {
			log.Println("rune is not valid", r)
			return false
		}
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		log.Println("invalid email: ", err)
		return false
	}

	return true
}
