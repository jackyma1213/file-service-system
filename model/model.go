package model

import (
	"errors"
	"fmt"
	"time"
)

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

func New() *Tree {
	return &Tree{
		&Node{
			FileId:     0,
			Name:       nil,
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

func (tree *Tree) Update(fileId int, content *string, name *string) (*Node, error) {
	node := tree.Find(fileId, tree.Root)

	temp := *tree.Root.Children

	fmt.Println("Update", node, temp[0])

	if node != nil {

		if node.ObjectType == 1 && content != nil {
			return nil, errors.New("not found")

		} else {

			if content != nil {

				*node.Content = *content
			}

			if name != nil {
				*node.Name = *name
			}
			node.LastModifiedDate = time.Now().Format(time.RFC3339)

			return node, nil
		}

	} else {
		return nil, nil
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
