package entity

type SelectOptions struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Color Color  `json:"color,omitempty"`
}

type (
	EmptyMetadata struct {
	}
	NumberMetadata struct {
		Format NumberFormat `json:"format"`
	}
	SelectMetadata struct {
		Options []SelectOptions `json:"options"`
	}
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
	ID             string               `json:"id,omitempty"`
	Type           DatabasePropertyType `json:"type"`
	Name           string               `json:"name,omitempty"`
	Title          *EmptyMetadata       `json:"title,omitempty"`
	RichText       *EmptyMetadata       `json:"rich_text,omitempty"`
	Date           *EmptyMetadata       `json:"date,omitempty"`
	People         *EmptyMetadata       `json:"people,omitempty"`
	Files          *EmptyMetadata       `json:"files,omitempty"`
	Checkbox       *EmptyMetadata       `json:"checkbox,omitempty"`
	URL            *EmptyMetadata       `json:"url,omitempty"`
	Email          *EmptyMetadata       `json:"email,omitempty"`
	PhoneNumber    *EmptyMetadata       `json:"phone_number,omitempty"`
	CreatedTime    *EmptyMetadata       `json:"created_time,omitempty"`
	CreatedBy      *EmptyMetadata       `json:"created_by,omitempty"`
	LastEditedTime *EmptyMetadata       `json:"last_edited_time,omitempty"`
	LastEditedBy   *EmptyMetadata       `json:"last_edited_by,omitempty"`

	Number      *NumberMetadata `json:"number,omitempty"`
	Select      *SelectMetadata `json:"select,omitempty"`
	MultiSelect *SelectMetadata `json:"multi_select,omitempty"`
}
