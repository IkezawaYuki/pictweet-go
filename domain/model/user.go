package model

import "time"

type User struct {
	ID        uint `gorm:"primary_key"`
	UID       string
	Name      string
	Email     string
	Avatar    string
	CreatedAt time.Time  `json:"-"`
	UpdateAt  time.Time  `json:"-"`
	DeleteAt  *time.Time `sql:"index" json:"-"`
}

type Users []User
