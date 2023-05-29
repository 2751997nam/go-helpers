package utils

import (
	"net/http"

	"github.com/2751997nam/go-helpers/message"
	ms "github.com/2751997nam/go-helpers/message_server"
	"github.com/gin-gonic/gin"
)

type HandleFunc func(data map[string]any) Response

func RestWrapH(handle HandleFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := GetRequestData(c)
		if err != nil {
			ResponseFail(c, err.Error(), http.StatusInternalServerError)
			return
		}
		res := handle(data)
		if len(res.Status) == 0 {
			res.Status = "successful"
		}
		if res.StatusCode == 0 {
			res.StatusCode = http.StatusOK
		}
		c.JSON(int(res.StatusCode), res)
	}
}

func GRPCWrapH(handle HandleFunc) ms.MessageFunc {
	return func(data map[string]any) (message.MessageResponse, error) {
		res := handle(data)
		if len(res.Status) == 0 {
			res.Status = "successful"
		}
		if res.StatusCode == 0 {
			res.StatusCode = http.StatusOK
		}
		return ResponseMessage(res)
	}
}
