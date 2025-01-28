package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"receipt/app/models"
	"receipt/app/router"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test Gin Server
func setupTestServer() *gin.Engine {
	r := gin.Default()
	router.SetupRoutes(r)
	return r
}

// Test Process Receipt - Target Receipt
func TestProcessReceiptAPI_Target(t *testing.T) {
	r := setupTestServer()

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

	jsonData, _ := json.Marshal(mockReceipt)
	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.NotEmpty(t, response["receipt_id"])

	// Fetch Points
	req, _ = http.NewRequest("GET", "/receipts/"+response["receipt_id"]+"/points", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var pointsResponse map[string]int
	json.Unmarshal(w.Body.Bytes(), &pointsResponse)
	assert.Equal(t, 28, pointsResponse["points"]) // Expected 28 points
}

// Test Process Receipt - M&M Corner Market
func TestProcessReceiptAPI_Market(t *testing.T) {
	r := setupTestServer()

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

	jsonData, _ := json.Marshal(mockReceipt)
	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.NotEmpty(t, response["receipt_id"])

	// Fetch Points
	req, _ = http.NewRequest("GET", "/receipts/"+response["receipt_id"]+"/points", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var pointsResponse map[string]int
	json.Unmarshal(w.Body.Bytes(), &pointsResponse)
	assert.Equal(t, 109, pointsResponse["points"]) // Expected 109 points
}
