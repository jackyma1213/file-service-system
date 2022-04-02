package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"file-service/storage"
)

var db *storage.Db

var fileId int = -1 //-1: storage is empty

func Start() {

	db = storage.New()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	//fileSystem
	r.Mount("/fileSystem", fileSystemRouter())

	//Get Status
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("welcome"))
	})

	http.ListenAndServe(":3000", r)
}

func fileSystemRouter() http.Handler {
	r := chi.NewRouter()

	//Create a folder or file
	r.Post("/create", func(w http.ResponseWriter, r *http.Request) {

		var req CreateRequestBody
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return
		}

		fileId++
		var dbRequest = storage.Item{
			FileId:           fileId,
			Name:             req.Name,
			ObjectType:       req.ObjectType,
			ParentFieldId:    req.ParentFieldId,
			LastModifiedDate: time.Now().Format(time.RFC3339),
		}

		db.Add(dbRequest)

		var res = CreateResponse{
			FileId:           dbRequest.FileId,
			LastModifiedDate: dbRequest.LastModifiedDate,
			Status:           0,
			Message:          "hi",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
	})

	//Delete a folder or file.
	//When deleting a folder, all containing files and sub-folders will also be deleted
	r.Delete("/{field}", func(w http.ResponseWriter, r *http.Request) {

	})

	//List a folder with a certain fileId with ALL its children including grandchildren.
	//If the file Id is a file, the fileList contains the file object alone
	r.Get("/list", func(w http.ResponseWriter, r *http.Request) {

	})

	//Update a file name and/or content or update a folder name. (Update content for a folder is not allowed)
	r.Post("/update/{field}", func(w http.ResponseWriter, r *http.Request) {

	})

	//Show the file content. Not applicable for folder
	r.Get("/{field}", func(w http.ResponseWriter, r *http.Request) {

	})

	return r
}
