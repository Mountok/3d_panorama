package api

import (
	"3d_panorama_back/internals/app/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRoute(
	subjectHandler *handler.SubjectHandler,
) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/images", subjectHandler.Image).Methods(http.MethodGet)
	router.HandleFunc("/images-all",subjectHandler.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/3d-upload", subjectHandler.Upload).Methods(http.MethodPost)

	return router
}
