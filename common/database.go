package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {

	//dsn: "root:root@tcp(127.0.0.1:3306)/shipofhope?charset=utf8&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		viper.Get("dataSource.username"),
		viper.Get("dataSource.password"),
		viper.Get("dataSource.host"),
		viper.Get("dataSource.port"),
		viper.Get("dataSource.database"),
		viper.Get("dataSource.charset"),
		viper.Get("dataSource.loc"),
	)
	log.Println("dsn: ", dsn)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,  // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		log.Println("database connect fail err:", err.Error())
	}
	DB = db
}

func GetDB() *gorm.DB {
	DB.Debug()
	return DB
}
