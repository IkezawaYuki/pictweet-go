package model

import "time"

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	UID       string     `json:"uid"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Avatar    string     `json:"avatar"`
	CreatedAt time.Time  `json:"-"`
	UpdateAt  time.Time  `json:"update_at"`
	DeleteAt  *time.Time `sql:"index" json:"-"`
}

type Users []User
