package order

import "time"

type CouponTemplate struct {
	ID         int
	Name       string
	ProductIDs []int
}

type Coupon struct {
	ID         int
	Name       string
	Desc       string
	Amount     int64
	LeftAmount int64
	ProductIDs []int
	UserKey    string
	StartAt    time.Time
	EndAt      time.Time
	CreatedAt  time.Time
}

type CouponDetail struct {
	ID          int64
	ProductID   int
	ResourceKey string
	ResourceID  string
	CouponID    int
	UsedAmount  int64
	LeftAmount  int64
	UserKey     string
	CreatedAt   time.Time
}
