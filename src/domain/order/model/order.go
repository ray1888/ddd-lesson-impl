package model

import "time"

type Order struct {
	StudentId int
	Price     float32
	PaidAt    time.Time
	Orderer   int
}
