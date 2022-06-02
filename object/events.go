package object

import (
	"encoding/json"
)

type MessageAllowObject struct {
	UserID int    `json:"user_id"`
	Key    string `json:"key"`
}

type MessageDenyObject struct {
	UserID int `json:"user_id"`
}

type MessageTypingStateObject struct {
	State  string `json:"state"`
	FromID int    `json:"from_id"`
	ToID   int    `json:"to_id"`
}

type MessageEventObject struct {
	UserID                int             `json:"user_id"`
	PeerID                int             `json:"peer_id"`
	EventID               string          `json:"event_id"`
	Payload               json.RawMessage `json:"payload"`
	ConversationMessageID int             `json:"conversation_message_id"`
}

type PhotoCommentDeleteObject struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	PhotoID   int `json:"photo_id"`
}

type VideoCommentDeleteObject struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	VideoID   int `json:"video_id"`
}

type WallReplyDeleteObject struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	DeleterID int `json:"deleter_id"`
	PostID    int `json:"post_id"`
}

type BoardPostDeleteObject struct {
	TopicOwnerID int `json:"topic_owner_id"`
	TopicID      int `json:"topic_id"`
	ID           int `json:"id"`
}

type MarketCommentDeleteObject struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	ItemID    int `json:"item_id"`
}

type GroupLeaveObject struct {
	UserID int  `json:"user_id"`
	Self   Bool `json:"self"`
}

type GroupJoinObject struct {
	UserID   int    `json:"user_id"`
	JoinType string `json:"join_type"`
}

type UserBlockObject struct {
	AdminID     int    `json:"admin_id"`
	UserID      int    `json:"user_id"`
	UnblockDate int    `json:"unblock_date"`
	Reason      int    `json:"reason"`
	Comment     string `json:"comment"`
}

type UserUnblockObject struct {
	AdminID   int `json:"admin_id"`
	UserID    int `json:"user_id"`
	ByEndDate int `json:"by_end_date"`
}

type PollVoteNewObject struct {
	OwnerID  int `json:"owner_id"`
	PollID   int `json:"poll_id"`
	OptionID int `json:"option_id"`
	UserID   int `json:"user_id"`
}

type GroupOfficersEditObject struct {
	AdminID  int `json:"admin_id"`
	UserID   int `json:"user_id"`
	LevelOld int `json:"level_old"`
	LevelNew int `json:"level_new"`
}

type Changes struct {
	OldValue string `json:"old_value"`
	NewValue string `json:"new_value"`
}

type ChangesInt struct {
	OldValue int `json:"old_value"`
	NewValue int `json:"new_value"`
}

type GroupChangeSettingsObject struct {
	UserID  int `json:"user_id"`
	Changes struct {
		Title                 Changes    `json:"title"`
		Description           Changes    `json:"description"`
		Access                ChangesInt `json:"access"`
		ScreenName            Changes    `json:"screen_name"`
		PublicCategory        ChangesInt `json:"public_category"`
		PublicSubcategory     ChangesInt `json:"public_subcategory"`
		AgeLimits             ChangesInt `json:"age_limits"`
		Website               Changes    `json:"website"`
		StatusDefault         Changes    `json:"status_default"`
		Wall                  ChangesInt `json:"wall"`
		Replies               ChangesInt `json:"replies"`
		Topics                ChangesInt `json:"topics"`
		Audio                 ChangesInt `json:"audio"`
		Photos                ChangesInt `json:"photos"`
		Video                 ChangesInt `json:"video"`
		Market                ChangesInt `json:"market"`
		Docs                  ChangesInt `json:"docs"`
		Messages              ChangesInt `json:"messages"`
		EventGroupID          ChangesInt `json:"event_group_id"`
		Links                 Changes    `json:"links"`
		Email                 Changes    `json:"email"`
		EventStartDate        ChangesInt `json:"event_start_date::"`
		EventFinishDate       ChangesInt `json:"event_finish_date:"`
		Subject               Changes    `json:"subject"`
		MarketWiki            Changes    `json:"market_wiki"`
		DisableMarketComments ChangesInt `json:"disable_market_comments"`
		Phone                 ChangesInt `json:"phone"`
		CountryID             ChangesInt `json:"country_id"`
		CityID                ChangesInt `json:"city_id"`
	} `json:"Changes"`
}

type GroupChangePhotoObject struct {
	UserID int   `json:"user_id"`
	Photo  Photo `json:"photo"`
}

type VkpayTransactionObject struct {
	FromID      int    `json:"from_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Date        int    `json:"date"`
}

type LeadFormsNewObject struct {
	LeadID   int    `json:"lead_id"`
	GroupID  int    `json:"group_id"`
	UserID   int    `json:"user_id"`
	FormID   int    `json:"form_id"`
	FormName string `json:"form_name"`
	AdID     int    `json:"ad_id"`
	Answers  []struct {
		Key      string `json:"key"`
		Question string `json:"question"`
		Answer   string `json:"answer"`
	} `json:"answers"`
}

type AppPayloadObject struct {
	UserID  int    `json:"user_id"`
	AppID   int    `json:"app_id"`
	Payload string `json:"payload"`
}

type MessageReadObject struct {
	FromID        int `json:"from_id"`
	PeerID        int `json:"peer_id"`
	ReadMessageID int `json:"read_message_id"`
}

type LikeAddObject struct {
	LikerID       int    `json:"liker_id"`
	ObjectType    string `json:"object_type"`
	ObjectOwnerID int    `json:"object_owner_id"`
	ObjectID      int    `json:"object_id"`
	ThreadReplyID int    `json:"thread_reply_id"`
	PostID        int    `json:"post_id"`
}

type LikeRemoveObject struct {
	LikerID       int    `json:"liker_id"`
	ObjectType    string `json:"object_type"`
	ObjectOwnerID int    `json:"object_owner_id"`
	ObjectID      int    `json:"object_id"`
	ThreadReplyID int    `json:"thread_reply_id"`
	PostID        int    `json:"post_id"`
}

type DonutSubscriptionCreateObject struct {
	Amount           int     `json:"amount"`
	AmountWithoutFee float64 `json:"amount_without_fee"`
	UserID           int     `json:"user_id"`
}

type DonutSubscriptionProlongedObject struct {
	Amount           int     `json:"amount"`
	AmountWithoutFee float64 `json:"amount_without_fee"`
	UserID           int     `json:"user_id"`
}

type DonutSubscriptionExpiredObject struct {
	UserID int `json:"user_id"`
}

type DonutSubscriptionCancelledObject struct {
	UserID int `json:"user_id"`
}

type DonutSubscriptionPriceChangedObject struct {
	AmountOld            int     `json:"amount_old"`
	AmountNew            int     `json:"amount_new"`
	AmountDiff           float64 `json:"amount_diff"`
	AmountDiffWithoutFee float64 `json:"amount_diff_without_fee"`
	UserID               int     `json:"user_id"`
}

type DonutMoneyWithdrawObject struct {
	Amount           int     `json:"amount"`
	AmountWithoutFee float64 `json:"amount_without_fee"`
}

type DonutMoneyWithdrawErrorObject struct {
	Reason string `json:"reason"`
}
