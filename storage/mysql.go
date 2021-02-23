package storage

import (
	"fmt"
	"log"
	"time"
	"webapp/apptoml"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var GDb *gorm.DB
var err error

//
func InitDB(serverAddr string, user string, pwd string, database string, maxOpen int, maxIdle int, idleTime int) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, serverAddr, database)
	GDb, err = gorm.Open("mysql", connStr)
	if err != nil {
		log.Printf("init mysql error: %+v", err)
		panic(err)
	}

	GDb.DB().SetMaxOpenConns(maxOpen)
	GDb.DB().SetMaxIdleConns(maxIdle)
	GDb.DB().SetConnMaxLifetime(time.Duration(idleTime) * time.Second)

	if apptoml.Config.Server.Debug {
		GDb.LogMode(true)
	}

	log.Printf("maxOpen:%+v, maxIdle:%+v, idleTime:%+v, init mysql ok.", maxOpen, maxIdle, idleTime)
}
