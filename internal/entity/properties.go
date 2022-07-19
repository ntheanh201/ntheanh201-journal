package entity

type SelectOptions struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Color Color  `json:"color,omitempty"`
}

type (
	FormulaMetadata struct {
		Expression string `json:"expression"`
	}
)

type Properties map[string]Property

type (
	DatabasePropertyType string
	NumberFormat         string
	FormulaResultType    string
	RollupResultType     string
	SortTimestamp        string
	SortDirection        string
)

type Property struct {
	ID          string               `json:"id,omitempty"`
	Type        DatabasePropertyType `json:"type"`
	Name        string               `json:"name,omitempty"`
	Title       interface{}          `json:"title,omitempty"`
	Date        *Date                `json:"date,omitempty"`
	RichText    *[]RichText          `json:"rich_text,omitempty"`
	Select      *SelectOptions       `json:"select,omitempty"`
	MultiSelect *[]SelectOptions     `json:"multi_select,omitempty"`
	Formula     *interface{}         `json:"formula,omitempty"`
	Relation    *interface{}         `json:"relation,omitempty"`
}
