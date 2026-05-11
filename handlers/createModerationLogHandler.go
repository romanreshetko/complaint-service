package handlers

import (
	"complaint-service/models"
	"complaint-service/monitoring"
	"complaint-service/repository"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateModerationLogHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	claims, ok := r.Context().Value("claims").(models.AuthContext)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	if claims.Role != "service" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	var req models.CreateLogRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	err := repository.InsertModerationLog(h.db, 0, req.ContentID, req.ContentType, req.Result)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if req.Result == "moderation_error" {
		go monitoring.CheckModerationError(h.db, h.mailer)
	}

	w.WriteHeader(http.StatusOK)
}
