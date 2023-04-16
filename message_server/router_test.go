package messageserver

import (
	"log"
	"testing"

	message "github.com/2751997nam/go-helpers/message"
)

func handle(data map[string]any) (message.MessageResponse, error) {
	return message.MessageResponse{}, nil
}
func TestGetRoute(t *testing.T) {
	router := MessageRouter{
		Handles: map[string]MessageHandle{},
	}

	router.GET("/find", handle)
	router.GET("/", handle)
	router.PUT("/:id", handle)

	handle, params := router.GetRoute("/6438dcdfddecf5127e5832fe", "PUT")
	log.Println("handle", handle)
	log.Println("params", params)
	if handle == nil {
		t.Errorf("handle is nil 1")
	}

	handle, params = router.GetRoute("/test/6438dcdfddecf5127e5832fe/oke", "GET")
	log.Println("handle", handle)
	log.Println("params", params)
	if handle == nil {
		t.Errorf("handle is nil 2")
	}
}
