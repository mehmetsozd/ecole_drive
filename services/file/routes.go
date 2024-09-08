package file

import (
	"ecole_drive/types"

	"github.com/gorilla/mux"
)

type Handler struct {
	store types.FileStore
}

func NewHandler(store types.FileStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) FileRoutes(router *mux.Router) {
	router.HandleFunc("/file", h.handleUploadFile).Methods("POST")
	router.HandleFunc("/file/{id}", h.handleUpdateFile).Methods("PUT")
	router.HandleFunc("/file/{id}", h.handleDeleteFile).Methods("DELETE")
	router.HandleFunc("/file", h.handleGetFiles).Methods("GET")
}