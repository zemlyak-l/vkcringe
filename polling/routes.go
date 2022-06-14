package polling

import (
	"github.com/zemlyak-l/vkcringe/object"
)

type Routes struct {
	MessageNew                    func(object.NewMessage)
	MessageReply                  func(object.NewMessage)
	MessageEdit                   func(object.NewMessage)
	MessageAllow                  func(object.MessageAllowObject)
	MessageDeny                   func(object.MessageDenyObject)
	MessageTypingState            func(object.MessageTypingStateObject)
	MessageEvent                  func(object.MessageEventObject)
	PhotoNew                      func(object.Photo)
	PhotoCommentNew               func(object.PhotoCommentNewObject)
	PhotoCommentEdit              func(object.PhotoCommentEditObject)
	PhotoCommentRestore           func(object.PhotoCommentRestoreObject)
	PhotoCommentDelete            func(object.PhotoCommentDeleteObject)
	AudioNew                      func(object.AudioNewObject)
	VideoNew                      func(object.VideoNewObject)
	VideoCommentNew               func(object.VideoCommentNewObject)
	VideoCommentEdit              func(object.VideoCommentEditObject)
	VideoCommentRestore           func(object.VideoCommentRestoreObject)
	VideoCommentDelete            func(object.VideoCommentDeleteObject)
	WallPostNew                   func(object.WallPostNewObject)
	WallRepost                    func(object.WallRepostObject)
	WallReplyNew                  func(object.WallReplyNewObject)
	WallReplyEdit                 func(object.WallReplyEditObject)
	WallReplyRestore              func(object.WallReplyRestoreObject)
	WallReplyDelete               func(object.WallReplyDeleteObject)
	BoardPostNew                  func(object.BoardPostNewObject)
	BoardPostEdit                 func(object.BoardPostEditObject)
	BoardPostRestore              func(object.BoardPostRestoreObject)
	BoardPostDelete               func(object.BoardPostDeleteObject)
	MarketCommentNew              func(object.MarketCommentNewObject)
	MarketCommentEdit             func(object.MarketCommentEditObject)
	MarketCommentRestore          func(object.MarketCommentRestoreObject)
	MarketCommentDelete           func(object.MarketCommentDeleteObject)
	MarketOrderNew                func(object.MarketOrderNewObject)
	MarketOrderEdit               func(object.MarketOrderEditObject)
	GroupLeave                    func(object.GroupLeaveObject)
	GroupJoin                     func(object.GroupJoinObject)
	UserBlock                     func(object.UserBlockObject)
	UserUnblock                   func(object.UserUnblockObject)
	PollVoteNew                   func(object.PollVoteNewObject)
	GroupOfficersEdit             func(object.GroupOfficersEditObject)
	GroupChangeSettings           func(object.GroupChangeSettingsObject)
	GroupChangePhoto              func(object.GroupChangePhotoObject)
	VkpayTransaction              func(object.VkpayTransactionObject)
	LeadFormsNew                  func(object.LeadFormsNewObject)
	AppPayload                    func(object.AppPayloadObject)
	MessageRead                   func(object.MessageReadObject)
	LikeAdd                       func(object.LikeAddObject)
	LikeRemove                    func(object.LikeRemoveObject)
	DonutSubscriptionCreate       func(object.DonutSubscriptionCreateObject)
	DonutSubscriptionProlonged    func(object.DonutSubscriptionProlongedObject)
	DonutSubscriptionExpired      func(object.DonutSubscriptionExpiredObject)
	DonutSubscriptionCancelled    func(object.DonutSubscriptionCancelledObject)
	DonutSubscriptionPriceChanged func(object.DonutSubscriptionPriceChangedObject)
	DonutMoneyWithdraw            func(object.DonutMoneyWithdrawObject)
	DonutMoneyWithdrawError       func(object.DonutMoneyWithdrawErrorObject)
}
