package entity

import "encoding/json"

type TableRow struct {
	// TableView that owns this row
	TableView *TableView

	// data for row is stored as properties of a page
	Page *Block

	// values extracted from Page for each column
	Columns [][]*TextSpan
}

// ColumnInfo describes a schema for a given cell (column)
type ColumnInfo struct {
	// TableView that owns this column
	TableView *TableView

	// so that we can access TableRow.Columns[Index]
	Index    int
	Schema   *ColumnSchema
	Property *TableProperty
}

// TableView represents a view of a table (Notion calls it a Collection View)
// Meant to be a representation that is easier to work with
type TableView struct {
	// original data
	Page           *Page
	CollectionView *CollectionView
	Collection     *Collection

	// easier to work representation we calculate
	Columns []*ColumnInfo
	Rows    []*TableRow
}

// Permission represents user permissions o
type Permission struct {
	Type string `json:"type"`

	// common to some permission types
	Role interface{} `json:"role"`

	// if Type == "user_permission"
	UserID *string `json:"user_id,omitempty"`

	AddedTimestamp int64 `json:"added_timestamp"`

	// if Type == "public_permission"
	AllowDuplicate            bool `json:"allow_duplicate"`
	AllowSearchEngineIndexing bool `json:"allow_search_engine_indexing"`
}

type NotionID struct {
	DashID   string
	NoDashID string
}

// BlockV3 describes a block
type BlockV3 struct {
	// values that come from JSON
	// a unique ID of the block
	ID string `json:"id"`
	// if false, the page is deleted
	Alive bool `json:"alive"`
	// List of block ids for that make up content of this block
	// Use Content to get corresponding block (they are in the same order)
	ContentIDs   []string `json:"content,omitempty"`
	CopiedFrom   string   `json:"copied_from,omitempty"`
	CollectionID string   `json:"collection_id,omitempty"` // for BlockCollectionView
	// ID of the user who created this block
	CreatedBy   string `json:"created_by"`
	CreatedTime int64  `json:"created_time"`

	CreatedByTable string `json:"created_by_table"` // e.g. "notion_user"
	CreatedByID    string `json:"created_by_id"`    // e.g. "bb760e2d-d679-4b64-b2a9-03005b21870a",
	// ID of the user who last edited this block
	LastEditedBy      string `json:"last_edited_by"`
	LastEditedTime    int64  `json:"last_edited_time"`
	LastEditedByTable string `json:"last_edited_by_table"` // e.g. "notion_user"
	LastEditedByID    string `json:"last_edited_by_id"`    // e.g. "bb760e2d-d679-4b64-b2a9-03005b21870a"

	// List of block ids with discussion content
	DiscussionIDs []string `json:"discussion,omitempty"`
	// those ids seem to map to storage in s3
	// https://s3-us-west-2.amazonaws.com/secure.notion-static.com/${id}/${name}
	FileIDs []string `json:"file_ids,omitempty"`

	// TODO: don't know what this means
	IgnoreBlockCount bool `json:"ignore_block_count,omitempty"`

	// ID of parent Block
	ParentID    string `json:"parent_id"`
	ParentTable string `json:"parent_table"`
	// not always available
	Permissions *[]Permission          `json:"permissions,omitempty"`
	Properties  map[string]interface{} `json:"properties,omitempty"`
	SpaceID     string                 `json:"space_id"`
	// type of the block e.g. TypeText, TypePage etc.
	Type string `json:"type"`
	// blocks are versioned
	Version int64 `json:"version"`
	// for BlockCollectionView
	ViewIDs []string `json:"view_ids,omitempty"`

	// Parent of this block
	Parent *Block `json:"-"`

	// maps ContentIDs array to Block type
	Content []*Block `json:"-"`
	// this is for some types like TypePage, TypeText, TypeHeader etc.
	InlineContent []*TextSpan `json:"-"`

	// for BlockPage
	Title string `json:"-"`

	// For BlockTodo, a checked state
	IsChecked bool `json:"-"`

	// for BlockBookmark
	Description string `json:"-"`
	Link        string `json:"-"`

	// for BlockBookmark it's the url of the page
	// for BlockGist it's the url for the gist
	// for BlockImage it's url of the image. Sometimes you need to use DownloadFile()
	//   to get this image
	// for BlockFile it's url of the file
	// for BlockEmbed it's url of the embed
	Source string `json:"-"`

	// for BlockFile
	FileSize string `json:"-"`

	// for BlockCode
	Code         string `json:"-"`
	CodeLanguage string `json:"-"`

	// for BlockCollectionView. There can be multiple views
	// those correspond to ViewIDs
	TableViews []*TableView `json:"-"`

	Page *Page `json:"-"`

	// RawJSON represents Block as
	RawJSON map[string]interface{} `json:"-"`

	notionID       *NotionID
	parentNotionID *NotionID
	isResolved     bool
}

type Record struct {
	// fields returned by the server
	Role string `json:"role"`
	// polymorphic value of the record, which we decode into Block, Space etc.
	Value json.RawMessage `json:"value"`

	// fields calculated from Value based on type
	ID       string    `json:"-"`
	Table    string    `json:"-"`
	Activity *Activity `json:"-"`
	Block    *BlockV3  `json:"-"`
	//Space          *Space          `json:"-"`
	//NotionUser     *NotionUser     `json:"-"`
	//UserRoot       *UserRoot       `json:"-"`
	//UserSettings   *UserSettings   `json:"-"`
	Collection     *Collection     `json:"-"`
	CollectionView *CollectionView `json:"-"`
	Comment        *Comment        `json:"-"`
	//Discussion     *Discussion     `json:"-"`
	// TODO: add more types
}

type RecordMap struct {
	Version         int                `json:"__version__"`
	Activities      map[string]*Record `json:"activity"`
	Blocks          map[string]*Record `json:"block"`
	Spaces          map[string]*Record `json:"space"`
	NotionUsers     map[string]*Record `json:"notion_user"`
	UsersRoot       map[string]*Record `json:"user_root"`
	UserSettings    map[string]*Record `json:"user_setting"`
	Collections     map[string]*Record `json:"collection"`
	CollectionViews map[string]*Record `json:"collection_view"`
	Comments        map[string]*Record `json:"comment"`
	Discussions     map[string]*Record `json:"discussion"`
}

// TODO: reference: https://github.dev/kjk/notionapi/blob/master/api_loadCachedPageChunk.go
