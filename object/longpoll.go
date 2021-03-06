package object

import "encoding/json"

type EventHandler struct {
	HandlerFunc func(event interface{})
}

type LongpollRoutes map[string]*EventHandler

type GetServerMessage struct {
	GroupID int `schema:"group_id"`
}

type LongpollMessage struct {
	Type    string          `json:"type"`
	Object  json.RawMessage `json:"object"`
	GroupID int             `json:"group_id"`
	EventID string          `json:"event_id"`
	Secret  string          `json:"secret"`
}

type LongpollResponse struct {
	Ts      string            `json:"ts"`
	Updates []LongpollMessage `json:"updates"`
}

type ResponseInit struct {
	Response struct {
		Key    string `json:"key"`
		Server string `json:"server"`
		Ts     string `json:"ts"`
	} `json:"response"`
}
