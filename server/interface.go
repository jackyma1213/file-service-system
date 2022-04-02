package server

type CreateRequestBody struct {
	Name          string `json: "name"`
	ObjectType    int    `json: "objectType"`
	ParentFieldId int    `json: "parentFieldId"`
}

type CreateResponse struct {
	FileId           int    `json: "fileId"`
	LastModifiedDate string `json: "lastmodifiedDate"`
	Status           int    `json: "status"`
	Message          string `json: "message"`
}
