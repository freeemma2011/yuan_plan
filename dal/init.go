package dal

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var SqlDB *gorm.DB

const (
	DBUser     = "root"
	DBPassport = "123455"
	Host       = "127.0.0.1"
	DB         = "yuan_plan"
	Port       = 3306
)

func InitDB() {
	var err error
	SqlDB, err = dbConn(DBUser, DBPassport, Host, DB, Port)
	if err != nil {
		panic(err.Error())
	}
}
