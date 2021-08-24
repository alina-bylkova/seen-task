package model

// Event is a database object
type Event struct {
	RecipientId string
	VideoId     string
	LpHits      int32
	VideoPlays  int32
}
