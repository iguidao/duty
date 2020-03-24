package main

import (
	"oncall/src/cfg"
	"oncall/src/http"
	"oncall/src/middleware/oncron"
	"oncall/src/mysql"

	"github.com/robfig/cron"
)

func init() {
	if err := cfg.Init(""); err != nil {
		panic(err)
	}
	mysql.Connect(cfg.Get_Info("MYSQL"))
}

func main() {
	c := cron.New()
	crontime := cfg.Get_Local("cronduty")
	c.AddFunc(crontime, func() {
		oncron.CronDuty()
	})
	c.Start()

	listen := cfg.Get_Local("addr")
	if listen == "" {
		listen = ":8080"
	}

	http.NewServer().Run(listen)
}
