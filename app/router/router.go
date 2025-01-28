package router

import (
	"receipt/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	receiptController := controllers.NewReceiptController()

	r.POST("/receipts/process", receiptController.ProcessReceipt)
	r.GET("/receipts/:id/points", receiptController.GetPoints)

}
