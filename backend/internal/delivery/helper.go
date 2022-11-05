package delivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const maxBodySize = 1024

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

func parseJSON(c *gin.Context, input interface{}) error {
	r := c.Request
	if r.Header.Get("Content-Type") != "" {
		value := r.Header.Get("Content-Type")
		if value != "application/json" {
			return errors.New("Content-Type header is not application/json")
		}
	} else {
		return errors.New("Content-Type header is not application/json")
	}

	c.Request.Body = http.MaxBytesReader(c.Writer, r.Body, maxBodySize)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(input)
	if err != nil {
		var syntaxErr *json.SyntaxError
		var unmarshalTypeErr *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxErr):
			msg := "request body contains badly-formed JSON"
			return errors.New(msg)
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := "request body contains badly-formed JSON"
			return errors.New(msg)
		case errors.As(err, &unmarshalTypeErr):
			msg := fmt.Sprintf("invalid value for the %q field, inalid type for this field", unmarshalTypeErr.Field)
			return errors.New(msg)
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return errors.New(msg)
		case errors.Is(err, io.EOF):
			msg := "request body must not be empty"
			return errors.New(msg)
		case err.Error() == "http: request body too large":
			msg := "request body must not be larger than 1MB"
			return errors.New(msg)
		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		msg := "request body must only contain a single JSON object"
		return errors.New(msg)
	}

	return nil
}
