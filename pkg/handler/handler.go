package handler

import (
	"L0WB/pkg/service"
	"encoding/json"
	"html/template"
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
	temp, _ := template.ParseFiles("web/order.html")
	uid := r.FormValue("uid")

	ord, _ := h.services.Cache[uid]
	data, _ := json.MarshalIndent(ord, "", "\t")
	temp.Execute(w, string(data))

}

func (h *Handler) InitRouters() {
	http.Handle("/", http.FileServer(http.Dir("web/")))
	http.HandleFunc("/order", h.GetOrderById)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
