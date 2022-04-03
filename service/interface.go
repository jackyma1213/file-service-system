package service

var ResonseStatusMessage = map[int]string{
	0:  "success",
	-1: "File ID not found",
	-2: "Invalid name",
}

type ReponseStatus struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type CreateRequestBody struct {
	Name         string `json:"name" validate:"required"`
	ObjectType   *int   `json:"objectType" validate:"required,min=1,max=2"`
	ParentFileId *int   `json:"parentFileId" validate:"required"`
}

type CreateResponse struct {
	ReponseStatus
	FileId           int    `json:"fileId"`
	LastModifiedDate string `json:"lastModifiedDate"`
}

type DeleteResponse struct {
	ReponseStatus
	FileId int `json:"fileId"`
	Count  int `json:"count"`
}

type GetChildrenResponse struct {
	ReponseStatus
	FileList []FileObject `json:"fileList"`
}

type FileObject struct {
	FileId           int          `json:"fileId"`
	ObjectType       int          `json:"objectType"`
	Name             string       `json:"name"`
	LastModifiedDate string       `json:"lastModifiedDate"`
	Children         []FileObject `json:"children"`
}

type UpdateRequestBody struct {
	Content *string `json:"content"`
	Name    *string `json:"name"`
}

type UpdateResponse struct {
	ReponseStatus
	FileId           int    `json:"fileId"`
	LastModifiedDate string `json:"lastModifiedDate"`
}
