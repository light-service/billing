package resource

import "time"

type Resource struct {
	ID                    int
	Key                   string
	Name                  string
	RecordUnit            int
	RecordUnitName        string
	DisplayUnitConversion int
	DisplayUnit           int
	DisplayUnitName       string
	DisplayUnitPrice      int
	ProductID             string
	CreatedAt             time.Time
}
