package entity

import "ntheanh201-journal/internal/request"

type FormulaArg struct {
	Name       *string `json:"name,omitempty"`
	ResultType string  `json:"result_type"`
	Type       string  `json:"type"`
	Value      *string `json:"value,omitempty"`
	ValueType  *string `json:"value_type,omitempty"`
}

type ColumnFormula struct {
	Args       []FormulaArg `json:"args"`
	Name       string       `json:"name"`
	Operator   string       `json:"operator"`
	ResultType string       `json:"result_type"`
	Type       string       `json:"type"`
}

type TextAttr = []string

// TextSpan describes a text with attributes
type TextSpan struct {
	Text  string     `json:"Text"`
	Attrs []TextAttr `json:"Attrs"`
}

// CollectionColumnOption describes options for ColumnTypeMultiSelect
// collection column
type CollectionColumnOption struct {
	Color string `json:"color"`
	ID    string `json:"id"`
	Value string `json:"value"`
}

// ColumnSchema describes a info of a collection column
type ColumnSchema struct {
	Name string `json:"name"`
	// ColumnTypeTitle etc.
	Type string `json:"type"`

	// for Type == ColumnTypeNumber, e.g. "dollar", "number"
	NumberFormat string `json:"number_format"`

	// For Type == ColumnTypeRollup
	Aggregation        string `json:"aggregation"` // e.g. "unique"
	TargetProperty     string `json:"target_property"`
	RelationProperty   string `json:"relation_property"`
	TargetPropertyType string `json:"target_property_type"`

	// for Type == ColumnTypeRelation
	CollectionID string `json:"collection_id"`
	Property     string `json:"property"`

	// for Type == ColumnTypeFormula
	Formula *ColumnFormula

	Options []*CollectionColumnOption `json:"options"`

	// TODO: would have to set it up from Collection.RawJSON
	//RawJSON map[string]interface{} `json:"-"`
}

// CollectionPageProperty describes properties of a collection
type CollectionPageProperty struct {
	Property string `json:"property"`
	Visible  bool   `json:"visible"`
}

// CollectionFormat describes format of a collection
type CollectionFormat struct {
	CoverPosition  float64                   `json:"collection_cover_position"`
	PageProperties []*CollectionPageProperty `json:"collection_page_properties"`
}

// Collection describes a collection
type Collection struct {
	ID          string                   `json:"id"`
	Version     int                      `json:"version"`
	Name        interface{}              `json:"name"`
	Schema      map[string]*ColumnSchema `json:"schema"`
	Format      *CollectionFormat        `json:"format"`
	ParentID    string                   `json:"parent_id"`
	ParentTable string                   `json:"parent_table"`
	Alive       bool                     `json:"alive"`
	CopiedFrom  string                   `json:"copied_from"`
	Cover       string                   `json:"cover"`
	Description []interface{}            `json:"description"`

	// TODO: are those ever present?
	Type          string   `json:"type"`
	FileIDs       []string `json:"file_ids"`
	Icon          string   `json:"icon"`
	TemplatePages []string `json:"template_pages"`

	// calculated by us
	name    []*TextSpan
	RawJSON map[string]interface{} `json:"-"`
}

// TableProperty describes property of a table
type TableProperty struct {
	Width    int    `json:"width"`
	Visible  bool   `json:"visible"`
	Property string `json:"property"`
}

// FormatTable describes format for BlockTable
type FormatTable struct {
	PageSort        []string         `json:"page_sort"`
	TableWrap       bool             `json:"table_wrap"`
	TableProperties []*TableProperty `json:"table_properties"`
}

// CollectionView represents a collection view
type CollectionView struct {
	ID          string         `json:"id"`
	Version     int64          `json:"version"`
	Type        string         `json:"type"` // "table"
	Format      *FormatTable   `json:"format"`
	Name        string         `json:"name"`
	ParentID    string         `json:"parent_id"`
	ParentTable string         `json:"parent_table"`
	Query       *request.Query `json:"query2"`
	Alive       bool           `json:"alive"`
	PageSort    []string       `json:"page_sort"`
	SpaceID     string         `json:"space_id"`

	// set by us
	RawJSON map[string]interface{} `json:"-"`
}
