package models

import (
	"github.com/getsentry/sentry-go"
	"log"
	"quizbe/configs/db"
	"time"
)

type App struct {
	ID          int       `gorm:"primaryKey;not null;unique;column:id" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Type        string    `gorm:"column:type"`
	AppId       string    `gorm:"column:app_id" json:"app_id"`
	AppSecret   string    `gorm:"column:app_secret" json:"app_secret"`
	Description string    `gorm:"column:description" json:"description"`
	Extend      string    `gorm:"column:extend" json:"extend"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func CreateApp(res App) error {

	if err := db.GetPostgresDB().Debug().Model(&App{}).Create(&res).Error; err != nil {
		log.Printf("%v\n", err)
		sentry.CaptureException(err)
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil

}

func GetApps() ([]App, error) {

	var apps []App

	result := db.GetPostgresDB().Debug().Model(&App{}).Find(&apps)

	if result.Error != nil {
		log.Printf(" %v\n", result.Error)
		sentry.CaptureException(result.Error)
		return apps, result.Error
	}

	return apps, result.Error
}

func GetApp(param string) (App, error) {
	var app App

	result := db.GetPostgresDB().Debug().Model(&App{}).Where("app_id = ? ", param).First(&app)

	if result.Error != nil {
		log.Printf(" %v\n", result.Error)
		sentry.CaptureException(result.Error)
	}

	return app, result.Error

}

func DeleteApp(param string) error {
	var app App

	// Checking, if app with given ID is exists.
	err := db.GetPostgresDB().Debug().Model(&App{}).Where("app_id = ?", param).First(&app).Error
	if err != nil {
		return err
	}

	if err := db.GetPostgresDB().Debug().Model(&App{}).Where("app_id = ?", param).Delete(&app).Error; err != nil {
		log.Printf(" %v\n", err)
		sentry.CaptureException(err)
		return err
	}

	return nil
}

func UpdateApp(appIdparam string, param map[string]interface{}) (App, error) {

	var app App

	if err := db.GetPostgresDB().Debug().Model(&App{}).Where("app_id = ?", appIdparam).Updates(param).First(&app).Error; err != nil {
		log.Printf(" %v\n", err)
		sentry.CaptureException(err)
		return app, err
	}

	return app, nil
}
