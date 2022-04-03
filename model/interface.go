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
