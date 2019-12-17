package conf

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func MysqlConnection(cfg Config) (err error, db *gorm.DB) {
	dbUser := cfg.GetString(`mysql.user`)
	dbPass := cfg.GetString(`mysql.pass`)
	dbName := cfg.GetString(`mysql.database`)
	dbHost := cfg.GetString(`mysql.address`)
	dbPort := cfg.GetString(`mysql.port`)

	db, err = gorm.Open("mysql", ""+dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err, nil
	}

	return
}
