package model

type Node struct {
	FileId           int
	Name             *string
	ObjectType       int
	LastModifiedDate string
	ParentFileId     int
	Children         *[]Node
	Content          *string
}

type Tree struct {
	Root *Node
}

type FileObject struct {
	FileId           int          `json:"fileId"`
	ObjectType       int          `json:"objectType"`
	Name             string       `json:"name"`
	LastModifiedDate string       `json:"lastModifiedDate"`
	Children         []FileObject `json:"children"`
}
