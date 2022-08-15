package model

import (
	"scheduler/library"
	"time"

	"gorm.io/gorm"
)

type DataBaseTables struct {
	gorm.Model
}

// ---------------------------------< 01º Table >--------------------------------- \\

type User struct {
	Id       uint64            `json:"id"        gorm:"column:id;           primaryKey"`
	Name     string            `json:"name"      gorm:"column:name"`
	Email    string            `json:"email"     gorm:"column:email"`
	Phone    int               `json:"phone"     gorm:"column:phone"`
	Password string            `json:"password"  gorm:"column:password"`
	Indisp   []Indisponibility `                 gorm:"foreignKey:user_id;  references:id"`
	Us_x_P   []User_x_Pastor   `                 gorm:"foreignKey:user_id;  references:id"`
}

func (User) TableName() string {
	return library.TB_USER
}

// ---------------------------------< 02º Table >--------------------------------- \\

type Indisponibility struct {
	YearMonthDay time.Time `json:"year_month_day"  gorm:"column:year_month_day;     primaryKey;    autoCreateTime:true"`
	UserId       uint64    `json:"user_id"         gorm:"column:user_id;            primaryKey"`
	NotPresent   bool      `json:"not_present"     gorm:"column:not_present"`
}

func (Indisponibility) TableName() string {
	return library.TB_INDISPONIBILITY
}

// ---------------------------------< 03º Table >--------------------------------- \\

type User_x_Pastor struct {
	UserId       uint64 `json:"id_user"        gorm:"column:user_id;            primaryKey"`
	PastorId     uint64 `json:"id_pastor"      gorm:"column:pastor_id;          primaryKey"`
	OccupationId uint64 `json:"id_occupation"  gorm:"column:occupation_id"`
}

func (User_x_Pastor) TableName() string {
	return library.TB_USER_X_PASTOR
}

// ---------------------------------< 04º Table >--------------------------------- \\

type Occupation struct {
	Id     uint64        `json:"id"    gorm:"column:id;                 primaryKey"`
	Name   string        `json:"name"  gorm:"column:name"`
	Us_x_P User_x_Pastor `             gorm:"foreignKey:occupation_id;  references:id"`
}

func (Occupation) TableName() string {
	return library.TB_OCCUPATION
}

// ---------------------------------< 05º Table >--------------------------------- \\

type Pastor struct {
	Id           uint64          `json:"id"             gorm:"column:id;             primaryKey"`
	Name         string          `json:"name"           gorm:"column:name"`
	Description  string          `json:"description"    gorm:"column:description"`
	CreatorId    int             `json:"creator_id"     gorm:"column:creator_id"`
	CreationDate int             `json:"creation_date"  gorm:"column:creation_date"`
	Us_x_P       []User_x_Pastor `                      gorm:"foreignKey:pastor_id;  references:id"`
	Month_sch    []MonthSchedule `                      gorm:"foreignKey:pastor_id;  references:id"`
	P_configs    PastorConfigs   `                      gorm:"foreignKey:pastor_id"`
}

func (Pastor) TableName() string {
	return library.TB_PASTOR
}

// ---------------------------------< 06º Table >--------------------------------- \\

type PastorConfigs struct {
	PastorId             uint64 `json:"pastor_id"               gorm:"column:pastor_id;               primaryKey"`
	OrderByExperience    bool   `json:"order_by_experience"     gorm:"column:order_by_experience"`
	OrderByGroup         bool   `json:"order_by_group"          gorm:"column:order_by_group"`
	ManualConfigInterval bool   `json:"manual_config_interval"  gorm:"column:manual_config_interval"`
}

func (PastorConfigs) TableName() string {
	return library.TB_PASTOR_CONFIGS
}

// ---------------------------------< 07º Table >--------------------------------- \\

type MonthSchedule struct {
	YearMonth      time.Time   `json:"year_month"        gorm:"column:year_month;         primaryKey;      autoCreateTime:true"`
	PastorId       uint64      `json:"pastor_id"         gorm:"column:pastor_id;          primaryKey"`
	IntervalInDays uint64      `json:"interval_in_days"  gorm:"column:interval_in_days"`
	Warning        string      `json:"warning"           gorm:"column:warning"`
	DailyDusty     []DailyDuty `                         gorm:"foreignKey:pastor_id       references:id"`
}

func (MonthSchedule) TableName() string {
	return library.TB_MONTH_SCHEDULE
}

// ---------------------------------< 08º Table >--------------------------------- \\

type CelebrationDay struct {
	YearMonthDay time.Time `json:"year_month_day"  gorm:"column:year_month_day;                      primaryKey;    autoCreateTime:true"`
	Hour         time.Time `json:"hour"            gorm:"column:hour;                                primaryKey;    autoCreateTime:false"`
	WorkDay      bool      `json:"work_day"        gorm:"column:work_day"`
	Solemnities  string    `json:"solemnities"     gorm:"column:solemnities"`
	LastEditorId uint64    `json:"last_editor_id"  gorm:"column:last_editor_id"`
	SyncDutyDay  DailyDuty `                       gorm:"foreignKey:year_month_day; references:year_month_day; foreignKey:hour; references:hour"`
}

func (CelebrationDay) TableName() string {
	return library.TB_CELEBRATION_DAY
}

// ---------------------------------< 09º Table >--------------------------------- \\

type DailyDuty struct {
	YearMonthDay   time.Time   `json:"year_month_day"  gorm:"column:year_month_day;                                       primaryKey;    autoCreateTime:true"`
	Hour           time.Time   `json:"hour"            gorm:"column:hour;                                                 primaryKey;    autoCreateTime:false"`
	PastorId       uint64      `json:"pastor_id"       gorm:"column:pastor_id;                                            primaryKey"`
	NumOfMembers   uint64      `json:"num_of_members"  gorm:"column:num_of_members"`
	NumOfGroups    uint64      `json:"num_of_groups"   gorm:"column:num_of_groups"`
	SyncDaySchDep1 DaySchedule `                       gorm:"foreignKey:year_month_day; references:year_month_day; foreignKey:hour; references:hour; foreignKey:pastor_id; references:pastor_id"`
}

func (DailyDuty) TableName() string {
	return library.TB_DAILY_DUTY
}

// ---------------------------------< 10º Table >--------------------------------- \\

type DaySchedule struct {
	YearMonthDay time.Time `json:"year_month_day"  gorm:"column:year_month_day;  primaryKey;    autoCreateTime:true"`
	Hour         time.Time `json:"hour"            gorm:"column:hour;            primaryKey;    autoCreateTime:false"`
	PastorId     uint64    `json:"pastor_id"       gorm:"column:pastor_id;       primaryKey"`
	UserId       uint64    `json:"user_id"         gorm:"column:user_id"`
	GroupId      uint64    `json:"group_id"        gorm:"column:group_id"`
	MembGroup    []Group   `                       gorm:"foreignKey:id;          references:group_id"`
}

func (DaySchedule) TableName() string {
	return library.TB_DAY_SCHEDULE
}

// ---------------------------------< 11º Table >--------------------------------- \\

type Group struct {
	Id     uint64 `json:"id"        gorm:"column:id;        primaryKey"`
	UserId uint64 `json:"user_id"   gorm:"column:user_id"`
}

func (Group) TableName() string {
	return library.TB_GROUP
}

// -----------------------------< General Interface >----------------------------- \\

type Tabler interface {
	TableName() string
}
