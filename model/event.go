package model

// Event is a database object
type Event struct {
	RecipientID int64 `gorm:"not null" json:"recipient_id"`
	VideoID     int64 `gorm:"not null" json:"video_id"`
	LpHits      int64 `gorm:"not null;default:0" json:"lp_hits"`
	VideoPlays  int64 `gorm:"not null;default:0" json:"video_plays"`
}
