package delivery

import (
	"github.com/gin-gonic/gin"
	"tumaris.hack-FemNovation/backend/internal/models"
)

func (h *Handler) CreateOrder(c *gin.Context) {
	order := &models.Order{}

	if err := parseJSON(c, order); err != nil {
		h.logger.Errorf("[ERROR]: [%s] bad request error: %s", err.Error())
		c.JSON(400, CreateResponse(models.StatusError, nil, models.ErrInvalidInput))
		return
	}

	if err := h.service.Order.Create(order); err != nil {
		h.logger.Errorf("[ERROR]: error creating order", err)
		c.JSON(500, CreateResponse(models.StatusError, nil, models.ErrInternalServer))
		return
	}

	h.logger.Infof("[SUCCESS]: created order %+v", order)
	c.JSON(200, CreateResponse(models.StatusOK, nil, nil))
}

func (h *Handler) UpdateOrder(c *gin.Context) {
	order := &models.Order{}

	if err := parseJSON(c, order); err != nil {
		h.logger.Errorf("[ERROR]: [%s] bad request error: %s", err.Error())
		c.JSON(400, CreateResponse(models.StatusError, nil, models.ErrInvalidInput))
		return
	}

	if err := h.service.Order.Update(order); err != nil {
		h.logger.Errorf("[ERROR]: error updating order", err)
		c.JSON(500, CreateResponse(models.StatusError, nil, models.ErrInternalServer))
		return
	}

	h.logger.Infof("[SUCCESS]: updated order %+v", order)
	c.JSON(200, CreateResponse(models.StatusOK, nil, nil))
}

func (h *Handler) GetAllOrders(c *gin.Context) {
	orders, err := h.service.Order.GetAll()
	if err != nil {
		h.logger.Errorf("[ERROR]: error getting ALL orders", err)
		c.JSON(500, CreateResponse(models.StatusError, nil, models.ErrInternalServer))
		return
	}

	h.logger.Infof("[SUCCESS]: getting all orders %+v", orders)
	c.JSON(200, CreateResponse(models.StatusOK, orders, nil))
}
