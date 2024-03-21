package domain

type Database interface {
	GetDatabase()
}

type DatabaseParam struct {
	DatabaseId string
}

type DatabaseResponse struct {
	Object         string
	Id             string
	CreatedTime    string
	LastEditedTime string
	Icon           Icon
	Cover          Cover
	Url            string
	Title          []Title
	Description    []Description
	Properties     Properties
	Parent         Parent
	Archived       bool
	IsInline       bool
	PublicUrl      string
}

type Icon struct {
	Type  string
	Emoji string
}

type Cover struct {
	Type     string
	External External
}

type External struct {
	Url string
}

type Title struct {
	Type        string
	Text        Text
	Annotations Annotations
	PlainText   string
	Href        string
}

type Text struct {
	Content string
	Link    string
}

type Annotations struct {
	Bold          bool
	Italic        bool
	Strikethrough bool
	Underline     bool
	Code          bool
	Color         string
}

type Description struct {
	Type         string
	Text         Text
	Annotiations Annotations
	PlainText    string
	Href         string
}

type Properties struct {
	PlusOne           PlusOne
	InStock           InStock
	Price             Price
	Description       Desc
	LastOrdered       LastOrdered
	Meals             Meals
	NumberOfMeals     NumberOfMeals
	StoreAvailability StoreAvailability
	Photo             Photo
	FoodGroup         FoodGroup
}

type PlusOne struct {
	Id     string
	Name   string
	Type   string
	People string
}

type InStock struct {
	Id       string
	Name     string
	Type     string
	CheckBox string
}

type Price struct {
	Id     string
	Name   string
	Type   string
	Number Number
}

type Number struct {
	Format string
}

type Desc struct {
	Id       string
	Name     string
	Type     string
	RichText string
}

type LastOrdered struct {
	Id   string
	Name string
	Type string
	Date interface{}
}

type Meals struct {
	Id       string
	Name     string
	Type     string
	Relation Relation
}

type Relation struct {
	DatabaseId         string
	SyncedPropertyName string
}

type NumberOfMeals struct {
	Id     string
	Name   string
	Type   string
	Rollup Rollup
}
type Rollup struct {
	RollupPropertyName   string
	RelationPropertyName string
	RollupPropertyId     string
	RelationPropertyId   string
	Function             string
}

type StoreAvailability struct {
	Id          string
	Name        string
	Type        string
	MultiSelect MultiSelect
}

type MultiSelect struct {
	Id    string
	Name  string
	Color string
}

type Photo struct {
	Id    string
	Name  string
	Type  string
	Files interface{}
}

type FoodGroup struct {
	Id     string
	Name   string
	Type   string
	Select Select
}

type Select struct {
	Id    string
	Name  string
	Color string
}

type Parent struct {
	Type   string
	PageId string
}
