package framework

import (
	"github.com/alicebob/miniredis"
	"github.com/bmizerany/assert"
	"github.com/go-redis/redis/v8"
	"log"
	"testing"
	"time"
)

const (
	miniRedisError = "an error '%s' was not expected when opening a stub database connection"
)

func TestConnectRedis(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf(miniRedisError, err)
	}

	t.Run("Without connection", func(t *testing.T) {
		redisConf := RedisDatabase{Host: mr.Addr(), Password: "hello", Database: 10}
		assert.Equal(t, (*redis.Client)(nil), redisConf.Connect())
	})
	t.Run("With connection", func(t *testing.T) {
		redisConf := RedisDatabase{Host: mr.Addr()}
		assert.NotEqual(t, (*redis.Client)(nil), redisConf.Connect())
	})
}

func TestCheckClientRedis(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf(miniRedisError, err)
	}
	redisMock := RedisDatabase{Host: mr.Addr()}

	t.Run("Without connection", func(t *testing.T) {
		redisMock.client = nil
		RedisDB = nil
		assert.Equal(t, (*redis.Client)(nil), redisMock.CheckClient())
	})
	t.Run("With global connection", func(t *testing.T) {
		redisMock.Connect()
		redisMock.client = nil
		assert.NotEqual(t, (*redis.Client)(nil), redisMock.CheckClient())
	})
	t.Run("With local connection", func(t *testing.T) {
		redisMock.Connect()
		RedisDB = nil
		assert.NotEqual(t, (*redis.Client)(nil), redisMock.CheckClient())
	})
}

func TestCheckPrefix(t *testing.T) {
	redisMock := RedisDatabase{}

	t.Run("Not connect", func(t *testing.T) {
		redisMock.Prefix = ""
		RedisPrefix = ""
		assert.Equal(t, "", redisMock.CheckPrefix())
	})
	t.Run("With global connection", func(t *testing.T) {
		target := "hello-world"
		redisMock.Prefix = ""
		RedisPrefix = target
		assert.Equal(t, target, redisMock.CheckPrefix())
	})
	t.Run("With local connection", func(t *testing.T) {
		target := "hello-world"
		redisMock.Prefix = target
		RedisPrefix = ""
		assert.Equal(t, target, redisMock.CheckPrefix())
	})
}

func TestSetCompress(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf(miniRedisError, err)
	}
	redisMock := RedisDatabase{Host: mr.Addr()}
	redisMock.Connect()
	redisMock.client.WithTimeout(1 * time.Second)

	t.Run("Valid data", func(t *testing.T) {
		output := redisMock.SetCompress("test", "1", 1, "data")
		assert.Equal(t, true, output)
	})
	t.Run("Long data", func(t *testing.T) {
		long := ""
		for i := 0; i < 100000; i++ {
			long += "a"
		}
		output := redisMock.SetCompress("test", "2", 1, long)
		assert.Equal(t, true, output)
	})
	t.Run("Error disconnect", func(t *testing.T) {
		mr.Close()
		output := redisMock.Set("test", "1", 1, "data")
		assert.Equal(t, false, output)
	})
	t.Run("Error no connection compress", func(t *testing.T) {
		redisMock.client = nil
		RedisDB = nil
		output := redisMock.SetCompress("test", "1", 1, "data")
		assert.Equal(t, false, output)
	})
	t.Run("No client", func(t *testing.T) {
		redisMock.client = nil
		RedisDB = nil
		output := redisMock.Set("test", "1", 1, "data")
		assert.Equal(t, false, output)
	})
}

func TestGetCompress(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf(miniRedisError, err)
	}
	redisMock := RedisDatabase{Host: mr.Addr()}
	redisMock.Connect()
	redisMock.client.WithTimeout(1 * time.Second)
	redisMock.SetCompress("test", "1", 1, "data")

	long := ""
	for i := 0; i < 100000; i++ {
		long += "a"
	}
	redisMock.SetCompress("long", "1", 1, long)

	t.Run("Valid data", func(t *testing.T) {
		isOkay, data := redisMock.GetCompress("test", "1")
		assert.Equal(t, true, isOkay)
		assert.Equal(t, "data", data)
	})
	t.Run("Valid data decompress", func(t *testing.T) {
		isOkay, data := redisMock.GetCompress("long", "1")
		assert.Equal(t, true, isOkay)
		assert.Equal(t, long, data)
	})
	t.Run("No data", func(t *testing.T) {
		isOkay, data := redisMock.GetCompress("test", "2")
		assert.Equal(t, false, isOkay)
		assert.Equal(t, "", data)
	})
	t.Run("Error disconnect decompress", func(t *testing.T) {
		mr.Close()
		isOkay, _ := redisMock.GetCompress("test", "1")
		assert.Equal(t, false, isOkay)
	})
	t.Run("Disconnect server", func(t *testing.T) {
		mr.Close()
		isOkay, _ := redisMock.Get("test", "1")
		assert.Equal(t, false, isOkay)
	})
	t.Run("Error no client decompress", func(t *testing.T) {
		redisMock.client = nil
		RedisDB = nil
		isOkay, _ := redisMock.GetCompress("test", "1")
		assert.Equal(t, false, isOkay)
	})
	t.Run("Error no client", func(t *testing.T) {
		redisMock.client = nil
		RedisDB = nil
		isOkay, _ := redisMock.Get("test", "1")
		assert.Equal(t, false, isOkay)
	})
}

func TestRemove(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf(miniRedisError, err)
	}
	redisMock := RedisDatabase{Host: mr.Addr()}
	redisMock.Connect()
	redisMock.client.WithTimeout(1 * time.Second)
	redisMock.SetCompress("test", "1", 1, "data")
	redisMock.SetCompress("test", "2", 1, "data")
	redisMock.SetCompress("test", "3", 1, "data")
	redisMock.SetCompress("test", "4", 1, "data")

	t.Run("Remove data", func(t *testing.T) {
		isOkay := redisMock.Remove("test", "1")
		assert.Equal(t, true, isOkay)
	})
	t.Run("Remove manual", func(t *testing.T) {
		hash := redisMock.CheckPrefix() + utils.SeedName("test") + "|" + "2"
		err := redisMock.removeData(redisMock.client, hash)
		assert.Equal(t, nil, err)
	})
	t.Run("Remove data all", func(t *testing.T) {
		isOkay := redisMock.Remove("test", "")
		assert.Equal(t, true, isOkay)
	})
	t.Run("Error disconnect delete", func(t *testing.T) {
		redisMock.Connect()
		mr.Close()
		hash := redisMock.CheckPrefix() + utils.SeedName("test") + "|" + "2"
		err := redisMock.removeData(redisMock.client, hash)
		assert.NotEqual(t, nil, err)
	})
	t.Run("Disconnect when delete", func(t *testing.T) {
		mr.Close()
		isOkay := redisMock.Remove("test", "1")
		assert.Equal(t, false, isOkay)
	})
	t.Run("Error no client", func(t *testing.T) {
		redisMock.client = nil
		RedisDB = nil
		isOkay := redisMock.Remove("test", "1")
		assert.Equal(t, false, isOkay)
	})
}
