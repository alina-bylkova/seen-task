package model

// Video is a database object
type Video struct {
	Id          string `gorm:"primaryKey"`
	CompanyName string
}
