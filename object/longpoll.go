package object

type LongpollMessage struct {
	Type    string                 `json:"type"`
	Object  map[string]interface{} `json:"object"`
	GroupID int                    `json:"group_id"`
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
