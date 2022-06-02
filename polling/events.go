package polling

import (
	"encoding/json"

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
	MessageEvent                  = "message_"
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

func (lp *Longpoll) CheckEvent(event object.LongpollMessage) error {
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
		message := object.NewMessage{}
		if err := json.Unmarshal(event.Object, &message); err != nil {
			return err
		}
		if lp.Routes.MessageReply != nil {
			lp.Routes.MessageReply(message)
		}
	case MessageEdit:
		message := object.NewMessage{}
		if err := json.Unmarshal(event.Object, &message); err != nil {
			return err
		}
		if lp.Routes.MessageEdit != nil {
			lp.Routes.MessageEdit(message)
		}
	case MessageAllow:
		obj := object.MessageAllowObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.MessageAllow != nil {
			lp.Routes.MessageAllow(obj)
		}
	case MessageDeny:
		obj := object.MessageDenyObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.MessageDeny != nil {
			lp.Routes.MessageDeny(obj)
		}
	case MessageEvent:
		obj := object.MessageEventObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.MessageEvent != nil {
			lp.Routes.MessageEvent(obj)
		}
	case PhotoNew:
		obj := object.Photo{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.PhotoNew != nil {
			lp.Routes.PhotoNew(obj)
		}
	case PhotoCommentNew:
		obj := object.PhotoCommentNewObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.PhotoCommentNew != nil {
			lp.Routes.PhotoCommentNew(obj)
		}
	case PhotoCommentEdit:
		obj := object.PhotoCommentEditObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.PhotoCommentEdit != nil {
			lp.Routes.PhotoCommentEdit(obj)
		}
	case PhotoCommentRestore:
		obj := object.PhotoCommentRestoreObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.PhotoCommentRestore != nil {
			lp.Routes.PhotoCommentRestore(obj)
		}
	case PhotoCommentDelete:
		obj := object.PhotoCommentDeleteObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.PhotoCommentDelete != nil {
			lp.Routes.PhotoCommentDelete(obj)
		}
	case AudioNew:
		obj := object.AudioNewObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.AudioNew != nil {
			lp.Routes.AudioNew(obj)
		}
	case VideoNew:
		obj := object.VideoNewObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.VideoNew != nil {
			lp.Routes.VideoNew(obj)
		}
	case VideoCommentNew:
		obj := object.VideoCommentNewObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.VideoCommentNew != nil {
			lp.Routes.VideoCommentNew(obj)
		}
	case VideoCommentEdit:
		obj := object.VideoCommentEditObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.VideoCommentEdit != nil {
			lp.Routes.VideoCommentEdit(obj)
		}
	case VideoCommentRestore:
		obj := object.VideoCommentRestoreObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.VideoCommentRestore != nil {
			lp.Routes.VideoCommentRestore(obj)
		}
	case VideoCommentDelete:
		obj := object.VideoCommentDeleteObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.VideoCommentDelete != nil {
			lp.Routes.VideoCommentDelete(obj)
		}
	case WallPostNew:
		obj := object.WallPostNewObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.WallPostNew != nil {
			lp.Routes.WallPostNew(obj)
		}
	case WallRepost:
		obj := object.WallRepostObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.WallRepost != nil {
			lp.Routes.WallRepost(obj)
		}
	case WallReplyNew:
		obj := object.WallReplyNewObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.WallReplyNew != nil {
			lp.Routes.WallReplyNew(obj)
		}
	case WallReplyEdit:
		obj := object.WallReplyEditObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.WallReplyEdit != nil {
			lp.Routes.WallReplyEdit(obj)
		}
	case WallReplyRestore:
		obj := object.WallReplyRestoreObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.WallReplyRestore != nil {
			lp.Routes.WallReplyRestore(obj)
		}
	case WallReplyDelete:
		obj := object.WallReplyDeleteObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.WallReplyDelete != nil {
			lp.Routes.WallReplyDelete(obj)
		}
	case BoardPostNew:
		obj := object.BoardPostNewObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.BoardPostNew != nil {
			lp.Routes.BoardPostNew(obj)
		}
	case BoardPostEdit:
		obj := object.BoardPostEditObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.BoardPostEdit != nil {
			lp.Routes.BoardPostEdit(obj)
		}
	case BoardPostRestore:
		obj := object.BoardPostRestoreObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.BoardPostRestore != nil {
			lp.Routes.BoardPostRestore(obj)
		}
	case BoardPostDelete:
		obj := object.BoardPostDeleteObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.BoardPostDelete != nil {
			lp.Routes.BoardPostDelete(obj)
		}
	case MarketCommentNew:
		obj := object.MarketCommentNewObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.MarketCommentNew != nil {
			lp.Routes.MarketCommentNew(obj)
		}
	case MarketCommentEdit:
		obj := object.MarketCommentEditObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.MarketCommentEdit != nil {
			lp.Routes.MarketCommentEdit(obj)
		}
	case MarketCommentRestore:
		obj := object.MarketCommentRestoreObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.MarketCommentRestore != nil {
			lp.Routes.MarketCommentRestore(obj)
		}
	case MarketCommentDelete:
		obj := object.MarketCommentDeleteObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.MarketCommentDelete != nil {
			lp.Routes.MarketCommentDelete(obj)
		}
	case MarketOrderNew:
		obj := object.MarketOrderNewObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.MarketOrderNew != nil {
			lp.Routes.MarketOrderNew(obj)
		}
	case MarketOrderEdit:
		obj := object.MarketOrderEditObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.MarketOrderEdit != nil {
			lp.Routes.MarketOrderEdit(obj)
		}
	case GroupLeave:
		obj := object.GroupLeaveObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.GroupLeave != nil {
			lp.Routes.GroupLeave(obj)
		}
	case GroupJoin:
		obj := object.GroupJoinObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.GroupJoin != nil {
			lp.Routes.GroupJoin(obj)
		}
	case UserBlock:
		obj := object.UserBlockObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.UserBlock != nil {
			lp.Routes.UserBlock(obj)
		}
	case UserUnblock:
		obj := object.UserUnblockObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.UserUnblock != nil {
			lp.Routes.UserUnblock(obj)
		}
	case PollVoteNew:
		obj := object.PollVoteNewObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.PollVoteNew != nil {
			lp.Routes.PollVoteNew(obj)
		}
	case GroupOfficersEdit:
		obj := object.GroupOfficersEditObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.GroupOfficersEdit != nil {
			lp.Routes.GroupOfficersEdit(obj)
		}
	case GroupChangeSettings:
		obj := object.GroupChangeSettingsObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.GroupChangeSettings != nil {
			lp.Routes.GroupChangeSettings(obj)
		}
	case GroupChangePhoto:
		obj := object.GroupChangePhotoObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.GroupChangePhoto != nil {
			lp.Routes.GroupChangePhoto(obj)
		}
	case VkpayTransaction:
		obj := object.VkpayTransactionObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.VkpayTransaction != nil {
			lp.Routes.VkpayTransaction(obj)
		}
	case LeadFormsNew:
		obj := object.LeadFormsNewObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.LeadFormsNew != nil {
			lp.Routes.LeadFormsNew(obj)
		}
	case AppPayload:
		obj := object.AppPayloadObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.AppPayload != nil {
			lp.Routes.AppPayload(obj)
		}
	case MessageRead:
		obj := object.MessageReadObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.MessageRead != nil {
			lp.Routes.MessageRead(obj)
		}
	case LikeAdd:
		obj := object.LikeAddObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.LikeAdd != nil {
			lp.Routes.LikeAdd(obj)
		}
	case LikeRemove:
		obj := object.LikeRemoveObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.LikeRemove != nil {
			lp.Routes.LikeRemove(obj)
		}
	case DonutSubscriptionCreate:
		obj := object.DonutSubscriptionCreateObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.DonutSubscriptionCreate != nil {
			lp.Routes.DonutSubscriptionCreate(obj)
		}
	case DonutSubscriptionProlonged:
		obj := object.DonutSubscriptionProlongedObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.DonutSubscriptionProlonged != nil {
			lp.Routes.DonutSubscriptionProlonged(obj)
		}
	case DonutSubscriptionExpired:
		obj := object.DonutSubscriptionExpiredObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.DonutSubscriptionExpired != nil {
			lp.Routes.DonutSubscriptionExpired(obj)
		}
	case DonutSubscriptionCancelled:
		obj := object.DonutSubscriptionCancelledObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.DonutSubscriptionCancelled != nil {
			lp.Routes.DonutSubscriptionCancelled(obj)
		}
	case DonutSubscriptionPriceChanged:
		obj := object.DonutSubscriptionPriceChangedObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.DonutSubscriptionPriceChanged != nil {
			lp.Routes.DonutSubscriptionPriceChanged(obj)
		}
	case DonutMoneyWithdraw:
		obj := object.DonutMoneyWithdrawObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.DonutMoneyWithdraw != nil {
			lp.Routes.DonutMoneyWithdraw(obj)
		}
	case DonutMoneyWithdrawError:
		obj := object.DonutMoneyWithdrawErrorObject{}
		if err := json.Unmarshal(event.Object, &obj); err != nil {
			return err
		}
		if lp.Routes.DonutMoneyWithdrawError != nil {
			lp.Routes.DonutMoneyWithdrawError(obj)
		}
	}
	return nil
}
