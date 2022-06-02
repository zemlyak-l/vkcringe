package object

type  WallWallpost struct {
	AccessKey      string                   `json:"access_key"` // Access key to private object
	ID             int                      `json:"id"`         // Post ID
	OwnerID        int                      `json:"owner_id"`   // Wall owner's ID
	FromID         int                      `json:"from_id"`    // Post author ID
	CreatedBy      int                      `json:"created_by"`
	Date           int                      `json:"date"` // Date of publishing in Unixtime
	Text           string                   `json:"text"` // Post text
	ReplyOwnerID   int                      `json:"reply_owner_id"`
	ReplyPostID    int                      `json:"reply_post_id"`
	FriendsOnly    int                      `json:"friends_only"`
	Comments       CommentsInfo         `json:"comments"`
	Likes          LikesInfo            `json:"likes"`   // Count of likes
	Reposts        RepostsInfo          `json:"reposts"` // Count of reposts
	Views          WallViews                `json:"views"`   // Count of views
	PostType       string                   `json:"post_type"`
	PostSource     WallPostSource           `json:"post_source"`
	Attachments    []WallWallpostAttachment `json:"attachments"`
	Geo            Geo                  `json:"geo"`
	SignerID       int                      `json:"signer_id"` // Post signer ID
	CopyHistory    []WallWallpost           `json:"copy_history"`
	CanPin         Bool              `json:"can_pin"`
	CanDelete      Bool              `json:"can_delete"`
	CanEdit        Bool              `json:"can_edit"`
	IsPinned       Bool              `json:"is_pinned"`
	IsFavorite     Bool              `json:"is_favorite"` // Information whether the post in favorites list
	IsArchived     Bool              `json:"is_archived"` // Is post archived, only for post owners
	IsDeleted      Bool              `json:"is_deleted"`
	MarkedAsAds    Bool              `json:"marked_as_ads"`
	Edited         int                      `json:"edited"` // Date of editing in Unixtime
	Copyright      WallPostCopyright        `json:"copyright"`
	PostID         int                      `json:"post_id"`
	ParentsStack   []int                    `json:"parents_stack"`
	Donut          WallWallpostDonut        `json:"donut"`
	ShortTextRate  float64                  `json:"short_text_rate"`
	CarouselOffset int                      `json:"carousel_offset"`
	Header         WallWallpostHeader       `json:"header"`
	Hash           string                   `json:"hash"`
}

type WallWallComment struct {
	Attachments    []WallCommentAttachment `json:"attachments"`
	Date           int                     `json:"date"` // Date when the comment has been added in Unixtime
	Deleted        Bool                    `json:"deleted"`
	FromID         int                     `json:"from_id"` // Author ID
	ID             int                     `json:"id"`      // Comment ID
	Likes          LikesInfo               `json:"likes"`
	RealOffset     int                     `json:"real_offset"`      // Real position of the comment
	ReplyToComment int                     `json:"reply_to_comment"` // Replied comment ID
	ReplyToUser    int                     `json:"reply_to_user"`    // Replied user ID
	Text           string                  `json:"text"`             // Comment text
	PostID         int                     `json:"post_id"`
	PostOwnerID    int                     `json:"post_owner_id"`
	PhotoID        int                     `json:"photo_id"`
	PhotoOwnerID   int                     `json:"photo_owner_id"`
	VideoID        int                     `json:"video_id"`
	VideoOwnerID   int                     `json:"video_owner_id"`
	ItemID         int                     `json:"item_id"`
	MarketOwnerID  int                     `json:"market_owner_id"`
	ParentsStack   []int                   `json:"parents_stack"`
	OwnerID        int                     `json:"owner_id"`
	Thread         WallWallCommentThread   `json:"thread"`
	Donut          WallWallCommentDonut    `json:"donut"`
}

type WallWallpostDonut struct {
	IsDonut            Bool          `json:"is_donut"`
	CanPublishFreeCopy Bool          `json:"can_publish_free_copy"`
	PaidDuration       int                  `json:"paid_duration"`
	EditMode           string               `json:"edit_mode"`
	Durations          []ObjectWithName `json:"durations"`
}

type WallWallCommentThread struct {
	Count           int               `json:"count"`
	Items           []WallWallComment `json:"items"`
	CanPost         Bool              `json:"can_post"`
	GroupsCanPost   Bool              `json:"groups_can_post"`
	ShowReplyButton Bool              `json:"show_reply_button"`
}

type WallWallCommentDonut struct {
	IsDonut     Bool   `json:"is_donut"`
	Placeholder string `json:"placeholder"`
}

type WallCommentAttachment struct {
	Audio    Audio        `json:"audio"`
	Doc      Doc               `json:"doc"`
	Link     Link              `json:"link"`
	Note     WallAttachedNote  `json:"note"`
	Page     PagesWikipageFull `json:"page"`
	Photo    Photo             `json:"photo"`
	Sticker  Sticker           `json:"sticker"`
	Type     string            `json:"type"`
	Video    Video             `json:"video"`
	Graffiti WallGraffiti      `json:"graffiti"`
}

type WallAttachedNote struct {
	Comments     int    `json:"comments"`      // Comments number
	Date         int    `json:"date"`          // Date when the note has been created in Unixtime
	ID           int    `json:"id"`            // Note ID
	OwnerID      int    `json:"owner_id"`      // Note owner's ID
	ReadComments int    `json:"read_comments"` // Read comments number
	Title        string `json:"title"`         // Note title
	ViewURL      string `json:"view_url"`      // URL of the page with note preview
}

type WallGraffiti struct {
	ID        int    `json:"id"`        // Graffiti ID
	OwnerID   int    `json:"owner_id"`  // Graffiti owner's ID
	Photo200  string `json:"photo_200"` // URL of the preview image with 200 px in width
	Photo586  string `json:"photo_586"` // URL of the preview image with 586 px in width
	URL       string `json:"url"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	AccessKey string `json:"access_key"`
}

type WallViews struct {
	Count int `json:"count"` // Count
}

type WallWallpostAttachment struct {
	AccessKey         string            `json:"access_key"` // Access key for the audio
	Album             PhotoAlbum  `json:"album"`
	App               WallAppPost       `json:"app"`
	Audio             Audio        `json:"audio"`
	Doc               Doc           `json:"doc"`
	Event             EventsEventAttach `json:"event"`
	Graffiti          WallGraffiti      `json:"graffiti"`
	Link              Link          `json:"link"`
	Note              WallAttachedNote  `json:"note"`
	Page              PagesWikipageFull `json:"page"`
	Photo             Photo       `json:"photo"`
	PhotosList        []string          `json:"photos_list"`
	Poll              Poll         `json:"poll"`
	PostedPhoto       WallPostedPhoto   `json:"posted_photo"`
	Type              string            `json:"type"`
	Video             Video        `json:"video"`
}

type WallPostSource struct {
	Link     Link `json:"link"`
	Data     string   `json:"data"`     // Additional data
	Platform string   `json:"platform"` // Platform name
	Type     string   `json:"type"`
	URL      string   `json:"url"` // URL to an external site used to publish the post
}

type WallPostedPhoto struct {
	ID       int    `json:"id"`        // Photo ID
	OwnerID  int    `json:"owner_id"`  // Photo owner's ID
	Photo130 string `json:"photo_130"` // URL of the preview image with 130 px in width
	Photo604 string `json:"photo_604"` // URL of the preview image with 604 px in width
}

type WallPostCopyright struct {
	ID   int    `json:"id,omitempty"`
	Link string `json:"link"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type WallAppPost struct {
	ID       int    `json:"id"`        // Application ID
	Name     string `json:"name"`      // Application name
	Photo130 string `json:"photo_130"` // URL of the preview image with 130 px in width
	Photo604 string `json:"photo_604"` // URL of the preview image with 604 px in width
}


type WallWallpostHeader struct {
	Type              string                              `json:"type"`
	CustomDescription WallWallpostHeaderCustomDescription `json:"custom_description"`
}

// WallWallpostHeaderCustomDescription struct.
type WallWallpostHeaderCustomDescription struct {
	SourceID int `json:"source_id"`
	Date     int `json:"date"`
}

type WallPostNewObject WallWallpost

type WallRepostObject WallWallpost

type WallReplyNewObject WallWallComment

type WallReplyEditObject WallWallComment

type WallReplyRestoreObject WallWallComment
