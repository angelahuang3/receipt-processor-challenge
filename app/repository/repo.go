package repository

import "receipt/app/models"

type ReceiptRepository interface {
	SaveReceipt(receipt models.Receipt) error
	SavePoints(id string, points int) int
	GetPoints(id string) int
}
