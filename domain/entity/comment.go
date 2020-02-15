package entity

type Comment struct {
	ID      int
	UserID  string
	TweetID string
	Text    string
}

type Comments []Comment
