package messageserver

import (
	message "github.com/2751997nam/go-helpers/message"
)

type MessageFunc func(data map[string]any) (message.MessageResponse, error)

type MessageHandle struct {
	Url    string
	Method string
	Handle MessageFunc
}

type MessageRouter struct {
	Handles []MessageHandle
}

func (r *MessageRouter) GET(url string, handle MessageFunc) {
	r.Handles = append(r.Handles, MessageHandle{
		Url:    url,
		Method: "GET",
		Handle: handle,
	})
}

func (r *MessageRouter) POST(url string, handle MessageFunc) {
	r.Handles = append(r.Handles, MessageHandle{
		Url:    url,
		Method: "POST",
		Handle: handle,
	})
}

func (r *MessageRouter) PUT(url string, handle MessageFunc) {
	r.Handles = append(r.Handles, MessageHandle{
		Url:    url,
		Method: "PUT",
		Handle: handle,
	})
}

func (r *MessageRouter) PATCH(url string, handle MessageFunc) {
	r.Handles = append(r.Handles, MessageHandle{
		Url:    url,
		Method: "PUT",
		Handle: handle,
	})
}

func (r *MessageRouter) DELETE(url string, handle MessageFunc) {
	r.Handles = append(r.Handles, MessageHandle{
		Url:    url,
		Method: "DELETE",
		Handle: handle,
	})
}
