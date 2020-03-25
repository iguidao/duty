package mysql

import (
	"log"
	"oncall/src/model"
	"strconv"
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
func (m *MySQL) ExistDutyByID(id string) bool {
	idint, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	var duty model.Duty
	if m.Where("id = ?", idint).First(&duty).RecordNotFound() {
		log.Println("buzai")
		return false
	}
	log.Println("zai")
	return true

}
func (m *MySQL) EditDuty(id, userid string) bool {
	idint, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	useridint, err := strconv.Atoi(userid)
	if err != nil {
		panic(err)
	}
	var duty model.Duty
	m.Where("id = ?", idint).First(&duty).Update("user_id", useridint)
	return true

}
