package service

import (
	"encoding/json"
	"file-service/model"
	"net/http"
	"time"
)

func FileSystemCreateService(db *model.Db, fileId *int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req CreateRequestBody
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
		}

		if req == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))

		}

		*fileId++
		var dbRequest = model.Item{
			FileId:           *fileId,
			Name:             req.Name,
			ObjectType:       req.ObjectType,
			ParentFieldId:    req.ParentFieldId,
			LastModifiedDate: time.Now().Format(time.RFC3339),
		}

		db.Add(dbRequest)

		var res = CreateResponse{
			FileId:           dbRequest.FileId,
			LastModifiedDate: dbRequest.LastModifiedDate,
		}
		res.Status = 0
		res.Message = ResonseStatusMessage[res.Status]

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
	}
}
