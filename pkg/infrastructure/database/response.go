package database

type DatabaseResponse struct {
	Object         string        `json:"object"`
	Id             string        `json:"id"`
	CreatedTime    string        `json:"created_time"`
	LastEditedTime string        `json:"last_edited_time"`
	Icon           Icon          `json:"icon"`
	Cover          Cover         `json:"cover"`
	Url            string        `json:"url"`
	Title          []Title       `json:"title"`
	Description    []Description `json:"description"`
	Properties     Properties    `json:"properties"`
	Parent         Parent        `json:"parent"`
	Archived       bool          `json:"archived"`
	IsInline       bool          `json:"is_inline"`
	PublicUrl      string        `json:"public_url"`
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

type Properties struct {
	PlusOne           PlusOne           `json:"+1"`
	InStock           InStock           `json:"in_stock"`
	Price             Price             `json:"price"`
	Description       Desc              `json:"description"`
	LastOrdered       LastOrdered       `json:"last_ordered"`
	Meals             Meals             `json:"meals"`
	NumberOfMeals     NumberOfMeals     `json:"number_of_meals"`
	StoreAvailability StoreAvailability `json:"store_availability"`
	Photo             Photo             `json:"photo"`
	FoodGroup         FoodGroup         `json:"food_group"`
}

type PlusOne struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	People string `json:"people"`
}

type InStock struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	CheckBox string `json:"checkbox"`
}

type Price struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Number Number `json:"number"`
}

type Number struct {
	Format string `json:"format"`
}

type Desc struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	RichText string `json:"rich_text"`
}

type LastOrdered struct {
	Id   string      `json:"id"`
	Name string      `json:"name"`
	Type string      `json:"type"`
	Date interface{} `json:"date"`
}

type Meals struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	Relation Relation `json:"relation"`
}

type Relation struct {
	DatabaseId         string `json:"database_id"`
	SyncedPropertyName string `json:"synced_property_name"`
}

type NumberOfMeals struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Rollup Rollup `json:"rollup"`
}
type Rollup struct {
	RollupPropertyName   string `json:"rollup_property_name"`
	RelationPropertyName string `json:"relation_property_name"`
	RollupPropertyId     string `json:"rollup_property_id"`
	RelationPropertyId   string `json:"relation_property_id"`
	Function             string `json:"function"`
}

type StoreAvailability struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	MultiSelect MultiSelect `json:"multi_select"`
}

type MultiSelect struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Photo struct {
	Id    string      `json:"id"`
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Files interface{} `json:"files"`
}

type FoodGroup struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Select Select `json:"select"`
}

type Select struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Parent struct {
	Type   string `json:"type"`
	PageId string `json:"pageId"`
}
