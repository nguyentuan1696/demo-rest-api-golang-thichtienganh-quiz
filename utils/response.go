package utils

type ResponseData struct {
	Signal    int         `json:"signal"`
	ErrorCode int         `json:"error_code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}
