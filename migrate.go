package main

import (
	"flag"
	"oncall/src/cfg"
	"oncall/src/mysql"
)

var migrate = flag.Bool("m", false, "migrate the database schemas.")

func init() {
	if err := cfg.Init(""); err != nil {
		panic(err)
	}
	mysql.Connect(cfg.Get_Info("MYSQL"))
}
func main() {
	flag.Parse()
	if *migrate {
		mysql.Migrate()
		return
	}
}
