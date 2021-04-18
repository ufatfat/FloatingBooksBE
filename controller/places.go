package controller

import (
	"FloatingBooks/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPlaces (c *gin.Context) {
	places, err := db.GetPlaces()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": places,
	})
}