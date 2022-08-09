package model

import (
	"scheduler/constants"
	"time"
)

type User struct {
	Id       uint64
	Name     string //`gorm:"index"`
	Email    string
	Phone    int
	PassWord string
	Indisp   []Indisponibility `gorm:"foreignKey:id_user;references:id"`
}

type Indisponibility struct {
	Id           uint64
	YearMonthDay time.Time `json:"year_month_day" gorm:"primaryKey"`
}

type User_x_Pastor struct {
	UserId       uint64 `json:"id_user" gorm:"primaryKey"`
	PastorId     uint64 `json:"id_pastor" gorm:"primaryKey"`
	OccupationId uint64 `json:"id_occupation" gorm:"primaryKey"`
}

type Occupation struct {
	Id     uint64
	Name   string
	Us_x_P User_x_Pastor
}

type Pastor struct {
	Id           uint64
	Name         string
	Description  string
	CreatorId    int
	CreationDate int
}

type PastorConfigs struct {
	PastorId             uint64
	OrderByExperience    bool
	OrderByGroup         bool
	ManualConfigInterval bool
}

type MonthSchedule struct {
	YearMonth      time.Time
	PastorId       uint64
	IntervalInDays uint64
	Warning        string
}

type CelebrationDay struct {
	YearMonthDay time.Time
	Hour         time.Time
	WorkDay      bool
	Solemnities  string
	LastEditorId uint64
}

type dailyDuty struct {
	YearMonthDay time.Time
	Hour         time.Time
	PastorId     uint64
	NumOfMembers uint64
}

type DaySchedule struct {
	YearMonthDay time.Time
	Hour         time.Time
	PastorId     uint64
	UserId       uint64
	GroupId      uint64
}

type DutyGroup struct {
	GroupId uint64
	UserId  uint64
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
