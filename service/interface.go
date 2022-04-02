package service

var ResonseStatusMessage = map[int]string{
	0:  "success",
	-1: "File ID not found",
	-2: "Invalid name",
}

type CreateRequestBody struct {
	Name          string `json:"name"`
	ObjectType    int    `json:"objectType"`
	ParentFieldId int    `json:"parentFieldId"`
}

type CreateResponse struct {
	FileId           int    `json:"fileId"`
	LastModifiedDate string `json:"lastModifiedDate"`
	Status           int    `json:"status"`
	Message          string `json:"message"`
}
