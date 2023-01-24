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

	ord, err := h.services.Cache.Get(uid)
	var res string
	if err != nil {
		log.Println(err)
		res = "Not found UID"
	} else {
		data, _ := json.MarshalIndent(ord, "", "\t")
		res = string(data)
	}

	temp.Execute(w, res)
}

func (h *Handler) InitRouters() {
	http.Handle("/", http.FileServer(http.Dir("web/")))
	http.HandleFunc("/order", h.GetOrderById)

	log.Println("SERVER START (port 8080)")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
