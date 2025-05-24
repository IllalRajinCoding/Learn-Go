package models

import "time"

type User struct {
	Id uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
	Username string `json:"username" gorm:"not null;unique"`
	Email string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

}