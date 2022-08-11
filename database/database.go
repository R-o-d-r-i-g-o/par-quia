package database

import (
	"fmt"
	"log"
	"scheduler/configs"
	"scheduler/database/model"

	"gorm.io/gorm"
)

var (
	db *gorm.DB

	tables = func(db *gorm.DB) error {
		return db.AutoMigrate(
			&model.User{},
			&model.Indisponibility{},
			&model.User_x_Pastor{},
			&model.Occupation{},
			&model.Pastor{},
			&model.PastorConfigs{},
			&model.MonthSchedule{},
			&model.CelebrationDay{},
			&model.DailyDuty{},
			&model.DaySchedule{},
			&model.Group{},
		)
	}
)

func GetGormDB() *gorm.DB {
	return db
}

func StartDatabase() {
	databaseConfiguration := createDatabaseStringConfig()
	mysql := NewMysql(databaseConfiguration)

	if err := mysql.connect(); err != nil {
		log.Fatal(err)
	}

	if tables(mysql.db) == nil {
		db = mysql.db
	}
}

func createDatabaseStringConfig() string {
	databaseStringConfig := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",

		configs.DBase.USER,
		configs.DBase.PASSWORD,
		configs.DBase.HOST,
		configs.DBase.PORT,
		configs.DBase.DB,
	)

	return databaseStringConfig
}
