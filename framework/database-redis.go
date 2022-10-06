package framework

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"time"
)

var (
	RedisDB     *redis.Client
	RedisPrefix string
)

type RedisDatabase struct {
	Prefix   string
	Host     string
	Password string
	Database int
	client   *redis.Client
}

func (w RedisDatabase) Connect() *redis.Client {
	var clientRedis *redis.Client
	clientRedis = redis.NewClient(&redis.Options{
		Addr:     w.Host,
		Password: w.Password,
		DB:       w.Database,
	})
	RedisPrefix = w.Prefix

	_, err := clientRedis.Ping(context.Background()).Result()
	if err != nil {
		log.Println("redis:", err)
		RedisDB = nil
		return nil
	}
	w.client = clientRedis
	RedisDB = clientRedis
	return clientRedis
}

func (w RedisDatabase) CheckClient() *redis.Client {
	if RedisDB == nil && w.client == nil {
		return nil
	} else if w.client != nil {
		return w.client
	} else {
		return RedisDB
	}
}

func (w RedisDatabase) CheckPrefix() string {
	if RedisPrefix == "" && w.Prefix == "" {
		return ""
	} else if w.Prefix != "" {
		return w.Prefix
	} else {
		return RedisPrefix
	}
}

func (w RedisDatabase) Set(urlPath string, payload string, timeInMinutes int, data string) bool {
	var clientRedis *redis.Client
	if clientRedis = w.CheckClient(); clientRedis == nil {
		return false
	}

	hash := w.CheckPrefix() + utils.SeedName(urlPath) + "|" + utils.SeedName(payload)

	err := clientRedis.Set(context.Background(), hash, data, time.Duration(timeInMinutes)*time.Minute).Err()
	if err != nil {
		log.Println("redis", err)
		return false
	}

	return true
}

func (w RedisDatabase) SetCompress(urlPath string, payload string, timeInMinutes int, data string) bool {
	var clientRedis *redis.Client
	if clientRedis = w.CheckClient(); clientRedis == nil {
		return false
	}

	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(data)); err != nil {
		log.Println("gzip", err)
	}
	if err := gz.Flush(); err != nil {
		log.Println("gzip", err)
	}
	if err := gz.Close(); err != nil {
		log.Println("gzip", err)
	}

	str := base64.StdEncoding.EncodeToString(b.Bytes())
	if len(str) >= len(data) {
		return w.Set(urlPath, payload, timeInMinutes, data)
	}

	return w.Set(urlPath, payload, timeInMinutes, str)
}

func (w RedisDatabase) Get(urlPath string, payload string) (bool, string) {
	var clientRedis *redis.Client
	if clientRedis = w.CheckClient(); clientRedis == nil {
		return false, ""
	}

	hash := w.CheckPrefix() + utils.SeedName(urlPath) + "|" + utils.SeedName(payload)
	res, err := clientRedis.Get(context.Background(), hash).Result()
	if err == redis.Nil {
		return false, ""
	} else if err != nil {
		log.Println("redis", err)
		return false, ""
	}

	return true, res
}

func (w RedisDatabase) GetCompress(urlPath string, payload string) (bool, string) {
	isOkay, res := w.Get(urlPath, payload)
	if !isOkay {
		return isOkay, res
	}

	data, _ := base64.StdEncoding.DecodeString(res)
	rdata := bytes.NewReader(data)
	r, err := gzip.NewReader(rdata)
	if err != nil {
		return isOkay, res
	}
	decompress, err := ioutil.ReadAll(r)

	return isOkay, string(decompress)
}

func (w RedisDatabase) SetJson(c echo.Context, timeInMinutes int, data map[string]interface{}) error {
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	jsonByte, _ := json.Marshal(data)
	w.SetCompress(c.Request().URL.Path, string(bodyBytes), timeInMinutes, string(jsonByte))

	return webserver.ResultAPIFromJson(c, data)
}

func (w RedisDatabase) GetJson(c echo.Context) error {
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	isCache, jsonString := w.GetCompress(c.Request().URL.Path, string(bodyBytes))
	if !isCache {
		return errors.New("no cache")
	}

	jsonMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(jsonString), &jsonMap)

	return webserver.ResultAPIFromJson(c, jsonMap)
}

func (w RedisDatabase) Remove(urlPath string, payload string) bool {
	var clientRedis *redis.Client
	if clientRedis = w.CheckClient(); clientRedis == nil {
		return false
	}

	hash := w.CheckPrefix() + utils.SeedName(urlPath) + "|"
	if payload == "" {
		hash += "*"
		iter := clientRedis.Scan(context.Background(), 0, hash, 0).Iterator()
		for iter.Next(context.Background()) {
			err := clientRedis.Del(context.Background(), iter.Val()).Err()
			if err != nil {
				log.Println("redis", err)
				return false
			}
		}
	} else {
		hash += utils.SeedName(payload)
		err := clientRedis.Del(context.Background(), hash).Err()
		if err != nil {
			log.Println("redis", err)
			return false
		}
	}
	return true
}
