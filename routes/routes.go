package routes

import (
	"github.com/cristopherecruz/erp-do-brasil-fiscal/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {

	r := gin.Default()

	//NFC-e
	r.POST("/fiscal/autorizacao", controllers.WebServiceAutorizacao)

	err := r.Run()
	if err != nil {
		log.Panic(err.Error())
	}

}
