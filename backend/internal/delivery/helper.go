package delivery

import "github.com/gin-gonic/gin"

func CreateResponse(status int, data interface{}, err error) gin.H {
	var errMessage gin.H

	if err != nil {
		errMessage = gin.H{
			"message": err.Error(),
		}
	} else {
		errMessage = nil
	}

	return gin.H{
		"response": gin.H{
			"status": status,
			"data":   data,
		},
		"error": errMessage,
	}
}
