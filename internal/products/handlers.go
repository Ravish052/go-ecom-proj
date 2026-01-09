package products

import (
	"log"
	"net/http"

	"github.com/Ravish052/goEcon/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// 1. call the service method to get products

	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// return JSON in HTTP response

	json.Write(w, http.StatusOK, products)

}
