package v1

import (
	//"log"

	"net/http"
	"oncall/src/hsc"
	"oncall/src/mysql"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	code := hsc.SUCCESS
	user := mysql.DB.GetAllUser()
	//log.Println( gettagtotal, count)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": user,
	})
}

func DeleteUser(c *gin.Context) {
	//id := com.StrTo(c.Param("id")).MustInt()
	id := c.Param("id")

	code := hsc.INVALID_PARAMS

	code = hsc.SUCCESS
	if mysql.DB.ExistUserID(id) {
		mysql.DB.DeleteUserId(id)
	} else {
		code = hsc.ERROR_NOT_EXIST
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": make(map[string]string),
	})
}

func AddUser(c *gin.Context) {
	name := c.Query("name")
	department := c.Query("department")
	state := c.DefaultQuery("state", "0")
	group := c.Query("group")

	code := hsc.INVALID_PARAMS

	if !mysql.DB.ExistUserByName(name, group) {
		code = hsc.SUCCESS
		mysql.DB.AddUser(name, department, state, group)
	} else {
		code = hsc.ERROR_EXIST
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": make(map[string]string),
	})
	c.JSON(http.StatusOK, gin.H{"ok": true})
}
