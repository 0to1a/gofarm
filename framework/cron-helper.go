package framework

import (
	"github.com/go-co-op/gocron"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type CronUtils struct{}

func (w CronUtils) Setup() {
	if scheduler == nil {
		scheduler = gocron.NewScheduler(time.UTC)
	}
}

func (w CronUtils) AddEverySecond(timeInSecond int, jobFun interface{}, params ...interface{}) {
	if scheduler != nil {
		_, _ = scheduler.Every(timeInSecond).Seconds().Do(jobFun, params...)
	}
}

func (w CronUtils) AddEveryDay(atHourMinute string, jobFun interface{}, params ...interface{}) {
	if scheduler != nil {
		_, _ = scheduler.Every(1).Day().At(atHourMinute).Do(jobFun, params...)
	}
}

func (w CronUtils) Start() {
	if scheduler != nil {
		scheduler.StartAsync()
	}
}

func (w CronUtils) Stop() {
	if scheduler != nil {
		scheduler.Stop()
	}
}
