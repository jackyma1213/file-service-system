package service

import (
	"encoding/json"
	"file-service/model"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

func FileSystemCreateService(db *model.Db, fileId *int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req CreateRequestBody

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return
		}

		validate := validator.New()

		if err := validate.Struct(req); err != nil {

			fmt.Println(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return
		}

		if match, _ := regexp.MatchString("^[a-zA-Z0-9]*$", req.Name); !match {
			var res ReponseStatus

			res.Status = -2
			res.Message = ResonseStatusMessage[res.Status]

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(res)
			return
		}

		if *req.ParentFileId > 0 {

			if _, err := db.Get(*req.ParentFileId); err != nil {
				var res ReponseStatus

				res.Status = -1
				res.Message = ResonseStatusMessage[res.Status]

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(res)
				return
			}
		}

		*fileId++
		var dbRequest = model.Item{
			FileId:           *fileId,
			Name:             req.Name,
			ObjectType:       *req.ObjectType,
			ParentFileId:     *req.ParentFileId,
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
