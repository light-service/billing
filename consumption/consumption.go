package consumption

import "time"

type Consumption struct {
	ID                  int
	ProductID           int
	ResourceKey         string
	ResourceID          string
	StartAt             time.Time
	EndAt               time.Time
	Accrual             int
	ResourcePackAccrual int
	UserKey             string
	Completed           bool
}
