package port

import (
	"time"
)

type OutputPort interface {
	Index()
}

type TweetsDto struct {
	ID       uint      `gorm:"primary_key" json:"id"`
	Author   string    `json:"author"`
	Image    string    `json:"image"`
	Title    string    `json:"title"`
	Text     string    `json:"text"`
	UpdateAt time.Time `json:"update_at"`
}
