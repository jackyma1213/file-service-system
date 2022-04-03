package service

import (
	"encoding/json"
	"file-service/model"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func FileSystemUpdateService(tree *model.Tree) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fileId := chi.URLParam(r, "fileId")

		var req UpdateRequestBody

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return
		}

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
		} else if node, err := tree.Update(fileId, req.Content, req.Name); err != nil {

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return

		} else if node == nil {

			var res ReponseStatus

			res.Status = -1
			res.Message = ResonseStatusMessage[res.Status]

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(res)
		} else {
			res := UpdateResponse{
				FileId:           node.FileId,
				LastModifiedDate: node.LastModifiedDate,
			}

			res.Status = 0
			res.Message = ResonseStatusMessage[res.Status]
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(res)
		}
	}

}
