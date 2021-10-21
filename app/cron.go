package app

import (
	"framework/app/controller"
	"github.com/go-co-op/gocron"
	"log"
	"time"
)

func CronJobMaker() {
	schedule := gocron.NewScheduler(time.UTC)

	_, _ = schedule.Every(15).Seconds().Do(func() {
		log.Println("Hello per 15 seconds")
	})
	_, _ = schedule.Every(1).Minutes().Do(controller.CronExampleLog)

	schedule.StartAsync()
}
