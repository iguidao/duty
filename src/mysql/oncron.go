package mysql

import (
	"log"
	"oncall/src/model"
	"time"
)

// func (m *MySQL) DutyLastDay() (duty model.Duty) {
// 	m.Last(&duty)
// 	return
// }

// func (m *MySQL) DutyLastExist() bool {
// 	var duty model.Duty
// 	if m.Last(&duty).RecordNotFound() {
// 		return false
// 	}
// 	return true

// }

func (m *MySQL) UserOnDuty(notstate int, group_name string) bool {
	var user model.User
	if m.Where("state_info = ? AND group_name = ? ", notstate, group_name).First(&user).RecordNotFound() {
		return false
	}
	return true

}

func (m *MySQL) DutyExist(duty_day, interval, group_name string) bool {
	loc, _ := time.LoadLocation("Local")
	the_day, err := time.ParseInLocation("2006-01-02", duty_day, loc)
	if err == nil {
		var duty model.Duty
		if m.Where("date_time = ? AND interval_time = ? AND group_name = ?", the_day, interval, group_name).First(&duty).RecordNotFound() {
			return false
		}
		return true
	} else {
		return false
	}
}

func (m *MySQL) DutyCheckState(state, notstate int, group_name string) bool {
	var user model.User
	m.Model(&user).Where("state_info = ? AND group_name = ? ", state, group_name).Update("state_info", notstate)
	return true

}

func (m *MySQL) DutyChengeState(name_id int, state int, group_name string) bool {
	var user model.User
	m.Model(&user).Where("id = ? AND group_name = ?", name_id, group_name).Update("state_info", state)
	return true

}

func (m *MySQL) ExistUserByName(name, group string) bool {
	var user model.User
	if m.Where("staff_name = ? AND group_name = ?", name, group).First(&user).RecordNotFound() {
		log.Println("meiyou")
		return false
	}
	return true
}

func (m *MySQL) Scheduling(staff_id int, duty_day string, interval string, group_name string) bool {
	loc, _ := time.LoadLocation("Local")
	the_day, err := time.ParseInLocation("2006-01-02", duty_day, loc)
	if err == nil {
		m.Create(&model.Duty{
			UserID:       staff_id,
			DateTime:     the_day,
			IntervalTime: interval,
			GroupName:    group_name,
		})
		return true
	} else {
		return false
	}

}
