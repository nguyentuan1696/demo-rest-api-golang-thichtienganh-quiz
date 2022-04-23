package db

import (
	"gorm.io/gorm/schema"
	"log"
	"net/url"
	"quizbe/configs"
	"quizbe/utils"
	"time"

	"fmt"
	"github.com/getsentry/sentry-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var PostgresDB *gorm.DB

func ConnectPostgres() *gorm.DB {
	connect := configs.Configs.Sql
	password := url.QueryEscape(connect.Password)
	dsn := "postgres://" + connect.Username + ":" + password + "@" + connect.Host + ":" + connect.Port + "/" + connect.Database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{ // bỏ thêm 's' vào tên bảng
			SingularTable: true,
		},
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Printf("%v\n", err)
		sentry.CaptureException(err)
	} else {
		fmt.Printf(utils.LogNoticeColor, "$$$ Connect DB "+connect.Name+" "+connect.Host+" Success $$$ \n")
	}

	postgresDB, _ := db.DB()
	postgresDB.SetMaxIdleConns(10)
	postgresDB.SetMaxOpenConns(100)
	postgresDB.SetConnMaxLifetime(24 * time.Hour)

	PostgresDB = db
	return PostgresDB
}

func GetPostgresDB() *gorm.DB {
	sqlDB, _ := PostgresDB.DB()
	err := sqlDB.Ping()
	if err != nil {
		fmt.Println("Connection to PostgresSQL closed => opening a new one")
		return ConnectPostgres()
	}
	return PostgresDB
}
