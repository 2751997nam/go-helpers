package messageserver

import (
	"fmt"

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
	r.Handles[fmt.Sprintf("%s_%s", url, "GET")] = MessageHandle{
		Type:   url,
		Method: "GET",
		Handle: handle,
	}
}

func (r *MessageRouter) POST(url string, handle MessageFunc) {
	r.Handles[fmt.Sprintf("%s_%s", url, "POST")] = MessageHandle{
		Type:   url,
		Method: "POST",
		Handle: handle,
	}
}

func (r *MessageRouter) PUT(url string, handle MessageFunc) {
	r.Handles[fmt.Sprintf("%s_%s", url, "PUT")] = MessageHandle{
		Type:   url,
		Method: "PUT",
		Handle: handle,
	}
}

func (r *MessageRouter) PATCH(url string, handle MessageFunc) {
	r.Handles[fmt.Sprintf("%s_%s", url, "PATCH")] = MessageHandle{
		Type:   url,
		Method: "PATCH",
		Handle: handle,
	}
}

func (r *MessageRouter) DELETE(url string, handle MessageFunc) {
	r.Handles[fmt.Sprintf("%s_%s", url, "DELETE")] = MessageHandle{
		Type:   url,
		Method: "DELETE",
		Handle: handle,
	}
}
