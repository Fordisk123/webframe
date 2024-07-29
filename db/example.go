package db

import "context"

type Example struct {
	ID     int64  `gorm:"primary_key"`
	Name   string `gorm:"column:name"`
	Type   string `gorm:"column:type"` // http, tcp, icmp etc...
	Switch *bool  `gorm:"column:switch"`
}

var models = []interface{}{
	Example{},
}

func MigrateTable() {
	ctx := context.Background()
	if err := GetDb(ctx).AutoMigrate(models...); err != nil {
		panic(err.Error())
	}
}
