package http

import (
	//"net/http"
	v1 "oncall/src/http/v1"

	"github.com/gin-gonic/gin"
)

// NewServer return a configured http server of gin
func NewServer() *gin.Engine {
	// 存储日志文件代码
	// gin.DisableConsoleColor()
	// f, _ := os.Create("./logs/oncall.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()

	api := r.Group("/api/v1")

	{
		api.GET("/health", v1.HealthCheck)
		api.GET("/cookie", v1.Cookie)
	}
	duty := r.Group("/duty/v1")
	{
		duty.GET("/info", v1.DutyInfo)
		duty.PUT("/shift:id", v1.DutyShift)
	}
	staff := r.Group("/staff/v1")
	{
		staff.GET("/user", v1.GetUser)
		staff.POST("/user", v1.AddUser)
		staff.DELETE("/user/:id", v1.DeleteUser)
	}
	return r
}
