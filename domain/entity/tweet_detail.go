package entity

type TweetDetail struct {
	ID       uint     `gorm:"primary_key" json:"id"`
	UserID   uint     `json:"user_id"`
	Author   string   `json:"author"`
	Avatar   string   `json:"avatar"`
	Image    string   `json:"image"`
	Title    string   `json:"title"`
	Text     string   `json:"text"`
	PostDate string   `json:"post_date"`
	Comments Comments `json:"comments"`
}
