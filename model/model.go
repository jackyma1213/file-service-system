package model

import (
	"errors"
	"fmt"
)

type Node struct {
	FileId           int
	Name             string
	ObjectType       int
	ParentFileId     int
	LastModifiedDate string
	Children         *[]Node
}

type Tree struct {
	Root *Node
}

func New() *Tree {
	return &Tree{
		&Node{
			FileId:     0,
			Name:       "root",
			ObjectType: 1,
			Children:   &[]Node{},
		},
	}
}

func (tree *Tree) Find(fileId int, node *Node) *Node {

	if node.FileId == fileId {
		return node
	} else {
		var returningNode *Node
		for _, child := range *node.Children {
			if returningNode = tree.Find(fileId, &child); returningNode != nil {
				return returningNode
			}
		}
	}

	return nil
}

func (tree *Tree) Add(node Node, parentFileId int) {
	parentNode := tree.Find(parentFileId, tree.Root)

	if parentNode != nil {
		*parentNode.Children = append(*parentNode.Children, node)
	}

}

func (tree *Tree) Remove(fileId int) (int, error) {
	node := tree.Find(fileId, tree.Root)

	if node != nil {
		count := childCount(0, *node)
		parentNode := tree.Find(node.ParentFileId, tree.Root)
		fmt.Println("parentNode", parentNode)
		indexOfNode := indexOf(fileId, *parentNode.Children)
		*parentNode.Children = removeElement(parentNode.Children, indexOfNode)

		return count, nil
	} else {
		return -1, errors.New("not found")
	}
}

func childCount(count int, node Node) int {
	children := node.Children
	for _, child := range *children {
		count = childCount(count, child)
	}
	return count + 1
}

func indexOf(targetFileId int, list []Node) int {
	for k, v := range list {
		fmt.Println("k", k, v, targetFileId)
		if targetFileId == v.FileId {
			return k
		}
	}
	return -1 //not found.
}

func removeElement(list *[]Node, index int) []Node {
	tempList := *list
	tempList[index] = tempList[len(tempList)-1]

	return tempList[:len(tempList)-1]
}
