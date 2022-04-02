package model

type Item struct {
	FileId           int    `json:"fileId"`
	Name             string `json:"name"`
	ObjectType       int    `json:"objectType"`
	ParentFieldId    int    `json:"parentField"`
	LastModifiedDate string `json:"lastModifiedDate"`
}

type Db struct {
	Items []Item
}

func New() *Db {
	return &Db{}
}

func (db *Db) Add(item Item) {
	db.Items = append(db.Items, item)
}

func (db *Db) Delete(fileId int) {

}
