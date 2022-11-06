package delivery

import (
	"github.com/gin-gonic/gin"
	"tumaris.hack-FemNovation/backend/internal/models"
)

func (h *Handler) GetUser(c *gin.Context) {
	email := c.Param("email")

	info, err := h.service.Users.GetByEmail(email)
	if err != nil {
		h.logger.Errorf("[ERROR]: error getting user", err)
		c.JSON(500, CreateResponse(models.StatusError, nil, models.ErrInternalServer))
		return
	}

	h.logger.Infof("[SUCCESS]: getting user %+v", info)
	c.JSON(200, CreateResponse(models.StatusOK, info, nil))
}
