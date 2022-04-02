package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"file-service/model"
	"file-service/service"
)

var db *model.Db
var fileId int = -1 //-1: storage is empty

func Init() *chi.Mux {

	db = model.New()

	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.Logger)
	//fileSystem
	r.Mount("/fileSystem", fileSystemRouter())

	//404 & 405
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("Method is not valid"))
	})

	return r
}

func fileSystemRouter() http.Handler {
	r := chi.NewRouter()

	//Create a folder or file
	r.Post("/create", service.FileSystemCreateService(db, &fileId))

	//Delete a folder or file.
	//When deleting a folder, all containing files and sub-folders will also be deleted
	r.Delete("/{fileId}", service.FileSystemDeleteService(db))

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
