package read

import "time"

type OrderRead struct {
	OrderId   int
	Price     float32
	StudentId int
	Paidat    time.Time
}
