package model

// Video is a database object
type Video struct {
	ID          int64  `gorm:"primaryKey" json:"id"`
	CompanyName string `json:"company_name"`
}
