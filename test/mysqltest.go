package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
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

	DateTime     string `json:"date"`
	IntervalTime string `json:"interval"`
	GroupName    string `json:"group"`
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

type Result struct {
	StaffName  string
	Department string
	GroupName  string
	StateInfo  int
}

func main() {
	db, err := gorm.Open("mysql", "root:11qqAAAA@(localhost:3306)/oncall?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.SingularTable(true)

	//var result Result
	user := []User{}
	//db.Find(&user)
	db.Where("state_info = ?" , -1).First(&user)
	//fmt.Println(user)
	//db.Debug().Create(user)
	// var usif UserInfo
	// db.Take(&usif)
	// fmt.Println(usif)
	//db.Take(&user).Scan(&result)
	fmt.Println(user)

	// test_b := result.ID
	// fmt.Println("===============")
	// fmt.Println(test_b)
	// fmt.Println("===============")
	//    test_uuid, err := uuid.FromBytes(test_b)
	//    if err != nil {
	//            log.Println("uuid get error: %s", err)
	//    }
	//    //test_string := result.UserPassword
	//    fmt.Println(test_uuid)
}
