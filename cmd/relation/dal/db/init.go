package db

import (
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

func Init() {
	DB, err := gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&Relation{})
	if err != nil {
		return
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

}
