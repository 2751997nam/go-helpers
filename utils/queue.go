package utils

import (
	"encoding/json"

	"github.com/2751997nam/go-helpers/event"
	"github.com/2751997nam/go-helpers/libs"
)

func SendToQueue(name string, data any) error {
	emitter, err := event.NewEventEmitter(libs.GetMQ())
	if err != nil {
		return err
	}
	payload := event.Event{
		Name: name,
		Data: data,
	}

	j, _ := json.Marshal(&payload)
	err = emitter.Push(string(j))

	if err != nil {
		return err
	}

	return nil
}
