package dto

import "time"

type UserDto struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	UID       string     `json:"uid"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Avatar    string     `json:"avatar"`
	CreatedAt time.Time  `json:"-"`
	UpdateAt  time.Time  `json:"updated_at"`
	DeleteAt  *time.Time `sql:"index" json:"-"`
}

type UsersDto []UserDto

func (UserDto) TableName() string {
	return "users"
}
