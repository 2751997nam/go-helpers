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
	router.GET("/:id", handle)
	router.GET("/show/:id", handle)
	router.GET("/view/:id", handle)
	router.GET("/product_sku/:id", handle)
	router.PUT("/:id", handle)

	handle, params := router.GetRoute("/123123", "GET")
	log.Println("handle", handle)
	log.Println("params", params)
	if handle == nil {
		t.Errorf("handle is nil 1")
	}

	handle, params = router.GetRoute("/show/12312", "GET")
	log.Println("handle", handle)
	log.Println("params", params)
	if handle == nil {
		t.Errorf("handle is nil 2")
	}

	handle, params = router.GetRoute("/view/12312", "GET")
	log.Println("handle", handle)
	log.Println("params", params)
	if handle == nil {
		t.Errorf("handle is nil 2")
	}

	handle, params = router.GetRoute("/product_sku/2131869259", "GET")
	log.Println("handle", handle)
	log.Println("params", params)
	if handle == nil {
		t.Errorf("handle is nil 333s")
	}

	handle, params = router.GetRoute("/test/6438dcdfddecf5127e5832fe/oke", "GET")
	log.Println("handle", handle)
	log.Println("params", params)
	if handle == nil {
		t.Errorf("handle is nil 3")
	}
}
