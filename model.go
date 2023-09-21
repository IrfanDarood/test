package web
import (
	"time"

)

type Response struct{
	// Status int64 `json:"status"`
	// Message string `json:"message"`
	Sum time.Duration `json:"sum"`
	// Data interface{} `json:"data"`
}