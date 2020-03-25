package v1

import (
	//"log"

	"net/http"
	"oncall/src/hsc"
	"oncall/src/mysql"

	"github.com/gin-gonic/gin"
	//"oncall/src/cfg"
)

func DutyInfo(c *gin.Context) {
	starttime := c.Query("starttime")
	endtime := c.Query("endtime")
	code := hsc.SUCCESS
	duty := mysql.DB.GetAllDuty(starttime, endtime)
	//log.Println( gettagtotal, count)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": duty,
	})
}


func DutyShift(c *gin.Context) {
	id := c.Param("id")
	userid := c.Query("userid")

	code := hsc.INVALID_PARAMS
	if mysql.DB.ExistDutyByID(id) {
		mysql.DB.EditDuty(id, userid)
		code = hsc.SUCCESS
	} else {
		code = hsc.ERROR_NOT_EXIST
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": make(map[string]string),
	})
}

