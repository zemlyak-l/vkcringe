package object

type Message struct {
	Text            string           `schema:"message,omitempty"`
	RandomID        int              `schema:"random_id"`
	UserID          int              `schema:"user_id,omitempty"`
	PeerID          int              `schema:"peer_id,omitempty"`
	PeerIDs         string           `schema:"peer_ids,omitempty"`
	Domain          string           `schema:"domain,omitempty"`
	ChatID          int              `schema:"chat_id,omitempty"`
	UserIDs         string           `schema:"user_ids,omitempty"`
	Guid            int              `schema:"guid,omitempty"`
	Lat             string           `schema:"lat,omitempty"`
	Long            string           `schema:"long,omitempty"`
	Attachment      string           `schema:"attachment,omitempty"`
	ReplyTo         int              `schema:"reply_to,omitempty"`
	ForwardMessages string           `schema:"forward_messages,omitempty"`
	Forward         *MessageForward  `schema:"forward,omitempty"`
	StickerID       int              `schema:"sticker_id,omitempty"`
	GroupID         int              `schema:"group_id,omitempty"`
	Keyboard        *MessageKeyboard `schema:"keyboard,omitempty"`
	Template        *MessageTemplate `schema:"template,omitempty"`
	Payload         string           `schema:"payload,omitempty"`
	ContentSource   string           `schema:"content_source,omitempty"`
	Intent          string           `schema:"intent,omitempty"`
	SubscribeID     int              `schema:"subscribe_id,omitempty"`
	DontParseLinks  int              `schema:"dont_parse_links,omitempty"`
	DisableMentions int              `schema:"disable_mentions,omitempty"`
}

type MessageForward struct {
	OwnerID                int   `schema:"owner_id"`
	PeerID                 int   `schema:"peer_id"`
	ConservationMessageIDs []int `schema:"conversation_message_ids"`
	MessageIDs             []int `schema:"message_ids"`
	IsReply                bool  `schema:"is_reply"`
}

type MessageKeyboard struct {
	AuthorID int                       `schema:"author_id,omitempty"`
	Buttons  [][]MessageKeyboardButton `schema:"buttons"`
	OneTime  bool                      `schema:"one_time,omitempty"`
	Inline   bool                      `schema:"inline,omitempty"`
}

type MessageKeyboardButton struct {
	Action MessageButtonAction `schema:"action"`
	Color  string              `schema:"color,omitempty"`
}

type MessageButtonAction struct {
	AppID   int    `schema:"app_id,omitempty"`
	Hash    string `schema:"hash,omitempty"`
	Label   string `schema:"label,omitempty"`
	OwnerID int    `schema:"owner_id,omitempty"`
	Payload string `schema:"payload,omitempty"`
	Type    string `schema:"type"`
	Link    string `schema:"link,omitempty"`
}

type MessageTemplate struct {
	Type     string            `schema:"type"`
	Elements []TemplateElement `schema:"elements"`
}

type MessageTemplateElement struct {
	MessageElementCarousel
}

type MessageElementCarousel struct {
	Title       string                  `schema:"title,omitempty"`
	Action      MessageCarouselAction   `schema:"action,omitempty"`
	Description string                  `schema:"description,omitempty"`
	Photo       *Photo                  `schema:"photo,omitempty"`
	PhotoID     string                  `schema:"photo_id,omitempty"`
	Buttons     []MessageKeyboardButton `schema:"buttons,omitempty"`
}

type MessageCarouselAction struct {
	Type string `schema:"type"`
	Link string `schema:"link,omitempty"`
}

type MessagePhoto struct {
	AccessKey          string              `schema:"access_key"`
	AlbumID            int                 `schema:"album_id"`
	Date               int                 `schema:"date"`
	Height             int                 `schema:"height"`
	ID                 int                 `schema:"id"`
	Images             []MessageImage      `schema:"images"`
	Lat                float64             `schema:"lat"`
	Long               float64             `schema:"long"`
	OwnerID            int                 `schema:"owner_id"`
	PostID             int                 `schema:"post_id"`
	Text               string              `schema:"text"`
	UserID             int                 `schema:"user_id"`
	Width              int                 `schema:"width"`
	CanUpload          bool                `schema:"can_upload"`
	CommentsDisabled   bool                `schema:"comments_disabled"`
	ThumbIsLast        bool                `schema:"thumb_is_last"`
	UploadByAdminsOnly bool                `schema:"upload_by_admins_only"`
	HasTags            bool                `schema:"has_tags"`
	Created            int                 `schema:"created"`
	Description        string              `schema:"description"`
	PrivacyComment     []string            `schema:"privacy_comment"`
	PrivacyView        []string            `schema:"privacy_view"`
	Size               int                 `schema:"size"`
	Sizes              []MessagePhotoSizes `schema:"sizes"`
	ThumbID            int                 `schema:"thumb_id"`
	ThumbSrc           string              `schema:"thumb_src"`
	Title              string              `schema:"title"`
	Updated            int                 `schema:"updated"`
	Color              string              `schema:"color"`
}

type MessagePhotoSizes struct {
	MessageBaseImage
}

type MessageImage struct {
	MessageBaseImage
	Type string `schema:"type"`
}

type MessageBaseImage struct {
	Height float64 `schema:"height"`
	URL    string  `schema:"url"`
	Width  float64 `schema:"width"`
	Type   string  `schema:"type"`
}
