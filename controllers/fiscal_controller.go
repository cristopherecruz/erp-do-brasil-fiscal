package controllers

import (
	"github.com/cristopherecruz/erp-do-brasil-fiscal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WebServiceAutorizacao(c *gin.Context) {

	autorizacao := &models.Autorizacao{}

	err := c.ShouldBindJSON(&autorizacao)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, autorizacao)
}
