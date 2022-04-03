package service

import (
	"encoding/json"
	"file-service/model"
	"net/http"
	"strconv"
)

func FileSystemGetChildrenService(tree *model.Tree) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fileId := r.URL.Query().Get("fileId")

		if fileId == "" {
			var res ReponseStatus

			res.Status = -1
			res.Message = ResonseStatusMessage[res.Status]

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(res)
			return
		}

		if fileId, err := strconv.Atoi(fileId); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return

		} else if fileId < 0 {

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))

			return
		} else if node := tree.Find(fileId, tree.Root); node == nil {
			var res ReponseStatus

			res.Status = -1
			res.Message = ResonseStatusMessage[res.Status]

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(res)
		} else {
			if node.ObjectType == 2 {
				var fileList = []FileObject{
					{
						FileId:           node.FileId,
						ObjectType:       node.ObjectType,
						Name:             *node.Name,
						LastModifiedDate: node.LastModifiedDate,
					},
				}

				var res = GetChildrenResponse{
					FileList: fileList,
				}
				res.Status = 0
				res.Message = ResonseStatusMessage[res.Status]

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(res)

			} else {
				fileList := getChildren(node.Children)

				var res = GetChildrenResponse{
					FileList: fileList,
				}
				res.Status = 0
				res.Message = ResonseStatusMessage[res.Status]

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(res)

			}

		}
	}

}

func getChildren(children *[]model.Node) (fileList []FileObject) {
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
