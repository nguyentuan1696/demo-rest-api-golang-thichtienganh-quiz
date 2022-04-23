package sentry

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"log"
	"quizbe/configs"
	"quizbe/utils"
)

func InitInstanceSentry() {
	connect := configs.Configs
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://" + connect.SecretKeySentry + "@o976125.ingest.sentry.io/" + connect.ProjectSentry,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	} else {
		fmt.Printf(utils.LogNoticeColor, "$$$ Connect Sentry "+connect.ProjectSentry+" Success $$$ \n")
	}
}
