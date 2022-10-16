package framework

import (
	"github.com/imroc/req/v3"
	"log"
	"strings"
	"time"
)

type RequestClient struct {
	Client *req.Client
}

func NewRequestClient(agentName string, timeout int, retry int, customRetryHook *func(resp *req.Response, err error), customRetryCondition *func(resp *req.Response, err error) bool) *RequestClient {
	if agentName == "" {
		agentName = "gofarm-client"
	}
	if timeout <= 0 {
		timeout = 15
	}
	if retry <= 0 {
		retry = 1
	}
	client := req.C().
		SetUserAgent(agentName).
		SetCommonRetryCount(retry).
		SetCommonRetryFixedInterval(3 * time.Second).
		SetTimeout(time.Duration(timeout) * time.Second)

	retryFunction := funcRetryHook
	if customRetryHook != nil {
		retryFunction = *customRetryHook
	}
	client.AddCommonRetryHook(retryFunction)

	conditionFunction := funcRetryCondition
	if customRetryHook != nil {
		conditionFunction = *customRetryCondition
	}
	client.SetCommonRetryCondition(conditionFunction)

	return &RequestClient{Client: client}
}

func funcRetryHook(resp *req.Response, err error) {
	req := resp.Request.RawRequest
	escapedMethod := strings.Replace(req.Method, "\n", "", -1)
	escapedMethod = strings.Replace(escapedMethod, "\r", "", -1)
	escapedUrl := strings.Replace(req.URL.String(), "\n", "", -1)
	escapedUrl = strings.Replace(escapedUrl, "\r", "", -1)
	escapedAgent := strings.Replace(req.UserAgent(), "\n", "", -1)
	escapedAgent = strings.Replace(escapedAgent, "\r", "", -1)
	log.Println("Retry request:", escapedAgent, escapedMethod, escapedUrl, err)
}

func funcRetryCondition(resp *req.Response, err error) bool {
	return err != nil || resp.StatusCode >= 500
}

func (w *RequestClient) AddHeader(key string, value string) {
	w.Client.SetCommonHeader(key, value)
}

func (w *RequestClient) AddBearer(token string) {
	w.Client.SetCommonBearerAuthToken(token)
}
