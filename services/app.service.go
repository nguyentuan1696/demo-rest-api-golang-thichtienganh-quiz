package services

import (
	"github.com/getsentry/sentry-go"
	"log"
	"quizbe/models"
	"quizbe/utils"
	"strings"
)

func CreateApp(param models.CreateAppParam) error {

	res := models.App{
		Name:        strings.ToLower(param.Name),
		Type:        strings.ToLower(param.Type),
		AppId:       utils.RandomString(16),
		AppSecret:   utils.RandomString(32),
		Description: param.Description,
		Extend:      param.Extend,
	}

	if err := models.CreateApp(res); err != nil {
		log.Printf("%v\n", err)
	}

	return nil
}

func GetApps() ([]models.App, error) {

	apps, err := models.GetApps()

	return apps, err
}

func GetApp(param string) (models.App, error) {

	app, err := models.GetApp(param)

	return app, err
}

func DeleleApp(param string) error {

	if err := models.DeleteApp(param); err != nil {
		log.Printf("%v\n", err)
		return err
	}

	return nil
}

func UpdateApp(appIdparam string, param map[string]interface{}) (models.App, error) {
	app, err := models.UpdateApp(appIdparam, param)
	if err != nil {
		log.Printf(" %v\n", err)
		sentry.CaptureException(err)
	}
	return app, nil
}
