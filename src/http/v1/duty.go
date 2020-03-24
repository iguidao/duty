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
	c.JSON(http.StatusOK, gin.H{"ok": true})
}
