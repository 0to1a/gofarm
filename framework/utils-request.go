package framework

import (
	"github.com/imroc/req/v3"
	"log"
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
	log.Println("Retry request:", req.UserAgent(), req.Method, req.URL, err)
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
