package model

import (
	"errors"
)

type Item struct {
	FileId           int    `json:"fileId"`
	Name             string `json:"name"`
	ObjectType       int    `json:"objectType"`
	ParentFileId     int    `json:"parentFileId"`
	LastModifiedDate string `json:"lastModifiedDate"`
}

type Db struct {
	Items []Item
	Directory
}

func New() *Db {
	return &Db{}
}

func (db *Db) Add(item Item) {
	db.Items = append(db.Items, item)
}

func (db *Db) Delete(fileId int) error {

	index, _ := binarySearch(db.Items, fileId)

	db.Items = append(db.Items[:index], db.Items[index+1:]...)

	if index == -1 {
		return errors.New("not found")
	}

	return nil

}

func (db *Db) Get(fileId int) (Item, error) {

	index, _ := binarySearch(db.Items, fileId)

	if index == -1 {
		return Item{}, errors.New("not found")
	}

	return db.Items[index], nil
}

func binarySearch(a []Item, search int) (result int, searchCount int) {
	if len(a) == 0 {
		return -1, nil
	}

	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1 // not found
	case a[mid].FileId > search:
		result, searchCount = binarySearch(a[:mid], search)
	case a[mid].FileId < search:
		result, searchCount = binarySearch(a[mid+1:], search)
		if result >= 0 { // if anything but the -1 "not found" result
			result += mid + 1
		}
	default: // a[mid] == search
		result = mid // found
	}
	searchCount++
	return
}
