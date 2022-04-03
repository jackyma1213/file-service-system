package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"file-service/model"
	"file-service/service"
)

var tree *model.Tree
var fileId int = 0

func Init() *chi.Mux {

	tree = model.New()

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

	r.Get("/list", service.FileSystemGetChildrenService(tree))
	r.Get("/{fileId}", service.FileSystemGetFileService(tree))

	r.Post("/create", service.FileSystemCreateService(tree, &fileId))
	r.Post("/update/{fileId}", service.FileSystemUpdateService(tree))

	r.Delete("/{fileId}", service.FileSystemDeleteService(tree))

	return r
}
