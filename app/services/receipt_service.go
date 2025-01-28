package services

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"unicode"

	"receipt/app/models"
	"receipt/app/repository"

	"github.com/google/uuid"
)

type ReceiptService struct {
	DB *repository.InMemoryDB
}

func NewReceiptService() *ReceiptService {
	return &ReceiptService{DB: repository.NewInMemoryDB()}
}

func (s *ReceiptService) ProcessReceipt(r models.Receipt) (string, error) {
	r.ID = uuid.New().String()

	totalPoints := 0
	totalPoints += CountRetailerPoints(r.Retailer)
	totalPoints += CountRoundDollar(r.Total)
	totalPoints += CountQuarter(r.Total)
	totalPoints += CountItemPair(r.Items)
	totalPoints += CountItemDescription(r.Items)
	totalPoints += CountOddDay(r.PurchaseDate)
	totalPoints += CountTime(r.PurchaseTime)

	s.DB.SaveReceipt(r)
	s.DB.SavePoints(r.ID, totalPoints)

	return r.ID, nil
}

func (s *ReceiptService) GetPoints(id string) (int, error) {
	points, err := s.DB.GetPoints(id)
	if err != nil {
		return 0, errors.New("no receipt found for that id")
	}
	return points, nil
}

// Count alphanumeric characters in the retailer name
func CountRetailerPoints(retailer string) int {
	count := 0
	for _, char := range retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// 50 points if the total is a round dollar amount with no cents.
func CountRoundDollar(total string) int {
	if strings.HasSuffix(total, ".00") {
		return 50
	}
	return 0
}

// 25 points if the total is a multiple of 0.25
func CountQuarter(total string) int {
	floatTotal, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return 0
	}
	if math.Mod(floatTotal, 0.25) == 0 {
		return 25
	}
	return 0
}

// 5 points for every two items on the receipt
func CountItemPair(items []models.Item) int {
	return (len(items) / 2) * 5
}

// If the length of the item description is a multiple of 3,
// multiply the price by 0.2 and round up.
func CountItemDescription(items []models.Item) int {
	points := 0
	for _, item := range items {
		trimmedStr := strings.TrimSpace(item.ShortDescription)
		if len(trimmedStr)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err == nil {
				points += int(math.Ceil(price * 0.2))
			}
		}
	}
	return points
}

// 6 points if the day in the purchase date is odd
func CountOddDay(purchaseDate string) int {
	dateParts := strings.Split(purchaseDate, "-")
	if len(dateParts) < 3 {
		return 0
	}
	day, err := strconv.Atoi(dateParts[2])
	if err != nil {
		return 0
	}
	if day%2 != 0 {
		return 6
	}
	return 0
}

// 10 points if the time of purchase is after 2:00pm and before 4:00pm
func CountTime(purchaseTime string) int {
	timeParts := strings.Split(purchaseTime, ":")
	if len(timeParts) < 2 {
		return 0
	}
	hour, err := strconv.Atoi(timeParts[0])
	if err != nil {
		return 0
	}
	if hour >= 14 && hour < 16 {
		return 10
	}
	return 0
}
