package model

// Recipient is a database object
type Recipient struct {
	ID    int64  `gorm:"primaryKey" json:"id" binding:"eq=0"`
	Name  string `gorm:"not null" json:"name" binding:"required"`
	Email string `json:"email" binding:"email"`
	Phone string `json:"phone" binding:"numeric,min=8,max=15"` //Todo: check numeric validation in string
}
