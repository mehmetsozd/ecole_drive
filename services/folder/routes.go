package folder

import (
	"ecole_drive/types"

	"github.com/gorilla/mux"
)

type Handler struct {
	store types.FolderStore
}

func NewHandler(store types.FolderStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) FolderRoutes(router *mux.Router) {
	router.HandleFunc("/folder", h.handleCreateFolder).Methods("POST")
	router.HandleFunc("/folder/{id}", h.handleUpdateFolder).Methods("PUT")
	router.HandleFunc("/folder/{id}", h.handleDeleteFolder).Methods("DELETE")
	router.HandleFunc("/folder", h.handleGetFolders).Methods("GET")
}