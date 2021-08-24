package model

// Recipient is a database object
type Recipient struct {
	Id    string `gorm:"primaryKey"`
	Name  string
	Email string
	Phone string
}
