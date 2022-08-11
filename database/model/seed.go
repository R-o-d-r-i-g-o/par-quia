package model

import (
	"scheduler/constants"

	"gorm.io/gorm"
)

func Handler(db *gorm.DB) {

	populateTableOccupation(db)
}

func populateTableOccupation(db *gorm.DB) {
	occupations := []Occupation{
		{Id: 1, Name: "Leader of pastor(s)"},
		{Id: 2, Name: "experienced member"},
		{Id: 3, Name: "non-experienced member"},
	}

	for _, occupation := range occupations {
		if err := db.Table(constants.TB_OCCUPATION).Find(&occupation).Error; err != nil {
			db.Table(constants.TB_OCCUPATION).Create(occupation)
		}
	}
}
