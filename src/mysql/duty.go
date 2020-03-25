package mysql

import (
	"log"
	"oncall/src/model"
	"time"
)

func (m *MySQL) GetAllDuty(starttime, endtime string) []model.Duty {
	loc, _ := time.LoadLocation("Local")
	startday, err := time.ParseInLocation("2006-01-02", starttime, loc)
	if err != nil {
		log.Println("转换时间失败")
	}
	endday, err := time.ParseInLocation("2006-01-02", endtime, loc)
	if err != nil {
		log.Println("转换时间失败")
	}
	var duty []model.Duty

	m.Preload("User").Where("date_time BETWEEN (?) AND (?)", startday, endday).Find(&duty)
	//m.Preload("User").Where("date_time > (?) AND date_time < (?)", startday, endday).Find(&duty)
	return duty
}
