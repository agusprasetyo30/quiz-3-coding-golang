package model

import "time"

type User struct {
	ID         int        `json:"id" gorm:"primary_key"`
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	CreatedAt  time.Time  `json:"created_at"`
	CreatedBy  *string    `json:"created_by"`
	ModifiedAt *time.Time `json:"modified_at"`
	ModifiedBy *string    `json:"modified_by"`
}
