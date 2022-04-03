package service

import (
	"encoding/json"
	"file-service/model"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func FileSystemDeleteService(tree *model.Tree) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fileId := chi.URLParam(r, "fileId")

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

		} else if fileId <= 0 {

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))

			return
		} else if count, err := tree.Remove(fileId); err != nil {
			var res ReponseStatus

			res.Status = -1
			res.Message = ResonseStatusMessage[res.Status]

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(res)
			return
		} else {

			var res = DeleteResponse{
				FileId: fileId,
				Count:  count,
			}

			res.Status = 0
			res.Message = ResonseStatusMessage[res.Status]
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(res)

			return
		}

	}
}
