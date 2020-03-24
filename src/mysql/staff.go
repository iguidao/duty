package mysql

import (
	"oncall/src/model"
	"strconv"
)

func (m *MySQL) ExistUserID(id string) bool {
	id_int, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	var user model.User
	if m.Where("id = ?", id_int).First(&user).RecordNotFound() {

		return false
	}
	return true
}

func (m *MySQL) DeleteUserId(id string) bool {
	delstate := -1
	id_int, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	//m.Debug().Model(&model.User{}).Where("id = ?", id_int).Update("state_info", delstate)
	m.Model(&model.User{}).Where("id = ?", id_int).Update("state_info", delstate).Delete(&model.User{})
	return true
}

func (m *MySQL) Getuser(notstate int, group_name string) (user model.User) {
	m.Where("state_info = ? AND group_name = ?", notstate, group_name).First(&user)
	return user
}

func (m *MySQL) GetAllUser() []model.User {
	var user []model.User
	m.Find(&user)
	return user
}

func (m *MySQL) AddUser(name string, department string, state string, group string) bool {
	state_int, err := strconv.Atoi(state)
	if err != nil {
		panic(err)
	}
	m.Debug().Create(&model.User{
		StaffName:  name,
		StateInfo:  state_int,
		GroupName:  group,
		Department: department,
	})

	return true
}
