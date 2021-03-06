package model

import (
	"errors"
	"time"
)

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

func (tree *Tree) Add(node Node, parentFileId int) error {
	parentNode := tree.Find(parentFileId, tree.Root)

	if parentNode != nil {
		*parentNode.Children = append(*parentNode.Children, node)
		return nil
	} else {
		return errors.New("not found")
	}
}

func (tree *Tree) Remove(fileId int) (int, error) {
	node := tree.Find(fileId, tree.Root)

	if node != nil {
		count := childCount(0, *node)
		parentNode := tree.Find(node.ParentFileId, tree.Root)
		indexOfNode := indexOf(fileId, *parentNode.Children)
		*parentNode.Children = removeElement(parentNode.Children, indexOfNode)

		return count, nil
	} else {
		return -1, errors.New("not found")
	}
}

func (tree *Tree) GetChildren(fileId int) ([]FileObject, error) {
	node := tree.Find(fileId, tree.Root)

	if node != nil {

		if node.ObjectType == 2 {
			fileList := []FileObject{
				{
					FileId:           node.FileId,
					ObjectType:       node.ObjectType,
					Name:             *node.Name,
					LastModifiedDate: node.LastModifiedDate,
				},
			}
			return fileList, nil
		} else {
			var fileList = []FileObject{}
			fileList = getChildren(node.Children)
			return fileList, nil
		}

	} else {
		return nil, errors.New("not found")
	}
}

func (tree *Tree) Update(fileId int, content *string, name *string) (*Node, error) {
	node := tree.Find(fileId, tree.Root)

	if node != nil {

		if node.ObjectType == 1 && content != nil {
			return nil, errors.New("not file")

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

func (tree *Tree) GetFileContent(fileId int) (*Node, error) {

	node := tree.Find(fileId, tree.Root)

	if node != nil {

		if node.ObjectType == 1 {
			return nil, errors.New("not file")

		} else {

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

func getChildren(children *[]Node) (fileList []FileObject) {
	for _, child := range *children {
		fileList = append(fileList, FileObject{
			FileId:           child.FileId,
			ObjectType:       child.ObjectType,
			Name:             *child.Name,
			LastModifiedDate: child.LastModifiedDate,
			Children:         getChildren(child.Children),
		})
	}
	return fileList
}
