package polling

import (
	"encoding/json"
	"log"

	"github.com/zemlyak-l/vkgottle/object"
)

const (
	Confirmation                  = "confirmation"
	MessageNew                    = "message_new"
	MessageReply                  = "message_reply"
	MessageEdit                   = "message_edit"
	MessageAllow                  = "message_allow"
	MessageDeny                   = "message_deny"
	MessageTypingState            = "message_typing_state"
	Message                       = "message_"
	PhotoNew                      = "photo_new"
	PhotoCommentNew               = "photo_comment_new"
	PhotoCommentEdit              = "photo_comment_edit"
	PhotoCommentRestore           = "photo_comment_restore"
	PhotoCommentDelete            = "photo_comment_delete"
	AudioNew                      = "audio_new"
	VideoNew                      = "video_new"
	VideoCommentNew               = "video_comment_new"
	VideoCommentEdit              = "video_comment_edit"
	VideoCommentRestore           = "video_comment_restore"
	VideoCommentDelete            = "video_comment_delete"
	WallPostNew                   = "wall_post_new"
	WallRepost                    = "wall_repost"
	WallReplyNew                  = "wall_reply_new"
	WallReplyEdit                 = "wall_reply_edit"
	WallReplyRestore              = "wall_reply_restore"
	WallReplyDelete               = "wall_reply_delete"
	BoardPostNew                  = "board_post_new"
	BoardPostEdit                 = "board_post_edit"
	BoardPostRestore              = "board_post_restore"
	BoardPostDelete               = "board_post_delete"
	MarketCommentNew              = "market_comment_new"
	MarketCommentEdit             = "market_comment_edit"
	MarketCommentRestore          = "market_comment_restore"
	MarketCommentDelete           = "market_comment_delete"
	MarketOrderNew                = "market_order_new"
	MarketOrderEdit               = "market_order_edit"
	GroupLeave                    = "group_leave"
	GroupJoin                     = "group_join"
	UserBlock                     = "user_block"
	UserUnblock                   = "user_unblock"
	PollVoteNew                   = "poll_vote_new"
	GroupOfficersEdit             = "group_officers_edit"
	GroupChangeSettings           = "group_change_settings"
	GroupChangePhoto              = "group_change_photo"
	VkpayTransaction              = "vkpay_transaction"
	LeadFormsNew                  = "lead_forms_new"
	AppPayload                    = "app_payload"
	MessageRead                   = "message_read"
	LikeAdd                       = "like_add"
	LikeRemove                    = "like_remove"
	DonutSubscriptionCreate       = "donut_subscription_create"
	DonutSubscriptionProlonged    = "donut_subscription_prolonged"
	DonutSubscriptionExpired      = "donut_subscription_expired"
	DonutSubscriptionCancelled    = "donut_subscription_cancelled"
	DonutSubscriptionPriceChanged = "donut_subscription_price_changed"
	DonutMoneyWithdraw            = "donut_money_withdraw"
	DonutMoneyWithdrawError       = "donut_money_withdraw_error"
)

func (lp *Longpoll) CheckEvents() {
	for event := range lp.NewEvent {
		if err := lp.checkCurrentEvent(event); err != nil {
			log.Fatal(err)
		}
	}
}

func (lp *Longpoll) checkCurrentEvent(event object.LongpollMessage) error {
	switch event.Type {
	case MessageNew:
		message := object.MessageJson{}
		if err := json.Unmarshal(event.Object, &message); err != nil {
			return err
		}
		if lp.Routes.MessageNew != nil {
			lp.Routes.MessageNew(message.CurrentMessage)
		}
	case MessageReply:
	case MessageEdit:
	case MessageAllow:
	case MessageDeny:
	case Message:
	case PhotoNew:
	case PhotoCommentNew:
	case PhotoCommentEdit:
	case PhotoCommentRestore:
	case PhotoCommentDelete:
	case AudioNew:
	case VideoNew:
	case VideoCommentNew:
	case VideoCommentEdit:
	case VideoCommentRestore:
	case VideoCommentDelete:
	case WallPostNew:
	case WallRepost:
	case WallReplyNew:
	case WallReplyEdit:
	case WallReplyRestore:
	case WallReplyDelete:
	case BoardPostNew:
	case BoardPostEdit:
	case BoardPostRestore:
	case BoardPostDelete:
	case MarketCommentNew:
	case MarketCommentEdit:
	case MarketCommentRestore:
	case MarketCommentDelete:
	case MarketOrderNew:
	case MarketOrderEdit:
	case GroupLeave:
	case GroupJoin:
	case UserBlock:
	case UserUnblock:
	case PollVoteNew:
	case GroupOfficersEdit:
	case GroupChangeSettings:
	case GroupChangePhoto:
	case VkpayTransaction:
	case LeadFormsNew:
	case AppPayload:
	case MessageRead:
	case LikeAdd:
	case LikeRemove:
	case DonutSubscriptionCreate:
	case DonutSubscriptionProlonged:
	case DonutSubscriptionExpired:
	case DonutSubscriptionCancelled:
	case DonutSubscriptionPriceChanged:
	case DonutMoneyWithdraw:
	case DonutMoneyWithdrawError:
	}
	return nil
}
