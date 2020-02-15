package entity

type Tweet struct {
	ID     int
	UserID string
	Image  string
	Title  string
	Text   string
}

type Tweets []Tweet
