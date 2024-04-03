package controllers

import (
	"github.com/cristopherecruz/erp-do-brasil-fiscal/models"
	"github.com/gin-gonic/gin"
	"github.com/hooklift/gowsdl/soap"
	"net/http"
)

func WebServiceAutorizacao(c *gin.Context) {

	/*autorizacao := &models.Autorizacao{}

	err := c.ShouldBindJSON(&autorizacao)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}*/

	//soap.WithHTTPClient(httpClient)
	client := soap.NewClient("https://nfe-homologacao.svrs.rs.gov.br/ws/NfeAutorizacao/NFeAutorizacao4.asmx")

	autorizacao4Soap := models.NewNFeAutorizacao4Soap(client)
	autorizacao4Soap.NfeAutorizacaoLote(&models.NfeDadosMsg{})

	c.JSON(http.StatusOK, nil)
}
