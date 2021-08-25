package model

// Recipient is a database object containing recipients information
type Recipient struct {
	ID    int64  `gorm:"primaryKey" json:"id" binding:"eq=0"`
	Name  string `gorm:"not null" json:"name" binding:"required,min=2"`
	Email string `json:"email"`
	Phone string `gorm:"index:idx_recipient_phone" json:"phone" binding:"numeric,min=8,max=15"`
}
