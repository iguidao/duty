package v1

import (
	//"log"

	"net/http"

	"github.com/gin-gonic/gin"
	//"oncall/src/cfg"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func Cookie(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")

	if err != nil {
		cookie = "NotSet"
		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
	}

	c.JSON(http.StatusOK, gin.H{"cookie": cookie})
}
