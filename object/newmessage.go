package object

type NewMessage struct {
	CmdArgs []string `json:"-"`

	AdminAuthorID int `json:"admin_author_id"`
	Action        struct {
		MemberID int    `json:"member_id"`
		Type     string `json:"type"`
	} `json:"action"`
	Attachments           []Attachments `json:"attachments"`
	ConversationMessageID int           `json:"conversation_message_id"`
	Date                  int           `json:"date"`
	FromID                int           `json:"from_id"`
	FwdMessages           []NewMessage  `json:"fwd_Messages"`
	ReplyMessage          *NewMessage   `json:"reply_message"`
	Geo                   Bool          `json:"geo"`
	PinnedAt              int           `json:"pinned_at,omitempty"`
	ID                    int           `json:"id"`
	Deleted               Bool          `json:"deleted"`
	Important             Bool          `json:"important"`
	IsHidden              Bool          `json:"is_hidden"`
	IsCropped             Bool          `json:"is_cropped"`
	IsSilent              Bool          `json:"is_silent"`
	Out                   int           `json:"out"`
	WasListened           Bool          `json:"was_listened,omitempty"`
	Payload               string        `json:"payload"`
	PeerID                int           `json:"peer_id"`
	RandomID              int           `json:"random_id"`
	Ref                   string        `json:"ref"`
	RefSource             string        `json:"ref_source"`
	Text                  string        `json:"text"`
	UpdateTime            int           `json:"update_time"`
	MembersCount          int           `json:"members_count"`
	ExpireTTL             int           `json:"expire_ttl"`
	MessageTag            string        `json:"message_tag"`
	Template              Template      `json:"template"`
}

type MessageJson struct {
	CurrentMessage NewMessage `json:"message"`
}
