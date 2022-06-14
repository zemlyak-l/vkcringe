package object

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

type NewGetByID struct {
	GroupIDs string `schema:"group_ids,omitempty"`
	GroupID  int    `schema:"group_id,omitempty"`
	Fields   string `schema:"fields,omitempty"`
}

type GroupResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ScreenName  string `json:"screen_name"`
	IsClosed    int    `json:"is_closed"`
	Deactivated string `json:"deactivated"`
	IsAdmin     int    `json:"is_admin"`
	AdminLevel  int    `json:"admin_level"`
	IsMember    int    `json:"is_member"`
	IsAdvestier int    `json:"is_advertiser"`
	InvitedBy   int    `json:"invited_by"`
	Type        string `json:"type"`
	Photo50     string `json:"photo_50"`
	Photo100    string `json:"photo_100"`
	Photo200    string `json:"photo_200"`
}

type GroupResponseSlice struct {
	Response []GroupResponse `json:"response"`
}

// GroupsAddress WorkInfoStatus of information about timetable.
const (
	WorkStatusNoInformation     = "no_information"
	WorkStatusTemporarilyClosed = "temporarily_closed"
	WorkStatusAlwaysOpened      = "always_opened"
	WorkStatusTimetable         = "timetable"
	WorkStatusForeverClosed     = "forever_closed"
)

// GroupsAddress struct.
type GroupsAddress struct {
	// Additional address to the place (6 floor, left door).
	AdditionalAddress string                 `json:"additional_address"`
	Address           string                 `json:"address"`          // String address to the place (Nevsky, 28)
	CityID            int                    `json:"city_id"`          // City id of address
	CountryID         int                    `json:"country_id"`       // Country id of address
	Distance          int                    `json:"distance"`         // Distance from the point
	ID                int                    `json:"id"`               // Address id
	Latitude          float64                `json:"latitude"`         // Address latitude
	Longitude         float64                `json:"longitude"`        // Address longitude
	MetroStationID    int                    `json:"metro_station_id"` // Metro id of address
	Phone             string                 `json:"phone"`            // Address phone
	TimeOffset        int                    `json:"time_offset"`      // Time offset int minutes from utc time
	Timetable         GroupsAddressTimetable `json:"timetable"`        // Week timetable for the address
	Title             string                 `json:"title"`            // Title of the place (Zinger, etc)
	WorkInfoStatus    string                 `json:"work_info_status"` // Status of information about timetable
	PlaceID           int                    `json:"place_id"`
}

// GroupsAddressTimetable Timetable for a week.
type GroupsAddressTimetable struct {
	Fri GroupsAddressTimetableDay `json:"fri"` // Timetable for friday
	Mon GroupsAddressTimetableDay `json:"mon"` // Timetable for monday
	Sat GroupsAddressTimetableDay `json:"sat"` // Timetable for saturday
	Sun GroupsAddressTimetableDay `json:"sun"` // Timetable for sunday
	Thu GroupsAddressTimetableDay `json:"thu"` // Timetable for thursday
	Tue GroupsAddressTimetableDay `json:"tue"` // Timetable for tuesday
	Wed GroupsAddressTimetableDay `json:"wed"` // Timetable for wednesday
}

// GroupsAddressTimetableDay Timetable for one day.
type GroupsAddressTimetableDay struct {
	BreakCloseTime int `json:"break_close_time"` // Close time of the break in minutes
	BreakOpenTime  int `json:"break_open_time"`  // Start time of the break in minutes
	CloseTime      int `json:"close_time"`       // Close time in minutes
	OpenTime       int `json:"open_time"`        // Open time in minutes
}

// GroupsAddressesInfo struct.
type GroupsAddressesInfo struct {
	IsEnabled     Bool `json:"is_enabled"`      // Information whether addresses is enabled
	MainAddressID int  `json:"main_address_id"` // Main address id for group
}

// Group AdminLevel type.
const (
	GroupsAdminLevelModerator = iota
	GroupsAdminLevelEditor
	GroupsAdminLevelAdministrator
)

// Group MainSection type.
const (
	GroupsMainSectionAbsent = iota
	GroupsMainSectionPhotos
	GroupsMainSectionTopics
	GroupsMainSectionAudio
	GroupsMainSectionVideo
	GroupsMainSectionMarket
)

// Group MemberStatus(events_event_attach, newsfeed_event_activity).
const (
	GroupsMemberStatusNotMember = iota
	GroupsMemberStatusMember
	GroupsMemberStatusNotSure
	GroupsMemberStatusDeclined
	GroupsMemberStatusHasSentRequest
	GroupsMemberStatusInvited
)

// Group Access or IsClosed type.
const (
	GroupOpen = iota
	GroupClosed
	GroupPrivate
)

// Group AgeLimits.
const (
	GroupsAgeLimitsNo = iota
	GroupsAgeLimitsOver16
	GroupsAgeLimitsOver18
)

// Group type.
const (
	GroupsTypeGroup = "group"
	GroupsTypePage  = "page"
	GroupsTypeEvent = "event"
)

// Group struct.
type Group struct {
	ID           int              `json:"id"`          // Community ID
	Name         string           `json:"name"`        // Community name
	ScreenName   string           `json:"screen_name"` // Domain of the community page
	Type         string           `json:"type"`
	IsClosed     int              `json:"is_closed"`
	AdminLevel   int              `json:"admin_level,omitempty"`
	Deactivated  string           `json:"deactivated,omitempty"` // Information whether community is banned
	FinishDate   int              `json:"finish_date,omitempty"` // Finish date in Unixtime format
	Photo100     string           `json:"photo_100,omitempty"`   // URL of square photo of the community with 100 pixels in width
	Photo200     string           `json:"photo_200,omitempty"`   // URL of square photo of the community with 200 pixels in width
	Photo50      string           `json:"photo_50,omitempty"`    // URL of square photo of the community with 50 pixels in width
	StartDate    int              `json:"start_date,omitempty"`  // Start date in Unixtime format
	Market       GroupsMarketInfo `json:"market,omitempty"`
	MemberStatus int              `json:"member_status,omitempty"` // Current user's member status
	City         Object           `json:"city,omitempty"`
	Country      Country          `json:"country,omitempty"`

	// Information whether current user is administrator.
	IsAdmin Bool `json:"is_admin"`

	// Information whether current user is advertiser.
	IsAdvertiser Bool `json:"is_advertiser,omitempty"`

	// Information whether current user is member.
	IsMember Bool `json:"is_member,omitempty"`

	// Information whether community is in faves.
	IsFavorite Bool `json:"is_favorite,omitempty"`

	// Information whether community is adult.
	IsAdult Bool `json:"is_adult,omitempty"`

	// Information whether current user is subscribed.
	IsSubscribed Bool `json:"is_subscribed,omitempty"`

	// Information whether current user can post on community's wall.
	CanPost Bool `json:"can_post,omitempty"`

	// Information whether current user can see all posts on community's wall.
	CanSeeAllPosts Bool `json:"can_see_all_posts,omitempty"`

	// Information whether current user can create topic.
	CanCreateTopic Bool `json:"can_create_topic,omitempty"`

	// Information whether current user can upload video.
	CanUploadVideo Bool `json:"can_upload_video,omitempty"`

	// Information whether current user can upload doc.
	CanUploadDoc Bool `json:"can_upload_doc,omitempty"`

	// Information whether community has photo.
	HasPhoto Bool `json:"has_photo,omitempty"`

	// Information whether current user can send a message to community.
	CanMessage Bool `json:"can_message,omitempty"`

	// Information whether community can send a message to current user.
	IsMessagesBlocked Bool `json:"is_messages_blocked,omitempty"`

	// Information whether community can send notifications by phone number to current user.
	CanSendNotify Bool `json:"can_send_notify,omitempty"`

	// Information whether current user is subscribed to podcasts.
	IsSubscribedPodcasts Bool `json:"is_subscribed_podcasts,omitempty"`

	// Owner in whitelist or not.
	CanSubscribePodcasts Bool `json:"can_subscribe_podcasts,omitempty"`

	// Can subscribe to wall.
	CanSubscribePosts Bool `json:"can_subscribe_posts,omitempty"`

	// Information whether community has market app.
	HasMarketApp        Bool `json:"has_market_app,omitempty"`
	IsHiddenFromFeed    Bool `json:"is_hidden_from_feed,omitempty"`
	IsMarketCartEnabled Bool `json:"is_market_cart_enabled,omitempty"`
	Verified            Bool `json:"verified,omitempty"` // Information whether community is verified

	// Information whether the community has a fire pictogram.
	Trending     Bool                `json:"trending,omitempty"`
	Description  string              `json:"description,omitempty"`   // Community description
	WikiPage     string              `json:"wiki_page,omitempty"`     // Community's main wiki page title
	MembersCount int                 `json:"members_count,omitempty"` // Community members number
	Counters     GroupsCountersGroup `json:"counters,omitempty"`
	Cover        GroupsCover         `json:"cover,omitempty"`

	// Type of group, start date of event or category of public page.
	Activity        string               `json:"activity,omitempty"`
	FixedPost       int                  `json:"fixed_post,omitempty"`    // Fixed post ID
	Status          string               `json:"status,omitempty"`        // Community status
	MainAlbumID     int                  `json:"main_album_id,omitempty"` // Community's main photo album ID
	Links           []GroupsLinksItem    `json:"links,omitempty"`
	Contacts        []GroupsContactsItem `json:"contacts,omitempty"`
	Site            string               `json:"site,omitempty"` // Community's website
	MainSection     int                  `json:"main_section,omitempty"`
	OnlineStatus    GroupsOnlineStatus   `json:"online_status,omitempty"` // Status of replies in community messages
	AgeLimits       int                  `json:"age_limits,omitempty"`    // Information whether age limit
	BanInfo         GroupBanInfo         `json:"ban_info,omitempty"`      // User ban info
	Addresses       GroupsAddressesInfo  `json:"addresses,omitempty"`     // Info about addresses in Groups
	LiveCovers      GroupsLiveCovers     `json:"live_covers,omitempty"`
	CropPhoto       UsersCropPhoto       `json:"crop_photo,omitempty"`
	Wall            int                  `json:"wall,omitempty"`
	ActionButton    GroupsActionButton   `json:"action_button,omitempty"`
	TrackCode       string               `json:"track_code,omitempty"`
	PublicDateLabel string               `json:"public_date_label,omitempty"`
	AuthorID        int                  `json:"author_id,omitempty"`
	Phone           string               `json:"phone,omitempty"`
	Like            GroupLike            `json:"like"`
}

// ToMention return mention.
func (group *Group) ToMention() string {
	return fmt.Sprintf("[club%d|%s]", group.ID, group.Name)
}

// GroupLike struct.
type GroupLike struct {
	IsLiked Bool             `json:"is_liked"`
	Friends GroupLikeFriends `json:"friends"`
}

// GroupLikeFriends struct.
type GroupLikeFriends struct {
	Count   int   `json:"count"`
	Preview []int `json:"preview"`
}

// GroupsLiveCovers struct.
type GroupsLiveCovers struct {
	IsEnabled  Bool     `json:"is_enabled"`
	IsScalable Bool     `json:"is_scalable"`
	StoryIds   []string `json:"story_ids"`
}

// GroupsBanInfo reason type.
const (
	GroupsBanReasonOther = iota
	GroupsBanReasonSpam
	GroupsBanReasonVerbalAbuse
	GroupsBanReasonStrongLanguage
	GroupsBanReasonFlood
)

// GroupsBanInfo struct.
type GroupsBanInfo struct {
	AdminID        int    `json:"admin_id"` // Administrator ID
	Comment        string `json:"comment"`  // Comment for a ban
	Date           int    `json:"date"`     // Date when user has been added to blacklist in Unixtime
	EndDate        int    `json:"end_date"` // Date when user will be removed from blacklist in Unixtime
	Reason         int    `json:"reason"`
	CommentVisible Bool   `json:"comment_visible"`
}

// GroupsCallbackServer struct.
type GroupsCallbackServer struct {
	CreatorID int    `json:"creator_id"`
	ID        int    `json:"id"`
	SecretKey string `json:"secret_key"`
	Status    string `json:"status"`
	Title     string `json:"title"`
	URL       string `json:"url"`
}

// GroupsCallbackSettings struct.
type GroupsCallbackSettings struct {
	APIVersion string               `json:"api_version"` // API version used for the events
	Events     GroupsLongPollEvents `json:"events"`
}

// GroupsContactsItem struct.
type GroupsContactsItem struct {
	Desc   string `json:"desc"`    // Contact description
	Email  string `json:"email"`   // Contact email
	Phone  string `json:"phone"`   // Contact phone
	UserID int    `json:"user_id"` // User ID
}

// GroupsCountersGroup struct.
type GroupsCountersGroup struct {
	Addresses      int `json:"addresses"`       // Addresses number
	Albums         int `json:"albums"`          // Photo albums number
	Articles       int `json:"articles"`        // Articles number
	Audios         int `json:"audios"`          // Audios number
	Docs           int `json:"docs"`            // Docs number
	Market         int `json:"market"`          // Market items number
	Photos         int `json:"photos"`          // Photos number
	Topics         int `json:"topics"`          // Topics number
	Videos         int `json:"videos"`          // Videos number
	Narratives     int `json:"narratives"`      // Narratives number
	Clips          int `json:"clips"`           // Clips number
	ClipsFollowers int `json:"clips_followers"` // Clips followers number
}

// UnmarshalJSON GroupsCountersGroup.
//
// BUG(VK): GroupsCountersGroup return [].
func (personal *GroupsCountersGroup) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("[]")) {
		return nil
	}

	type renamedGroupsCountersGroup GroupsCountersGroup

	var r renamedGroupsCountersGroup

	err := json.Unmarshal(data, &r)
	if err != nil {
		return err
	}

	*personal = GroupsCountersGroup(r)

	return nil
}

// GroupsCover struct.
type GroupsCover struct {
	Enabled Bool    `json:"enabled"` // Information whether cover is enabled
	Images  []Image `json:"images"`
}

// GroupBanInfo struct.
type GroupBanInfo struct {
	Comment string `json:"comment"`  // Ban comment
	EndDate int    `json:"end_date"` // End date of ban in Unixtime
}

// GroupCategory struct.
type GroupCategory struct {
	ID            int              `json:"id"`   // Category ID
	Name          string           `json:"name"` // Category name
	Subcategories []ObjectWithName `json:"subcategories"`
}

// GroupCategoryFull struct.
type GroupCategoryFull struct {
	ID            int                 `json:"id"`         // Category ID
	Name          string              `json:"name"`       // Category name
	PageCount     int                 `json:"page_count"` // Pages number
	PagePreviews  []Group             `json:"page_previews"`
	Subcategories []GroupCategoryFull `json:"subcategories"`
}

// GroupCategoryType struct.
type GroupCategoryType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GroupLink struct.
type GroupLink struct {
	Desc            string `json:"desc"`             // Link description
	EditTitle       Bool   `json:"edit_title"`       // Information whether the title can be edited
	ImageProcessing Bool   `json:"image_processing"` // Information whether the image on processing
	Name            string `json:"name"`
	ID              int    `json:"id"`  // Link ID
	URL             string `json:"url"` // Link URL
}

// GroupPublicCategoryList struct.
type GroupPublicCategoryList struct {
	ID            int                 `json:"id"`
	Name          string              `json:"name"`
	Subcategories []GroupCategoryType `json:"subcategories"`
}

// GroupSettings Photos type.
const (
	GroupPhotosDisabled = iota
	GroupPhotosOpen
	GroupPhotosLimited
)

// GroupSettings Subject type.
const (
	_ = iota
	GroupSubjectAuto
	GroupSubjectActivityHolidays
	GroupSubjectBusiness
	GroupSubjectPets
	GroupSubjectHealth
	GroupSubjectDatingAndCommunication
	GroupSubjectGames
	GroupSubjectIt
	GroupSubjectCinema
	GroupSubjectBeautyAndFashion
	GroupSubjectCooking
	GroupSubjectArtAndCulture
	GroupSubjectLiterature
	GroupSubjectMobileServicesAndInternet
	GroupSubjectMusic
	GroupSubjectScienceAndTechnology
	GroupSubjectRealEstate
	GroupSubjectNewsAndMedia
	GroupSubjectSecurity
	GroupSubjectEducation
	GroupSubjectHomeAndRenovations
	GroupSubjectPolitics
	GroupSubjectFood
	GroupSubjectIndustry
	GroupSubjectTravel
	GroupSubjectWork
	GroupSubjectEntertainment
	GroupSubjectReligion
	GroupSubjectFamily
	GroupSubjectSports
	GroupSubjectInsurance
	GroupSubjectTelevision
	GroupSubjectGoodsAndServices
	GroupSubjectHobbies
	GroupSubjectFinance
	GroupSubjectPhoto
	GroupSubjectEsoterics
	GroupSubjectElectronicsAndAppliances
	GroupSubjectErotic
	GroupSubjectHumor
	GroupSubjectSocietyHumanities
	GroupSubjectDesignAndGraphics
)

// GroupSettings Topics type.
const (
	GroupTopicsDisabled = iota
	GroupTopicsOpen
	GroupTopicsLimited
)

// GroupSettings Docs type.
const (
	GroupDocsDisabled = iota
	GroupDocsOpen
	GroupDocsLimited
)

// GroupSettings Audio type.
const (
	GroupAudioDisabled = iota
	GroupAudioOpen
	GroupAudioLimited
)

// GroupSettings Video type.
const (
	GroupVideoDisabled = iota
	GroupVideoOpen
	GroupVideoLimited
)

// GroupSettings Wall type.
const (
	GroupWallDisabled = iota
	GroupWallOpen
	GroupWallLimited
	GroupWallClosed
)

// GroupSettings Wiki type.
const (
	GroupWikiDisabled = iota
	GroupWikiOpen
	GroupWikiLimited
)

// GroupSettings struct.
type GroupSettings struct {
	Access             int                       `json:"access"`          // Community access settings
	Address            string                    `json:"address"`         // Community's page domain
	Audio              int                       `json:"audio"`           // Audio settings
	Description        string                    `json:"description"`     // Community description
	Docs               int                       `json:"docs"`            // Docs settings
	ObsceneWords       []string                  `json:"obscene_words"`   // The list of stop words
	Photos             int                       `json:"photos"`          // Photos settings
	PublicCategory     int                       `json:"public_category"` // Information about the group category
	PublicCategoryList []GroupPublicCategoryList `json:"public_category_list"`

	// Information about the group subcategory.
	PublicSubcategory int                 `json:"public_subcategory"`
	Rss               string              `json:"rss"`     // URL of the RSS feed
	Subject           int                 `json:"subject"` // Community subject ID
	SubjectList       []GroupsSubjectItem `json:"subject_list"`
	Title             string              `json:"title"`   // Community title
	Topics            int                 `json:"topics"`  // Topics settings
	Video             int                 `json:"video"`   // Video settings
	Wall              int                 `json:"wall"`    // Wall settings
	Website           string              `json:"website"` // Community website
	Wiki              int                 `json:"wiki"`    // Wiki settings
	CountryID         int                 `json:"country_id"`
	CityID            int                 `json:"city_id"`
	Messages          int                 `json:"messages"`
	Articles          int                 `json:"articles"`
	Events            int                 `json:"events"`
	AgeLimits         int                 `json:"age_limits"`

	// Information whether the obscene filter is enabled.
	ObsceneFilter Bool `json:"obscene_filter"`

	// Information whether the stopwords filter is enabled.
	ObsceneStopwords Bool `json:"obscene_stopwords"`
	LiveCovers       struct {
		IsEnabled Bool `json:"is_enabled"`
	} `json:"live_covers"`
	Market           GroupsMarketInfo     `json:"market"`
	SectionsList     []GroupsSectionsList `json:"sections_list"`
	MainSection      int                  `json:"main_section"`
	SecondarySection int                  `json:"secondary_section"`
	ActionButton     GroupsActionButton   `json:"action_button"`
	Phone            string               `json:"phone"`

	RecognizePhoto int `json:"recognize_photo"`

	MarketServices GroupsMarketServices `json:"market_services"`
	Narratives     int                  `json:"narratives"`
	Clips          int                  `json:"clips"`
	Textlives      int                  `json:"textlives"`
	Youla          GroupsYoula          `json:"youla"`
}

// GroupsMarketServices struct.
type GroupsMarketServices struct {
	Enabled         Bool                `json:"enabled"`
	CanMessage      Bool                `json:"can_message"`
	CommentsEnabled Bool                `json:"comments_enabled"`
	ContactID       int                 `json:"contact_id"`
	Currency        MarketCurrency      `json:"currency"`
	ViewType        GroupsSelectedItems `json:"view_type"`
	BlockName       GroupsSelectedItems `json:"block_name"`
	ButtonLabel     GroupsSelectedItems `json:"button_label"`
}

// GroupsSelectedItems struct.
type GroupsSelectedItems struct {
	SelectedItemID int64            `json:"selected_item_id"`
	Items          []ObjectWithName `json:"items"`
}

// GroupsYoula struct.
type GroupsYoula struct {
	CategoryTree  GroupsYoulaCategory `json:"category_tree"`
	GroupSettings GroupsYoulaSettings `json:"group_settings"`
}

// GroupsYoulaCategory struct.
type GroupsYoulaCategory struct {
	ID            int                      `json:"id"`
	Title         string                   `json:"title"`
	Subcategories []GroupsYoulaSubcategory `json:"subcategories"`
}

// GroupsYoulaSubcategory struct.
type GroupsYoulaSubcategory struct {
	ID            int                      `json:"id"`
	Title         string                   `json:"title"`
	ParentID      int                      `json:"parent_id"`
	Subcategories []GroupsYoulaSubcategory `json:"subcategories"`
}

// GroupsYoulaSettings struct.
type GroupsYoulaSettings struct {
	IsActive              Bool      `json:"is_active"`
	IsModerated           Bool      `json:"is_moderated"`
	ShowModerationSetting Bool      `json:"show_moderation_setting"`
	ModerationStatus      int       `json:"moderation_status"`
	DeclineReason         string    `json:"decline_reason"`
	GroupMode             int       `json:"group_mode"`
	SelectedCategoryIDS   []int     `json:"selected_category_ids"`
	Lat                   float64   `json:"lat"`
	Long                  float64   `json:"long"`
	Radius                float64   `json:"radius"`
	RadiusArea            string    `json:"radius_area"`
	Address               string    `json:"address"`
	Radiuses              []float64 `json:"radiuses"`
}

// GroupsSectionsList struct.
type GroupsSectionsList struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UnmarshalJSON need for unmarshal dynamic array (Example: [1, "Фотографии"]) to struct.
//
// To unmarshal JSON into a value implementing the Unmarshaler interface,
// Unmarshal calls that value's UnmarshalJSON method.
// See more https://golang.org/pkg/encoding/json/#Unmarshal
func (g *GroupsSectionsList) UnmarshalJSON(data []byte) error {
	var alias []interface{}
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	if len(alias) != 2 {
		return &json.UnmarshalTypeError{
			Value: string(data),
			Type:  reflect.TypeOf((*GroupsSectionsList)(nil)),
		}
	}

	// default concrete Go type float64 for JSON numbers
	id, ok := alias[0].(float64)
	if !ok {
		return &json.UnmarshalTypeError{
			Value:  string(data),
			Type:   reflect.TypeOf((*GroupsSectionsList)(nil)),
			Struct: "GroupsSectionsList",
			Field:  "ID",
		}
	}

	name, ok := alias[1].(string)
	if !ok {
		return &json.UnmarshalTypeError{
			Value:  string(data),
			Type:   reflect.TypeOf((*GroupsSectionsList)(nil)),
			Struct: "GroupsSectionsList",
			Field:  "Name",
		}
	}

	g.ID = int(id)
	g.Name = name

	return nil
}

// GroupsActionType for action_button in groups.
type GroupsActionType string

// GroupsActionType enums.
const (
	GroupsActionTypeOpenURL      GroupsActionType = "open_url"
	GroupsActionTypeSendEmail    GroupsActionType = "send_email"
	GroupsActionTypeCallPhone    GroupsActionType = "call_phone"
	GroupsActionTypeCallVK       GroupsActionType = "call_vk"
	GroupsActionTypeOpenGroupApp GroupsActionType = "open_group_app"
	GroupsActionTypeOpenApp      GroupsActionType = "open_app"
)

// GroupsActionButton struct.
type GroupsActionButton struct {
	ActionType GroupsActionType         `json:"action_type"`
	Target     GroupsActionButtonTarget `json:"target"`
	Title      string                   `json:"title"`

	// IsEnabled for GroupSettings
	IsEnabled Bool `json:"is_enabled,omitempty"`
}

// GroupsActionButtonTarget struct.
type GroupsActionButtonTarget struct {
	// ActionType == ActionTypeSendEmail
	Email string `json:"email"`

	// ActionType == ActionTypeCallPhone
	Phone string `json:"phone"`

	// ActionType == ActionTypeCallVK
	UserID int `json:"user_id"`

	// ActionType == ActionTypeOpenURL
	URL string `json:"url"`

	// ActionType == ActionTypeOpenApp
	GoogleStoreURL string `json:"google_store_url"`
	ItunesURL      string `json:"itunes_url"`
	// URL string `json:"url"`

	// ActionType == ActionTypeOpenGroupApp
	AppID int `json:"app_id"`

	IsInternal Bool `json:"is_internal"`
}

// GroupXtrInvitedBy struct.
type GroupXtrInvitedBy struct {
	AdminLevel   int    `json:"admin_level"`
	ID           int    `json:"id"`          // Community ID
	InvitedBy    int    `json:"invited_by"`  // Inviter ID
	Name         string `json:"name"`        // Community name
	Photo100     string `json:"photo_100"`   // URL of square photo of the community with 100 pixels in width
	Photo200     string `json:"photo_200"`   // URL of square photo of the community with 200 pixels in width
	Photo50      string `json:"photo_50"`    // URL of square photo of the community with 50 pixels in width
	ScreenName   string `json:"screen_name"` // Domain of the community page
	Type         string `json:"type"`
	IsClosed     int    `json:"is_closed"`     // Information whether community is closed
	IsAdmin      Bool   `json:"is_admin"`      // Information whether current user is manager
	IsMember     Bool   `json:"is_member"`     // Information whether current user is member
	IsAdvertiser Bool   `json:"is_advertiser"` // Information whether current user is advertiser
}

// ToMention return mention.
func (group *GroupXtrInvitedBy) ToMention() string {
	return fmt.Sprintf("[club%d|%s]", group.ID, group.Name)
}

// GroupsLinksItem struct.
type GroupsLinksItem struct {
	Desc      string `json:"desc"`       // Link description
	EditTitle Bool   `json:"edit_title"` // Information whether the link title can be edited
	ID        int    `json:"id"`         // Link ID
	Name      string `json:"name"`       // Link title
	Photo100  string `json:"photo_100"`  // URL of square image of the link with 100 pixels in width
	Photo50   string `json:"photo_50"`   // URL of square image of the link with 50 pixels in width
	URL       string `json:"url"`        // Link URL
}

// GroupsLongPollEvents struct.
type GroupsLongPollEvents struct {
	MessageNew                    Bool `json:"message_new"`
	MessageReply                  Bool `json:"message_reply"`
	PhotoNew                      Bool `json:"photo_new"`
	AudioNew                      Bool `json:"audio_new"`
	VideoNew                      Bool `json:"video_new"`
	WallReplyNew                  Bool `json:"wall_reply_new"`
	WallReplyEdit                 Bool `json:"wall_reply_edit"`
	WallReplyDelete               Bool `json:"wall_reply_delete"`
	WallReplyRestore              Bool `json:"wall_reply_restore"`
	WallPostNew                   Bool `json:"wall_post_new"`
	BoardPostNew                  Bool `json:"board_post_new"`
	BoardPostEdit                 Bool `json:"board_post_edit"`
	BoardPostRestore              Bool `json:"board_post_restore"`
	BoardPostDelete               Bool `json:"board_post_delete"`
	PhotoCommentNew               Bool `json:"photo_comment_new"`
	PhotoCommentEdit              Bool `json:"photo_comment_edit"`
	PhotoCommentDelete            Bool `json:"photo_comment_delete"`
	PhotoCommentRestore           Bool `json:"photo_comment_restore"`
	VideoCommentNew               Bool `json:"video_comment_new"`
	VideoCommentEdit              Bool `json:"video_comment_edit"`
	VideoCommentDelete            Bool `json:"video_comment_delete"`
	VideoCommentRestore           Bool `json:"video_comment_restore"`
	MarketCommentNew              Bool `json:"market_comment_new"`
	MarketCommentEdit             Bool `json:"market_comment_edit"`
	MarketCommentDelete           Bool `json:"market_comment_delete"`
	MarketCommentRestore          Bool `json:"market_comment_restore"`
	MarketOrderNew                Bool `json:"market_order_new"`
	MarketOrderEdit               Bool `json:"market_order_edit"`
	PollVoteNew                   Bool `json:"poll_vote_new"`
	GroupJoin                     Bool `json:"group_join"`
	GroupLeave                    Bool `json:"group_leave"`
	GroupChangeSettings           Bool `json:"group_change_settings"`
	GroupChangePhoto              Bool `json:"group_change_photo"`
	GroupOfficersEdit             Bool `json:"group_officers_edit"`
	MessageAllow                  Bool `json:"message_allow"`
	MessageDeny                   Bool `json:"message_deny"`
	WallRepost                    Bool `json:"wall_repost"`
	UserBlock                     Bool `json:"user_block"`
	UserUnblock                   Bool `json:"user_unblock"`
	MessageEdit                   Bool `json:"message_edit"`
	MessageTypingState            Bool `json:"message_typing_state"`
	LeadFormsNew                  Bool `json:"lead_forms_new"`
	LikeAdd                       Bool `json:"like_add"`
	LikeRemove                    Bool `json:"like_remove"`
	VkpayTransaction              Bool `json:"vkpay_transaction"`
	AppPayload                    Bool `json:"app_payload"`
	MessageRead                   Bool `json:"message_read"`
	MessageEvent                  Bool `json:"message_event"`
	DonutSubscriptionCreate       Bool `json:"donut_subscription_create"`
	DonutSubscriptionProlonged    Bool `json:"donut_subscription_prolonged"`
	DonutSubscriptionExpired      Bool `json:"donut_subscription_expired"`
	DonutSubscriptionCancelled    Bool `json:"donut_subscription_cancelled"`
	DonutSubscriptionPriceChanged Bool `json:"donut_subscription_price_changed"`
	DonutMoneyWithdraw            Bool `json:"donut_money_withdraw"`
	DonutMoneyWithdrawError       Bool `json:"donut_money_withdraw_error"`

	// Bugs
	// MessagesEdit  Bool `json:"messages_edit"`
	// WallNew       Bool `json:"wall_new"`
	// WallNewReply  Bool `json:"wall_new_reply"`
	// WallEditReply Bool `json:"wall_edit_reply"`
}

// GroupsLongPollServer struct.
type GroupsLongPollServer struct {
	Key    string `json:"key"`    // Long Poll key
	Server string `json:"server"` // Long Poll server address
	Ts     string `json:"ts"`     // Number of the last event
}

// GetURL return link.
func (lp GroupsLongPollServer) GetURL(wait int) string {
	return fmt.Sprintf("%s?act=a_check&key=%s&ts=%s&wait=%d", lp.Server, lp.Key, lp.Ts, wait)
}

// GroupsLongPollSettings struct.
type GroupsLongPollSettings struct {
	APIVersion string               `json:"api_version"` // API version used for the events
	Events     GroupsLongPollEvents `json:"events"`
	IsEnabled  Bool                 `json:"is_enabled"` // Shows whether Long Poll is enabled
}

// GroupsMarketType ...
type GroupsMarketType string

// Possible values.
const (
	GroupsMarketBasic    GroupsMarketType = "basic"
	GroupsMarketAdvanced GroupsMarketType = "advanced"
)

// GroupsMarketInfo struct.
type GroupsMarketInfo struct {
	// information about the type of store. Returned if the group includes
	// the "Products" section.
	Type            GroupsMarketType  `json:"type,omitempty"`
	ContactID       int               `json:"contact_id,omitempty"` // Contact person ID
	Currency        MarketCurrency    `json:"currency,omitempty"`
	CurrencyText    string            `json:"currency_text,omitempty"` // Currency name
	Enabled         Bool              `json:"enabled"`                 // Information whether the market is enabled
	CommentsEnabled Bool              `json:"comments_enabled,omitempty"`
	CanMessage      Bool              `json:"can_message,omitempty"`
	IsHsEnabled     Bool              `json:"is_hs_enabled,omitempty"`
	MainAlbumID     int               `json:"main_album_id,omitempty"` // Main market album ID
	PriceMax        string            `json:"price_max,omitempty"`     // Maximum price
	PriceMin        string            `json:"price_min,omitempty"`     // Minimum price
	Wiki            PagesWikipageFull `json:"wiki,omitempty"`
	CityIDs         []int             `json:"city_ids"`
	CountryIDs      []int             `json:"country_ids,omitempty"`
	MinOrderPrice   MarketPrice       `json:"min_order_price,omitempty"`
}

// GroupRole Role type.
const (
	GroupRoleModerator     = "moderator"
	GroupRoleEditor        = "editor"
	GroupRoleAdministrator = "administrator"
	GroupRoleCreator       = "creator"
)

// GroupsMemberRole struct.
type GroupsMemberRole struct {
	ID          int      `json:"id"` // User ID
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
}

// GroupsMemberRoleXtrUser struct.
type GroupsMemberRoleXtrUser struct {
	User
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
}

// GroupsMemberStatus struct.
type GroupsMemberStatus struct {
	Member      Bool     `json:"member"`  // Information whether user is a member of the group
	UserID      int      `json:"user_id"` // User ID
	Permissions []string `json:"permissions"`
}

// GroupsMemberStatusFull struct.
type GroupsMemberStatusFull struct {
	Invitation Bool `json:"invitation"` // Information whether user has been invited to the group
	Member     Bool `json:"member"`     // Information whether user is a member of the group
	Request    Bool `json:"request"`    // Information whether user has send request to the group
	CanInvite  Bool `json:"can_invite"` // Information whether user can be invite
	CanRecall  Bool `json:"can_recall"` // Information whether user's invite to the group can be recalled
	UserID     int  `json:"user_id"`    // User ID
}

// GroupsOnlineStatus Status type.
const (
	GroupsOnlineStatusTypeNone       = "none"
	GroupsOnlineStatusTypeOnline     = "online"
	GroupsOnlineStatusTypeAnswerMark = "answer_mark"
)

// GroupsOnlineStatus struct.
type GroupsOnlineStatus struct {
	Minutes int    `json:"minutes"` // Estimated time of answer (for status = answer_mark)
	Status  string `json:"status"`
}

// GroupsOwnerXtrBanInfo struct.
type GroupsOwnerXtrBanInfo struct {
	BanInfo GroupsBanInfo `json:"ban_info"`
	Group   Group         `json:"group"`
	Profile User          `json:"profile"`
	Type    string        `json:"type"`
}

// GroupsSubjectItem struct.
type GroupsSubjectItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GroupsTokenPermissionSetting struct.
type GroupsTokenPermissionSetting struct {
	Name    string `json:"name"`
	Setting int    `json:"setting"`
}

// GroupsTokenPermissions struct.
type GroupsTokenPermissions struct {
	Mask        int                            `json:"mask"`
	Permissions []GroupsTokenPermissionSetting `json:"permissions"`
}

// GroupsTag struct.
type GroupsTag struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}
