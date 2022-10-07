package framework

import (
	"github.com/bmizerany/assert"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

func TestSetup(t *testing.T) {
	t.Run("Create instance", func(t *testing.T) {
		cron.Setup()
		assert.NotEqual(t, nil, cron.scheduler)
		assert.NotEqual(t, nil, scheduler)
	})
}

func TestAddEverySecond(t *testing.T) {
	t.Run("Add on global", func(t *testing.T) {
		cron.Setup()
		cron.scheduler = nil
		cron.AddEverySecond(10, func() { log.Println("hello") })
		assert.NotEqual(t, nil, scheduler)
		assert.Equal(t, 1, scheduler.Len())
	})
	t.Run("Add on local", func(t *testing.T) {
		cron.Setup()
		cron.AddEverySecond(10, func() { log.Println("hello") })
		assert.NotEqual(t, nil, cron.scheduler)
		assert.Equal(t, 1, cron.scheduler.Len())
	})
}

func TestAddEveryDay(t *testing.T) {
	t.Run("Add on global", func(t *testing.T) {
		cron.Setup()
		cron.scheduler = nil
		cron.AddEveryDay("1:00", func() { log.Println("hello") })
		assert.NotEqual(t, nil, scheduler)
		assert.Equal(t, 1, scheduler.Len())
	})
	t.Run("Add on local", func(t *testing.T) {
		cron.Setup()
		cron.AddEveryDay("1:00", func() { log.Println("hello") })
		assert.NotEqual(t, nil, cron.scheduler)
		assert.Equal(t, 1, cron.scheduler.Len())
	})
}

func TestStart(t *testing.T) {
	t.Run("Start on global", func(t *testing.T) {
		cron.Setup()
		cron.scheduler = nil
		cron.Start()
		assert.NotEqual(t, nil, scheduler)

	})
	t.Run("Start on local", func(t *testing.T) {
		cron.Setup()
		cron.Start()
		assert.NotEqual(t, nil, cron.scheduler)
	})
}

func TestStop(t *testing.T) {
	t.Run("Start on global", func(t *testing.T) {
		cron.Setup()
		cron.scheduler = nil
		cron.Stop()
		assert.NotEqual(t, nil, scheduler)
	})
	t.Run("Start on local", func(t *testing.T) {
		cron.Setup()
		cron.Stop()
		assert.NotEqual(t, nil, cron.scheduler)
	})
}
