package database

import (
	"scheduler/constants"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string //`gorm:"index"`
	Email    string
	Phone    int
	PassWord string
	Indisp   []Indisponibility `gorm:"foreignKey:user_id;references:id"`
}

type Indisponibility struct {
	yearMonthDay time.Time `gorm:"primaryKey"`
	user_id      int       `gorm:"primaryKey"`
}

type Pastor struct {
	Id           int `gorm:"primaryKey"`
	Name         string
	Description  string
	IdCreator    int
	CreationDate int
}

type Tabler interface {
	TableName() string
}

// ---------------------------------< Methods to name database tables >--------------------------------- \\

func (User) TableName() string {
	return constants.TB_USER
}

func (Indisponibility) TableName() string {
	return constants.TB_INDISPONIBILITY
}

func (Pastor) TableName() string {
	return constants.TB_PASTOR
}
