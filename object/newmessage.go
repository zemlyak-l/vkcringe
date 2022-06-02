package object

type NewMessage struct {
	AdminAuthorID int `json:"admin_author_id"`
	Action        struct {
		MemberID int    `json:"member_id"`
		Type     string `json:"type"`
	} `json:"action"`
	Attachments           []Attachments   `json:"attachments"`
	ConversationMessageID int             `json:"conversation_message_id"`
	Date                  int             `json:"date"`
	FromID                int             `json:"from_id"`
	FwdMessages           []NewMessage    `json:"fwd_Messages"`
	ReplyMessage          *NewMessage     `json:"reply_message"`
	Geo                   bool            `json:"geo"`
	PinnedAt              int             `json:"pinned_at,omitempty"`
	ID                    int             `json:"id"`
	Deleted               bool            `json:"deleted"`
	Important             bool            `json:"important"`
	IsHidden              bool            `json:"is_hidden"`
	IsCropped             bool            `json:"is_cropped"`
	IsSilent              bool            `json:"is_silent"`
	Out                   bool            `json:"out"`
	WasListened           bool            `json:"was_listened,omitempty"`
	Payload               string          `json:"payload"`
	PeerID                int             `json:"peer_id"`
	RandomID              int             `json:"random_id"`
	Ref                   string          `json:"ref"`
	RefSource             string          `json:"ref_source"`
	Text                  string          `json:"text"`
	UpdateTime            int             `json:"update_time"`
	MembersCount          int             `json:"members_count"`
	ExpireTTL             int             `json:"expire_ttl"`
	MessageTag            string          `json:"message_tag"`
	Template              MessageTemplate `json:"template"`
}

type MessageJson struct {
	CurrentMessage NewMessage `json:"message"`
}
