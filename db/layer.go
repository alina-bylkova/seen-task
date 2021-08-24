package db

type Layer interface {
	Add(Recipient)
	GetAll() []Recipient
}

type Recipient struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone uint32 `json:"phone"`
}
