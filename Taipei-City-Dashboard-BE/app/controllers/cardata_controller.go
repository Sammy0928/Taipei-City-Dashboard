package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCardata handles the request to get car data.

func GetCardata(c *gin.Context) {

    c.JSON(http.StatusOK, gin.H{

        "message": "Car data retrieved successfully",

    })

}
