package object

type Poll struct {
	AnswerID      int             `json:"answer_id"` // Current user's answer ID
	Answers       []PollsAnswer   `json:"answers"`
	Created       int             `json:"created"`  // Date when poll has been created in Unixtime
	ID            int             `json:"id"`       // Poll ID
	OwnerID       int             `json:"owner_id"` // Poll owner's ID
	Question      string          `json:"question"` // Poll question
	Votes         int             `json:"votes"`    // Votes number
	AnswerIDs     []int           `json:"answer_ids"`
	EndDate       int             `json:"end_date"`
	Anonymous     Bool     `json:"anonymous"` // Information whether the pole is anonymous
	Closed        Bool     `json:"closed"`
	IsBoard       Bool     `json:"is_board"`
	CanEdit       Bool     `json:"can_edit"`
	CanVote       Bool     `json:"can_vote"`
	CanReport     Bool     `json:"can_report"`
	CanShare      Bool     `json:"can_share"`
	Multiple      Bool     `json:"multiple"`
	DisableUnvote Bool     `json:"disable_unvote"`
	Photo         Photo     `json:"photo"`
	AuthorID      int             `json:"author_id"`
	Background    PollsBackground `json:"background"`
	Friends       []PollsFriend   `json:"friends"`
	Profiles      []User     `json:"profiles"`
	Groups        []Group   `json:"groups"`
	EmbedHash     string          `json:"embed_hash"`
}

type PollsFriend struct {
	ID int `json:"id"`
}

type PollsVoters struct {
	AnswerID int              `json:"answer_id"` // Answer ID
	Users    PollsVotersUsers `json:"users"`
}

type PollsAnswer struct {
	ID    int     `json:"id"`
	Rate  float64 `json:"rate"`
	Text  string  `json:"text"`
	Votes int     `json:"votes"`
}

type PollsVotersUsers struct {
	Count int   `json:"count"` // Votes number
	Items []int `json:"items"`
}

type PollsBackground struct {
	Type   string `json:"type"`
	Angle  int    `json:"angle"`
	Color  string `json:"color"`
	Points []struct {
		Position float64 `json:"position"`
		Color    string  `json:"color"`
	} `json:"points"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

