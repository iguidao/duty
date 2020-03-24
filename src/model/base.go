package model

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Base struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
	DeletedAt *time.Time
	//创建索引`sql:"index"`
}

type User struct {
	Base
	StaffName  string `gorm:"size:255"`
	Department string `gorm:"size:255"`
	GroupName  string `gorm:"size:255"`
	StateInfo  int    `gorm:"size:255"`
}

type Duty struct {
	Base
	UserID int  `json:"user_id" gorm:"index"`
	User   User `json:"user"`

	DateTime     time.Time `json:"date"`
	IntervalTime string    `json:"interval"`
	GroupName    string    `json:"group"`
}

type Alarm struct {
	Base
	UserID int  `json:"user_id" gorm:"index"`
	User   User `json:"user"`

	AlarmName   string `json:"name"`
	AlarmMetric string `json:"metric"`
	AlarmState  int    `json:"state"`
	AlarmResult string `json:"result"`
	ProductName string `json:"product"`
	AlarmNum    int    `json:"num"`
	AlarmTime   string `json:"time"`
	DateTime    string `json:"date"`
}
