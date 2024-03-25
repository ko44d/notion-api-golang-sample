package notionapi

type DatabaseResponse struct {
	Object         string        `json:"object"`
	Id             string        `json:"id"`
	CreatedTime    string        `json:"created_time"`
	CreatedBy      CreatedBy     `json:"created_by"`
	LastEditedTime string        `json:"last_edited_time"`
	Icon           Icon          `json:"icon"`
	Cover          Cover         `json:"cover"`
	Url            string        `json:"url"`
	Title          []Title       `json:"title"`
	Description    []Description `json:"description"`
	Properties     interface{}   `json:"properties"`
	Parent         Parent        `json:"parent"`
	Archived       bool          `json:"archived"`
	IsInline       bool          `json:"is_inline"`
	PublicUrl      string        `json:"public_url"`
}

type CreatedBy struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}
type Icon struct {
	Type  string `json:"type"`
	Emoji string `json:"emoji"`
}

type Cover struct {
	Type     string   `json:"type"`
	External External `json:"external"`
}

type External struct {
	Url string `json:"url"`
}

type Title struct {
	Type        string      `json:"type"`
	Text        Text        `json:"text"`
	Annotations Annotations `json:"annotations"`
	PlainText   string      `json:"plain_text"`
	Href        string      `json:"href"`
}

type Text struct {
	Content string `json:"content"`
	Link    string `json:"link"`
}

type Annotations struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}

type Description struct {
	Type        string      `json:"type"`
	Text        Text        `json:"text"`
	Annotations Annotations `json:"annotations"`
	PlainText   string      `json:"plain_text"`
	Href        string      `json:"href"`
}

type Parent struct {
	Type   string `json:"type"`
	PageId string `json:"page_id"`
}
