package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"pf-agent/internal/config"
	"pf-agent/internal/constant"
	"pf-agent/internal/logger"
	"time"
)

var (
	O *gorm.DB
)

func init() {
	var err error
	O, err = gorm.Open("mysql", config.C.GetString(constant.DataSource))
	if err != nil {
		logger.L.Fatalf("Fatal error db: %v\n", err)
	}
	O.SingularTable(true)
}

type Model struct {
	Id        int64      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}
