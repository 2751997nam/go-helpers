package messageserver

import (
	"fmt"
	"strings"

	message "github.com/2751997nam/go-helpers/message"
	"github.com/2751997nam/go-helpers/utils"
)

type MessageFunc func(data map[string]any) (message.MessageResponse, error)

type MessageHandle struct {
	Type   string
	Method string
	Handle MessageFunc
}

type MessageRouter struct {
	Handles map[string]MessageHandle
}

func (r *MessageRouter) GET(url string, handle MessageFunc) {
	r.registerRoute(url, "GET", handle)
}

func (r *MessageRouter) POST(url string, handle MessageFunc) {
	r.registerRoute(url, "POST", handle)
}

func (r *MessageRouter) PUT(url string, handle MessageFunc) {
	r.registerRoute(url, "PUT", handle)
}

func (r *MessageRouter) PATCH(url string, handle MessageFunc) {
	r.registerRoute(url, "PATCH", handle)
}

func (r *MessageRouter) DELETE(url string, handle MessageFunc) {
	r.registerRoute(url, "DELETE", handle)
}

func (r *MessageRouter) registerRoute(url string, method string, handle MessageFunc) {
	r.Handles[fmt.Sprintf("%s_%s", url, method)] = MessageHandle{
		Type:   url,
		Method: method,
		Handle: handle,
	}
}

func (r *MessageRouter) GetRoute(url string, method string) (*MessageHandle, map[string]string) {
	if handle, ok := r.Handles[fmt.Sprintf("%s_%s", url, method)]; ok {
		return &handle, nil
	}

	for key, handle := range r.Handles {
		index := strings.Index(key, "_")
		imethod := key[index+1:]
		key = key[0:index]
		if imethod == method && strings.Contains(key, ":") {
			params := map[string]string{}
			arr1 := strings.Split(strings.Trim(key, "/"), "/")
			utils.Log("arr1", arr1)
			arr2 := strings.Split(strings.Trim(url, "/"), "/")
			utils.Log("arr2", arr2)
			if len(arr1) == len(arr2) && len(arr1) > 0 {
				for index := range arr1 {
					if arr1[index][0:1] == ":" {
						params[arr1[index][1:]] = arr2[index]
					} else if arr1[index] != arr2[index] {
						return nil, nil
					}
				}
			} else {
				return nil, nil
			}
			return &handle, params
		}
	}

	return nil, nil
}
