package delivery

import (
	// "errors"
	// "log"
	// "net/http"
	// "net/mail"
	// "unicode"
	"strconv"
	"github.com/gin-gonic/gin"
	"tumaris.hack-FemNovation/backend/internal/models"
)

func (h *Handler)CreateOrUpdateProduct(c *gin.Context){
	var err error
	categoryID := c.Param("category_id")
	productID := c.Param("product_id")
	input := &models.Product{}
	if err := parseJSON(c, input); err != nil {
		h.logger.Errorf("[ERROR]: [%s] bad request error: %s", "product create", err.Error())
		c.JSON(400, CreateResponse(models.StatusError, nil, models.ErrInvalidInput))
		return
	}
	
	input.CategoryID, err = strconv.Atoi(categoryID)
	if err != nil {
		h.logger.Errorf("[ERROR]: [%s] error during conversion: %s", categoryID, err.Error())

		c.JSON(500, CreateResponse(models.StatusError, nil, models.ErrInternalServer))
		return
	}
	if productID == "" {
		id, err := h.service.Products.CreateProduct(input)
		if err != nil {
			c.JSON(500, CreateResponse(models.StatusError, nil, models.ErrInternalServer))
			return
		}
		h.logger.Infof("[SUCCESS]: created product with id %s", id)
		c.JSON(200, CreateResponse(models.StatusOK, id, nil))
	}
	input.ProductID, err = strconv.Atoi(productID)
	if err != nil {
		h.logger.Errorf("[ERROR]: [%s] error during conversion: %s", productID, err.Error())

		c.JSON(500, CreateResponse(models.StatusError, nil, models.ErrInternalServer))
		return
	}
	// here update product
}