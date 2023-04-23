package messageserver

import (
	"fmt"
	"strconv"
	"strings"

	message "github.com/2751997nam/go-helpers/message"
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
			arr2 := strings.Split(strings.Trim(url, "/"), "/")
			if len(arr1) == len(arr2) && len(arr1) > 0 && IsSameRoute(arr1, arr2) {
				for index := range arr1 {
					if arr1[index][0:1] == ":" {
						params[arr1[index][1:]] = arr2[index]
					} else if arr1[index] != arr2[index] {
						break
					}
				}
				return &handle, params
			}
		}
	}

	return nil, nil
}

func IsSameRoute(keyArr, requestArr []string) bool {
	for key := range keyArr {
		if len(keyArr[key]) == 0 {
			return false
		}
		if keyArr[key][0:1] == "*" || keyArr[key][0:1] == ":" {
			if keyArr[key] == ":id" {
				if _, err := strconv.ParseInt(requestArr[key], 10, 64); err == nil {
					continue
				} else {
					return false
				}
			}
			continue
		}
		if keyArr[key] != requestArr[key] {
			return false
		}
	}

	return true
}
