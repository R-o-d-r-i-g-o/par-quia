package model

import (
	"transaction/library"

	"gorm.io/gorm"
)

func Handler(db *gorm.DB) {

	populateTableOccupation(db)
	// populateStatusData(db)
}

func populateTableOccupation(db *gorm.DB) {
	occupations := []Occupation{
		{Id: 1, Name: "Leader of pastor(s)"},
		{Id: 2, Name: "experienced member"},
		{Id: 3, Name: "non-experienced member"},
	}

	for _, occupation := range occupations {
		if err := db.Table(library.TB_CATEGORIES).Find(&occupation).Error; err != nil {
			db.Table(library.TB_CATEGORIES).Create(occupation)
		}
	}
}
