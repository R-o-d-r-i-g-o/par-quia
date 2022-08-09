package model

import (
	"scheduler/constants"
	"time"
)

type DataBaseTables struct {
}

// ---------------------------------< 01º Table >--------------------------------- \\

type User struct {
	Id       uint64
	Name     string //`gorm:"index"`
	Email    string
	Phone    int
	PassWord string
	Indisp   []Indisponibility `gorm:"foreignKey:id_user;references:id"`
}

func (User) TableName() string {
	return constants.TB_USER
}

// ---------------------------------< 02º Table >--------------------------------- \\

type Indisponibility struct {
	YearMonthDay time.Time `json:"year_month_day  gorm:"column:year_month_day;primaryKey"`
	UserId       uint64    `json:"user_id" 		  gorm:"column:user_id;primaryKey"`
	NotPresent   bool      `json:"not_present" 	  gorm:"column:not_present"`
}

func (Indisponibility) TableName() string {
	return constants.TB_INDISPONIBILITY
}

// ---------------------------------< 03º Table >--------------------------------- \\

type User_x_Pastor struct {
	UserId       uint64 `json:"id_user" gorm:"primaryKey"`
	PastorId     uint64 `json:"id_pastor" gorm:"primaryKey"`
	OccupationId uint64 `json:"id_occupation" gorm:"primaryKey"`
}

func (User_x_Pastor) TableName() string {
	return constants.TB_USER_X_PASTOR
}

// ---------------------------------< 04º Table >--------------------------------- \\

type Occupation struct {
	Id     uint64
	Name   string
	Us_x_P User_x_Pastor
}

func (Occupation) TableName() string {
	return constants.TB_OCCUPATION
}

// ---------------------------------< 05º Table >--------------------------------- \\

type Pastor struct {
	Id           uint64
	Name         string
	Description  string
	CreatorId    int
	CreationDate int
}

func (Pastor) TableName() string {
	return constants.TB_PASTOR
}

// ---------------------------------< 06º Table >--------------------------------- \\

type PastorConfigs struct {
	PastorId             uint64
	OrderByExperience    bool
	OrderByGroup         bool
	ManualConfigInterval bool
}

func (PastorConfigs) TableName() string {
	return constants.TB_PASTOR_CONFIGS
}

// ---------------------------------< 07º Table >--------------------------------- \\

type MonthSchedule struct {
	YearMonth      time.Time
	PastorId       uint64
	IntervalInDays uint64
	Warning        string
}

func (MonthSchedule) TableName() string {
	return constants.TB_MONTH_SCHEDULE
}

// ---------------------------------< 08º Table >--------------------------------- \\

type CelebrationDay struct {
	YearMonthDay time.Time
	Hour         time.Time
	WorkDay      bool
	Solemnities  string
	LastEditorId uint64
}

func (CelebrationDay) TableName() string {
	return constants.TB_CELEBRATION_DAY
}

// ---------------------------------< 09º Table >--------------------------------- \\

type DailyDuty struct {
	YearMonthDay time.Time
	Hour         time.Time
	PastorId     uint64
	NumOfMembers uint64
}

func (DailyDuty) TableName() string {
	return constants.TB_DAILY_DUTY
}

// ---------------------------------< 10º Table >--------------------------------- \\

type DaySchedule struct {
	YearMonthDay time.Time
	Hour         time.Time
	PastorId     uint64
	UserId       uint64
	GroupId      uint64
}

func (DaySchedule) TableName() string {
	return constants.TB_DAY_SCHEDULE
}

// ---------------------------------< 11º Table >--------------------------------- \\

type DutyGroup struct {
	GroupId uint64
	UserId  uint64
}

func (DutyGroup) TableName() string {
	return constants.TB_DUTY_GROUP
}

// -----------------------------< General Interface >----------------------------- \\

type Tabler interface {
	TableName() string
}
