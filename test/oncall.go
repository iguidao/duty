	// 遇到节假日顺延，则没有周三开始值班一说
	// duty_exist := mysql.DB.DutyLastExist()
	// var nextday string
	// if !duty_exist {
	// 	currentTime := time.Now()
	// 	offset := int(time.Wednesday - currentTime.Weekday() + 7)
	// 	weekwendesday := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	// 	nextday = weekwendesday.Format("2006-01-02")
	// } else {
	// 	lastday := mysql.DB.DutyLastDay()
	// 	log.Println(lastday)
	// }
	// var duty_num int = 8
	// for i := 0; i < duty_num; i++ {
	// 	loc, _ := time.LoadLocation("Asia/Shanghai")
	// 	tt, _ := time.ParseInLocation("2006-01-02", nextday, loc)
	// 	date_day := time.Date(tt.Year(), tt.Month(), tt.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, i)
	// 	nextday := date_day.Format("2006-01-02")
	// 	if !ExistHoliday(nextday) {
	// 		log.Println(nextday)
	// 	} else {
	// 		duty_num = duty_num + 1
	// 	}
	// }