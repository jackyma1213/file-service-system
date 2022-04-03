// package main

// import (
// 	"file-service/model"
// 	"fmt"
// )

// func main() {
// 	var tree = model.New()

// 	tree.Add(model.Node{
// 		FileId:           1,
// 		Name:             "node1",
// 		ObjectType:       1,
// 		ParentFileId:     0,
// 		LastModifiedDate: "123sd",
// 		Children:         &[]model.Node{},
// 	}, 0)
// 	fmt.Println(*tree.Root, *tree.Root.Children)

// 	tree.Add(model.Node{
// 		FileId:           2,
// 		Name:             "node2",
// 		ObjectType:       1,
// 		ParentFileId:     1,
// 		LastModifiedDate: "123sd",
// 		Children:         &[]model.Node{},
// 	}, 1)
// 	temp := *tree.Root.Children
// 	fmt.Println(*tree.Root, *tree.Root.Children, *temp[0].Children)

// 	tree.Add(model.Node{
// 		FileId:           3,
// 		Name:             "node3",
// 		ObjectType:       1,
// 		ParentFileId:     2,
// 		LastModifiedDate: "123sd",
// 		Children:         &[]model.Node{},
// 	}, 2)
// 	tree.Add(model.Node{
// 		FileId:           4,
// 		Name:             "node4",
// 		ObjectType:       1,
// 		ParentFileId:     1,
// 		LastModifiedDate: "123sd",
// 		Children:         &[]model.Node{},
// 	}, 1)

// 	fmt.Println(*tree.Root, *tree.Root.Children, *temp[0].Children)
// 	count, _ := tree.Remove(1)
// 	fmt.Println("count", count)

// 	fmt.Println(*tree.Root, *tree.Root.Children)

// }
