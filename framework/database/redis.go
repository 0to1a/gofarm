package database

import (
	"bytes"
	"context"
	"encoding/json"
	"framework/app/structure"
	"framework/framework/utils"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"time"
)

func RedisConnect() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     structure.SystemConf.RedisHost,
		Password: structure.SystemConf.RedisPassword,
		DB:       structure.SystemConf.RedisDatabase,
	})

	_, err := RedisDB.Ping(context.Background()).Result()
	if err != nil {
		RedisDB = nil
		return
	}
}

func RedisCacheSet(urlPath string, payload string, timeInMinutes int, data string) bool {
	if RedisDB == nil {
		return false
	}

	hash := SeedName(urlPath) + "|" + SeedName(payload)

	err := RedisDB.Set(context.Background(), hash, data, time.Duration(timeInMinutes)*time.Minute).Err()
	if err != nil {
		log.Println("redis", err)
		utils.LogError("redis:" + err.Error())
		return false
	}

	return true
}

func RedisCacheGet(urlPath string, payload string) (bool, string) {
	if RedisDB == nil {
		return false, ""
	}

	hash := SeedName(urlPath) + "|" + SeedName(payload)
	res, err := RedisDB.Get(context.Background(), hash).Result()
	if err == redis.Nil {
		return false, ""
	} else if err != nil {
		log.Println("redis", err)
		utils.LogError("redis:" + err.Error())
		return false, ""
	}

	return true, res
}

func RedisCacheJson(c echo.Context, timeInMinutes int, data map[string]interface{}) bool {
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	jsonByte, _ := json.Marshal(data)
	return RedisCacheSet(c.Request().URL.Path, string(bodyBytes), timeInMinutes, string(jsonByte))
}

func RedisCacheJsonRead(c echo.Context) (bool, map[string]interface{}) {
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	isCache, jsonString := RedisCacheGet(c.Request().URL.Path, string(bodyBytes))
	if !isCache {
		return false, nil
	}
	jsonMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(jsonString), &jsonMap)
	return true, jsonMap
}

func RedisCacheRemove(urlPath string, payload string) bool {
	if RedisDB == nil {
		return false
	}

	hash := SeedName(urlPath) + "|"
	if payload == "" {
		hash += "*"
		iter := RedisDB.Scan(context.Background(), 0, hash, 0).Iterator()
		for iter.Next(context.Background()) {
			err := RedisDB.Del(context.Background(), iter.Val()).Err()
			if err != nil {
				log.Println("redis", err)
				utils.LogError("redis:" + err.Error())
				return false
			}
		}
	} else {
		hash += SeedName(payload)
		err := RedisDB.Del(context.Background(), hash).Err()
		if err != nil {
			log.Println("redis", err)
			utils.LogError("redis:" + err.Error())
			return false
		}
	}
	return true
}
