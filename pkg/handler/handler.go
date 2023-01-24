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

	ord, _ := h.services.Cache.Get(uid)
	data, _ := json.MarshalIndent(ord, "", "\t")
	res := string(data)
	if ord.Order_uid == "" {
		res = "Not found UID"
	}

	temp.Execute(w, res)
}

func (h *Handler) InitRouters() {
	http.Handle("/", http.FileServer(http.Dir("web/")))
	http.HandleFunc("/order", h.GetOrderById)

	log.Println("SERVER START (port 8080)")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
