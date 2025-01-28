package repository

import (
	"errors"
	"receipt/app/models"
)

type InMemoryDB struct {
	Receipts         map[string]models.Receipt
	PointsPerReceipt map[string]int
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		Receipts:         make(map[string]models.Receipt),
		PointsPerReceipt: make(map[string]int),
	}
}

func (db *InMemoryDB) SaveReceipt(r models.Receipt) models.Receipt {
	db.Receipts[r.ID] = r
	return r
}

func (db *InMemoryDB) GetPoints(id string) (int, error) {
	points, exists := db.PointsPerReceipt[id]
	if !exists {
		return -1, errors.New("receipt not found")
	}
	return points, nil
}

func (db *InMemoryDB) SavePoints(id string, points int) {
	db.PointsPerReceipt[id] = points
}
