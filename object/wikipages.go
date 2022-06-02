package object

type PagesWikipageFull struct {
	Created                  int    `json:"created"`
	CreatorID                int    `json:"creator_id"`
	CurrentUserCanEdit       Bool   `json:"current_user_can_edit"`
	CurrentUserCanEditAccess Bool   `json:"current_user_can_edit_access"`
	Edited                   int    `json:"edited"`
	EditorID                 int    `json:"editor_id"`
	PageID                   int    `json:"page_id"`
	GroupID                  int    `json:"group_id"`
	HTML                     string `json:"html"`
	ID                       int    `json:"id"`
	Source                   string `json:"source"`
	Title                    string `json:"title"`
	ViewURL                  string `json:"view_url"`
	Views                    int    `json:"views"`
	WhoCanEdit               int    `json:"who_can_edit"`
	WhoCanView               int    `json:"who_can_view"`
	VersionCreated           int    `json:"version_created"`
}
