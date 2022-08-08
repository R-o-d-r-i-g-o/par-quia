package database

import (
	"scheduler/constants"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Phone    int
	PassWord string
}

type Pastor struct {
	Id           int
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

func (Pastor) TableName() string {
	return constants.TB_PASTOR
}
