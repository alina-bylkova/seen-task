package model

// Event is a database object containing events infomation
type Event struct {
	RecipientID int64 `gorm:"primaryKey" json:"recipient_id"`
	VideoID     int64 `gorm:"primaryKey" json:"video_id"`
	LpHits      int64 `gorm:"not null;default:0" json:"lp_hits"`
	VideoPlays  int64 `gorm:"not null;default:0" json:"video_plays"`
}
