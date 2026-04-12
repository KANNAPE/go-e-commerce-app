package products

import (
	"log"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/kannape/go-e-commerce-app/internal/json"
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
	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.Write(w, http.StatusOK, products); err != nil {
		slog.Error("Something went wrong", "error", err)
	}
}

func (h *handler) FindProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		log.Println("Missing argument for product in GET request")
		http.Error(w, "Missing argument!", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Invalid ID in the GET request")
		http.Error(w, "Not an ID!", http.StatusBadRequest)
		return
	}

	product, err := h.service.FindProductByID(r.Context(), int64(id))
	if err != nil {
		log.Println("Product doesn't exist")
		http.Error(w, "Product doesn't exist", http.StatusNotFound)
		return
	}

	if err := json.Write(w, http.StatusOK, product); err != nil {
		slog.Error("Something went wrong", "error", err)
	}
}
