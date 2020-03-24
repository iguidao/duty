package oncron

import (
	"oncall/src/cfg"
	"oncall/src/model"
	"oncall/src/mysql"
	"time"
)

func ExistHoliday(nextday string) bool {
	holiday := cfg.Get_Slice("holiday")
	for _, aaa := range holiday {
		if nextday == aaa {
			return true
		}
	}
	return false
}

func DutyResult(dutystate, notdutystate int, group_name string) (user model.User) {

	if !mysql.DB.UserOnDuty(notdutystate, group_name) {
		mysql.DB.DutyCheckState(dutystate, notdutystate, group_name)
	}
	staff_user := mysql.DB.Getuser(notdutystate, group_name)
	return staff_user
}

func DutyDate() []string {

	var nextday string

	currentTime := time.Now()
	offset := int(time.Wednesday - currentTime.Weekday() + 7)
	weekwendesday := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	nextday = weekwendesday.Format("2006-01-02")

	var dutydays []string
	var duty_num int = 8
	for i := 0; i < duty_num; i++ {
		loc, _ := time.LoadLocation("Asia/Shanghai")
		tt, _ := time.ParseInLocation("2006-01-02", nextday, loc)
		date_day := time.Date(tt.Year(), tt.Month(), tt.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, i)
		nextday := date_day.Format("2006-01-02")
		if !ExistHoliday(nextday) {
			dutydays = append(dutydays, nextday)
		}
	}

	return dutydays
}

func DutyExec(staff_id int, duty_day, interval, group_name string) bool {
	duty_exist := mysql.DB.DutyExist(duty_day, interval, group_name)
	if !duty_exist {
		mysql.DB.Scheduling(staff_id, duty_day, interval, group_name)
		return true
	} else {
		return false
	}

}
func DutyStart(duty_date []string, staff_id int, group_name string) bool {
	exec_num := 0
	if len(duty_date) != 0 {
		duty_len := len(duty_date) - 1
		for i, duty_day := range duty_date {
			if i == 0 {
				interval := "PM"
				duty_result := DutyExec(staff_id, duty_day, interval, group_name)
				if !duty_result {
					exec_num = exec_num + 1
				}
			} else if i == duty_len {
				interval := "AM"
				duty_result := DutyExec(staff_id, duty_day, interval, group_name)
				if !duty_result {
					exec_num = exec_num + 1
				}
			} else {
				interval := "AM"
				duty_am_result := DutyExec(staff_id, duty_day, interval, group_name)
				interval = "PM"
				duty_pm_result := DutyExec(staff_id, duty_day, interval, group_name)
				if !duty_am_result && !duty_pm_result {
					exec_num = exec_num + 1
				}
			}
		}
		if exec_num == len(duty_date) {
			return false
		} else {
			return true
		}
	} else {
		return false
	}

}
func CronDuty() bool {
	groups := cfg.Get_Slice("group")
	notdutystate := 1
	dutystate := 0
	for _, group_name := range groups {
		staff_id := int(DutyResult(dutystate, notdutystate, group_name).ID)
		duty_date := DutyDate()
		duty_result := DutyStart(duty_date, staff_id, group_name)
		if duty_result {
			mysql.DB.DutyChengeState(staff_id, dutystate, group_name)
		}

	}
	return true
}
