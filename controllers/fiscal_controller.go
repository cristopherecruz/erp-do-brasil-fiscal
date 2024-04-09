package controllers

import (
	"crypto/tls"
	"github.com/cristopherecruz/erp-do-brasil-fiscal/models"
	"github.com/gin-gonic/gin"
	"github.com/hooklift/gowsdl/soap"
	"log"
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

	cert := lerCertificadoDigital()

	// Configura o certificado no TLS
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	client := soap.NewClient("https://nfe-homologacao.svrs.rs.gov.br/ws/NfeAutorizacao/NFeAutorizacao4.asmx", soap.WithTLS(tlsConfig))

	autorizacao4Soap := models.NewNFeAutorizacao4Soap(client)
	autorizacao4Soap.NfeAutorizacaoLote(&models.NfeDadosMsg{})

	c.JSON(http.StatusOK, nil)
}

func lerCertificadoDigital() tls.Certificate {

	cer, err := tls.LoadX509KeyPair("certificados_incluindo_intermediarios.pem", "chave_privada.pem")
	if err != nil {
		log.Println(err)
	}

	return cer
}
