package model

import (
	model "transaction/db/model"
	"transaction/library"

	"gorm.io/gorm"
)

func Handler(db *gorm.DB) {

	populateTableOccupation(db)
	populateStatusData(db)
}

func populateTableOccupation(db *gorm.DB) {
	occupations := []Occupation{
		{Id: 1, Name: "Leader of pastor(s)", Users: []model.User{}},
		{Id: 2, Name: "experienced member", Users: []model.User{}},
		{Id: 3, Name: "non-experienced member", Users: []model.User{}},
	}

	for _, occupation := range occupations {
		if err := db.Table(library.TB_CATEGORIES).Find(&occupation).Error; err != nil {
			db.Table(library.TB_CATEGORIES).Create(occupation)
		}
	}
}

func populateStatusData(db *gorm.DB) {
	statusArr := []model.Status{
		{Id: 1, Name: "Autorizado", Transaction: []model.Transaction{}},
		{Id: 2, Name: "Não Autorizado", Transaction: []model.Transaction{}},
	}

	for _, status := range statusArr {
		if err := db.Table(library.TB_STATUS).Find(&status).Error; err != nil {
			db.Table(library.TB_STATUS).Create(&status)
		}
	}
}
