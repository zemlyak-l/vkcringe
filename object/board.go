package object

type BoardTopicComment struct {
	Attachments []WallCommentAttachment `json:"attachments"`
	Date        int                     `json:"date"`
	FromID      int                     `json:"from_id"`
	ID          int                     `json:"id"`
	Text        string                  `json:"text"`
	Likes       LikesInfo               `json:"likes"`
	CanEdit     Bool                    `json:"can_edit"`
}

type BoardPostNewObject BoardTopicComment

type BoardPostEditObject BoardTopicComment

type BoardPostRestoreObject BoardTopicComment
