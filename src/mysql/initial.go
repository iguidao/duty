package mysql

import (
	"log"
	"oncall/src/model"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MySQL refrence a mysql db
type MySQL struct {
	*gorm.DB
}

// DB as the mysql client
var DB MySQL

// Connect create db connection
func Connect(dsn string) {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalln("Cannot open mysql database: ", err.Error())
		panic(err)
	}
	DB = MySQL{db}
	DB.SingularTable(true)
	DB.SetLogger(log.New(os.Stdout, "", 0))

}

// Migrate the db schema
func Migrate() {
	log.Println("start to auto migrate data schemas...")
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Duty{})
	DB.AutoMigrate(&model.Alarm{})
	//DB.Debug().AutoMigrate(&model.UserInfo{})
	log.Println("auto migrate data schemas done.")
}
