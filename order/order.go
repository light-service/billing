package order

import "time"

type Order struct {
	ID                int
	OrderType         int
	ProductID         int
	ResourceKey       string
	ResourceID        string
	TotalFee          int
	PayByDiscount     int
	PayByCoupon       int
	PayByAccounting   int
	NeedPay           int
	Paid              bool
	ConsumptionID     int
	AccountingTransID string
	UserKey           string
	UpdatedAt         time.Time
	CreatedAt         time.Time
}
