package object

type Message struct {
	Text            string          `json:"message"`
	RandomID        int             `json:"random_id"`
	UserID          int             `json:"user_id"`
	PeerID          int             `json:"peer_id"`
	PeerIDs         string          `json:"peer_ids"`
	Domain          string          `json:"domain"`
	ChatID          int             `json:"chat_id"`
	UserIDs         string          `json:"user_ids"`
	Guid            int             `json:"guid"`
	Lat             string          `json:"lat"`
	Long            string          `json:"long"`
	Attachment      string          `json:"attachment"`
	ReplyTo         int             `json:"reply_to"`
	ForwardMessages string          `json:"forward_messages"`
	Forward         MessageForward  `json:"forward"`
	StickerID       int             `json:"sticker_id"`
	GroupID         int             `json:"group_id"`
	Keyboard        MessageKeyboard `json:"keyboard"`
	Template        MessageTemplate `json:"template"`
	Payload         string          `json:"payload"`
	ContentSource   string          `json:"content_source"`
	Intent          string          `json:"intent"`
	SubscribeID     int             `json:"subscribe_id"`
	DontParseLinks  int             `json:"dont_parse_links,omitempty"`
	DisableMentions int             `json:"disable_mentions,omitempty"`
}

type MessageForward struct {
	OwnerID                int   `json:"owner_id"`
	PeerID                 int   `json:"peer_id"`
	ConservationMessageIDs []int `json:"conversation_message_ids"`
	MessageIDs             []int `json:"message_ids"`
	IsReply                bool  `json:"is_reply"`
}
