package user

import (
	"github.com/jinzhu/gorm"
	"pf-agent/internal/db"
)

var (
	U *userService
)

func init() {
	U = new()
}

type userService struct {
	o *gorm.DB
}

func new() *userService {
	return &userService{o: db.O}
}
