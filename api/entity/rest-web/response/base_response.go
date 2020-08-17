package response

import "time"

type BaseResponse struct {
	ServerTime	time.Time	`json:"server_time"`
	Code 		int64		`json:"code"`
	Message 	string		`json:"message"`
	Data 		interface{}	`json:"data"`
}
