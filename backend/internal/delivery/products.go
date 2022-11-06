package delivery

import (
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
	}else {
		input.ProductID, err = strconv.Atoi(productID)
		if err != nil {
			h.logger.Errorf("[ERROR]: [%s] error during conversion: %s", productID, err.Error())
	
			c.JSON(500, CreateResponse(models.StatusError, nil, models.ErrInternalServer))
			return
		}

		id, err := h.service.Products.UpdateProduct(input)
		if err != nil {
			c.JSON(500, CreateResponse(models.StatusError, nil, models.ErrInternalServer))
			return
		}
		h.logger.Infof("[SUCCESS]: updated product with id %s", id)
		c.JSON(200, CreateResponse(models.StatusOK, id, nil))
	}
}

func (h *Handler)DeleteProduct(c *gin.Context){
	var err error
	// categoryID := c.Param("category_id")
	productID := c.Param("product_id")
	var id int
	if productID == "" {
		h.logger.Errorf("[ERROR]: bad request product id: %s", productID)
		c.JSON(400, CreateResponse(models.StatusError, nil, models.ErrInvalidInput))
		return
	}else {
		id, err = strconv.Atoi(productID)
		if err != nil {
			h.logger.Errorf("[ERROR]: [%s] error during conversion: %s", productID, err.Error())
	
			c.JSON(500, CreateResponse(models.StatusError, nil, models.ErrInternalServer))
			return
		}

		err := h.service.Products.DeleteProduct(id)
		if err != nil {
			if err == models.ErrNoRows {
				c.JSON(400, CreateResponse(models.StatusError, nil, models.ErrInvalidInput))
				return
			}
			c.JSON(500, CreateResponse(models.StatusError, nil, models.ErrInternalServer))
			return
		}
		h.logger.Infof("[SUCCESS]: deleted product with id %s", id)
		c.JSON(200, CreateResponse(models.StatusOK, id, nil))
	}
}

func (h *Handler)GetProduct(c *gin.Context){
	var err error
	// categoryID := c.Param("category_id")
	productID := c.Param("product_id")
	var id int
	if productID == "" {
		h.logger.Errorf("[ERROR]: bad request product id: %s", productID)
		c.JSON(400, CreateResponse(models.StatusError, nil, models.ErrInvalidInput))
		return
	} else {
		id, err = strconv.Atoi(productID)
		if err != nil {
			h.logger.Errorf("[ERROR]: [%s] error during conversion: %s", productID, err.Error())
	
			c.JSON(500, CreateResponse(models.StatusError, nil, models.ErrInternalServer))
			return
		}

		res, err := h.service.Products.GetProduct(id)
		if err != nil {
			if err == models.ErrNoRows {
				c.JSON(400, CreateResponse(models.StatusError, nil, models.ErrInvalidInput))
				return
			}
			c.JSON(500, CreateResponse(models.StatusError, nil, models.ErrInternalServer))
			return
		}
		h.logger.Infof("[SUCCESS]: deleted product with id %s", id)
		c.JSON(200, CreateResponse(models.StatusOK, res, nil))
	}
}

func (h *Handler)GetProducts(c *gin.Context){
	var err error
	categoryID := c.Param("category_id")
	var id int
	if categoryID == "" {
		res, err := h.service.Products.GetProducts()
		if err != nil {
			if err == models.ErrNoRows {
				c.JSON(400, CreateResponse(models.StatusError, nil, models.ErrInvalidInput))
				return
			}
			c.JSON(500, CreateResponse(models.StatusError, nil, models.ErrInternalServer))
			return
		}
		h.logger.Infof("[SUCCESS]: gets products")
		c.JSON(200, CreateResponse(models.StatusOK, res, nil))
	} else {
		id, err = strconv.Atoi(categoryID)
		if err != nil {
			h.logger.Errorf("[ERROR]: [%s] error during conversion: %s", categoryID, err.Error())
	
			c.JSON(500, CreateResponse(models.StatusError, nil, models.ErrInternalServer))
			return
		}

		res, err := h.service.Products.GetProductsByCategory(id)
		if err != nil {
			if err == models.ErrNoRows {
				c.JSON(400, CreateResponse(models.StatusError, nil, models.ErrInvalidInput))
				return
			}
			c.JSON(500, CreateResponse(models.StatusError, nil, models.ErrInternalServer))
			return
		}
		h.logger.Infof("[SUCCESS]: get products with category id %s", id)
		c.JSON(200, CreateResponse(models.StatusOK, res, nil))
	}
}