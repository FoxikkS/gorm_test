package server

import (
	"gorm_test_proj/internal"
)

func Init() {
	internal.DBInit()
	routersInit()
}
