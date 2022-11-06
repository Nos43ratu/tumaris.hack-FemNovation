package delivery

import (
	"github.com/gin-gonic/gin"
	"tumaris.hack-FemNovation/backend/internal/models"
)

func (h *Handler) GetUser(c *gin.Context) {
	user := &models.User{}

	if err := parseJSON(c, user); err != nil {
		h.logger.Errorf("[ERROR]: [%s] bad request error: %s", err.Error())
		c.JSON(400, CreateResponse(models.StatusError, nil, models.ErrInvalidInput))
		return
	}

	info, err := h.service.Users.GetByEmail(user.Email)
	if err != nil {
		h.logger.Errorf("[ERROR]: error getting user", err)
		c.JSON(500, CreateResponse(models.StatusError, nil, models.ErrInternalServer))
		return
	}

	h.logger.Infof("[SUCCESS]: getting user %+v", info)
	c.JSON(200, CreateResponse(models.StatusOK, info, nil))
}
