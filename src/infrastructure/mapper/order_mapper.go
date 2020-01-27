package mapper

import (
	"time"

	"github.com/jinzhu/gorm"
)

type OrderModel struct {
	gorm.Model
	StudentId int
	Price     float32
	PaidAt    time.Time
	Orderer   int
}
