package sdkdemo

import "fmt"

type SDKError struct {
	Code      string
	Message   string
	RequestID string // 非常重要，用于排查服务端问题
}

func (e *SDKError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}
