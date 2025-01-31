package structs

import "time"

type JSONResponse struct {
	Timestamp time.Time   `json:"timestamp"`
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}
