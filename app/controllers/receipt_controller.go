package controllers

import (
	"log"
	"net/http"

	"receipt/app/models"
	"receipt/app/services"

	"github.com/gin-gonic/gin"
)

// Handles all API endpoints
type ReceiptController struct {
	Service *services.ReceiptService
}

// Initializes the controller with the service
func NewReceiptController() *ReceiptController {
	return &ReceiptController{Service: services.NewReceiptService()}
}

// Handles receipt submission and processing
func (rc *ReceiptController) ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt

	// Validate JSON format
	if err := c.ShouldBindJSON(&receipt); err != nil {
		log.Println("Invalid receipt format:", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid receipt format. Please check your JSON input."})
		return
	}

	// Process receipt and generate id
	id, err := rc.Service.ProcessReceipt(receipt)
	if err != nil {
		log.Println("Error processing receipt:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to process the receipt at this time."})
		return
	}

	log.Printf("Receipt processed successfully. ID: %s\n", id)
	c.JSON(http.StatusCreated, gin.H{"receipt_id": id})
}

// get the points associated with a given receipt ID
func (rc *ReceiptController) GetPoints(c *gin.Context) {
	id := c.Param("id")

	points, err := rc.Service.GetPoints(id)
	if err != nil {
		log.Printf("Receipt ID not found: %s\n", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for the provided ID."})
		return
	}

	log.Printf("Points retrieved for receipt ID %s: %d\n", id, points)
	c.JSON(http.StatusOK, gin.H{"receipt_id": id, "points": points})
}

func (rc *ReceiptController) CheckStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Receipt processing service is running."})
}
