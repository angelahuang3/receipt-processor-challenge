package tests

import (
	"receipt/app/models"
	"receipt/app/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test Receipt for Target Receipt
func TestProcessReceiptTarget(t *testing.T) {
	//store := repository.NewInMemoryDB()
	service := services.NewReceiptService()

	mockReceipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Total:        "35.35",
		Items: []models.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		},
	}

	receiptID, err := service.ProcessReceipt(mockReceipt)
	assert.NoError(t, err)
	assert.NotEmpty(t, receiptID)

	// Validate Points Calculation
	points, err := service.GetPoints(receiptID)
	assert.NoError(t, err)
	assert.Equal(t, 28, points) // Expected to be 28 points
}

// Test Receipt for M&M Corner Market
func TestProcessReceiptMarket(t *testing.T) {
	//store := repository.NewInMemoryDB()
	service := services.NewReceiptService()

	mockReceipt := models.Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Total:        "9.00",
		Items: []models.Item{
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
		},
	}

	receiptID, err := service.ProcessReceipt(mockReceipt)
	assert.NoError(t, err)
	assert.NotEmpty(t, receiptID)

	// Validate Points Calculation
	points, err := service.GetPoints(receiptID)
	assert.NoError(t, err)
	assert.Equal(t, 109, points) // Expected to be 109 points
}
