package handler

import (
	"L0WB/pkg/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")

	ord, _ := h.services.Cache[uid]
	data, _ := json.MarshalIndent(ord, "", "\t")
	fmt.Fprintf(w, "%s", data)

}

func (h *Handler) InitRouters() {
	http.HandleFunc("/order", h.GetOrderById)
	http.Handle("/", http.FileServer(http.Dir("web/")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
