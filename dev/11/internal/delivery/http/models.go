package http

import (
	"dev_11/internal/model"
	"encoding/json"
	"strings"
	"time"
)

type InputDate time.Time

type createEventInput struct {
	UserId      string    `json:"user_id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Date        InputDate `json:"date" validate:"required"`
}

type updateEventInput struct {
	UserId      string    `json:"user_id" validate:"required"`
	EventId     string    `json:"event_id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Date        InputDate `json:"date" validate:"required"`
}

type deleteEventInput struct {
	UserId  string `json:"user_id" validate:"required"`
	EventId string `json:"event_id" validate:"required"`
}

type successEventOutput struct {
	Result string `json:"result"`
}

type eventsOutput struct {
	Result []model.Event `json:"result"`
}

type errorOutput struct {
	Error string `json:"error"`
}

func newSuccessEventOutput(result string) successEventOutput {
	return successEventOutput{Result: result}
}

func newEventsOutput(result []model.Event) eventsOutput {
	return eventsOutput{Result: result}
}

func newErrorOutput(message string) errorOutput {
	return errorOutput{Error: message}
}

func (i *InputDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*i = InputDate(t)
	return nil
}

func (i InputDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(i))
}
