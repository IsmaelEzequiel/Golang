package compaign

import "time"

type Contacts struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	CreatedAt time.Time
	Content   string
	Contacts  []Contacts
}

func NewCompaign(name string, content string) {
}
