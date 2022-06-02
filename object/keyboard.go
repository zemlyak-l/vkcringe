package object

type Keyboard struct {
	AuthorID int                `json:"author_id,omitempty"`
	Buttons  [][]KeyboardButton `json:"buttons"`
	OneTime  Bool               `json:"one_time,omitempty"`
	Inline   Bool               `json:"inline,omitempty"`
}

type KeyboardButton struct {
	Action ButtonAction `json:"action"`
	Color  string       `json:"color,omitempty"`
}

type ButtonAction struct {
	AppID   int    `json:"app_id,omitempty"`
	Hash    string `json:"hash,omitempty"`
	Label   string `json:"label,omitempty"`
	OwnerID int    `json:"owner_id,omitempty"`
	Payload string `json:"payload,omitempty"`
	Type    string `json:"type"`
	Link    string `json:"link,omitempty"`
}
