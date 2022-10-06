package framework

import (
	"github.com/go-co-op/gocron"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type CronUtils struct {
	scheduler *gocron.Scheduler
}

func (w *CronUtils) Setup() {
	if w.scheduler == nil {
		w.scheduler = gocron.NewScheduler(time.UTC)
		scheduler = gocron.NewScheduler(time.UTC)
	}
}

func (w *CronUtils) AddEverySecond(timeInSecond int, jobFun interface{}, params ...interface{}) {
	if w.scheduler != nil {
		_, _ = w.scheduler.Every(timeInSecond).Seconds().Do(jobFun, params...)
	} else if scheduler != nil {
		_, _ = scheduler.Every(timeInSecond).Seconds().Do(jobFun, params...)
	}
}

func (w *CronUtils) AddEveryDay(atHourMinute string, jobFun interface{}, params ...interface{}) {
	if w.scheduler != nil {
		_, _ = w.scheduler.Every(1).Day().At(atHourMinute).Do(jobFun, params...)
	} else if scheduler != nil {
		_, _ = scheduler.Every(1).Day().At(atHourMinute).Do(jobFun, params...)
	}
}

func (w *CronUtils) Start() {
	if w.scheduler != nil {
		w.scheduler.StartAsync()
	} else if scheduler != nil {
		scheduler.StartAsync()
	}
}

func (w *CronUtils) Stop() {
	if w.scheduler != nil {
		w.scheduler.Stop()
	} else if scheduler != nil {
		scheduler.Stop()
	}
}
