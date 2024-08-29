package account

import (
	"pf-agent/internal/db"
	"time"
)

type Account struct {
	db.Model
	Mobile string    `json:"mobile" gorm:"column:mobile" binding:"required"`
	Expire time.Time `json:"expire" gorm:"column:expire"`
	Token  string    `json:"token" gorm:"column:token"`
}
