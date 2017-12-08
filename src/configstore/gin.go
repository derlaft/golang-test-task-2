package main

import (
	"net/http"

	"configstore/models"

	"github.com/gin-gonic/gin"
)

type ViewFunc func(interface{}) (interface{}, error)

func errorObject(err string) map[string]interface{} {
	return map[string]interface{}{
		"error": err,
	}
}

// Use our custom views with gin
func GinAdapter(input interface{}, view ViewFunc) gin.HandlerFunc {

	return func(c *gin.Context) {

		req := models.NewStruct(input)

		err := c.BindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorObject("Bad input"))
			return
		}

		res, err := view(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorObject(err.Error()))
			return
		}

		c.JSON(200, res)
	}
}
