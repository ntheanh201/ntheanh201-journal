package entity

type FileBlock struct {
	Caption  []RichText  `json:"caption,omitempty"`
	Type     FileType    `json:"type"`
	File     *FileObject `json:"file,omitempty"`
	External *FileObject `json:"external,omitempty"`
}
