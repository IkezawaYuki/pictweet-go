package entity

type Comment struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	UserID   uint   `json:"user_id"`
	Author   string `json:"author"`
	Avatar   string `json:"avatar"`
	Text     string `json:"text"`
	PostDate string `json:"post_date"`
}

type Comments []Comment
