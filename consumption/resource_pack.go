package consumption

import "time"

type ResourcePackTemplate struct {
	ID          int
	Name        string
	ResourceKey string
	Amount      int
	Fixed       bool
	Hide        bool
	CreatedAt   time.Time
}

type ResourcePack struct {
	ID          int
	Name        string
	Desc        string
	ResourceKey string
	Amount      int
	LeftAmount  int
	UsedAmount  int
	Fixed       bool
	UserKey     string
	StartAt     time.Time
	EndAt       time.Time
	CreatedAt   time.Time
}

type ResourcePackDetail struct {
	ID             int
	ProductID      int
	ResourceKey    string
	ResourceID     string
	UseAmount      int
	LeftAmount     int
	ResourcePackID int
	UserKey        string
	CreatedAt      time.Time
}
