package accounting

import "time"

type Journal struct {
	ID            int
	Name          string
	SubName       string
	Describe      string
	DebitAccount  string
	CreditAccount string
}

type JournalResource struct {
	ID         int
	ResourceID int
	JournalID  int
}

type JournalEntry struct {
	ID            int64
	TransactionID string
	Describe      string
	Debit         int64
	Credit        int64
	JournalID     int
	AccountKey    string
	UserKey       string
	CreatedAt     time.Time
}
