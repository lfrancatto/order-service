package http

import (
	"encoding/json"
	"net/http"
	"order-service/internal/application"
	"order-service/internal/domain"
)

type Handler struct {
	usecase *application.CreateOrder
}

func NewHandler(uc *application.CreateOrder) *Handler {
	return &Handler{
		usecase: uc,
	}
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID     string
		UserID string
		Amount float64
	}

	_ = json.NewDecoder(r.Body).Decode(&req)

	order := domain.NewOrder(req.ID, req.UserID, req.Amount)

	err := h.usecase.Execute(r.Context(), order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
