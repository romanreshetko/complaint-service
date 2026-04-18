package handlers

import (
	"complaint-service/models"
	"complaint-service/repository"
	service_integrations "complaint-service/service-integrations"
	"net/http"
	"strconv"
)

func (h *Handler) CreateReviewResolutionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	claims, ok := r.Context().Value("claims").(models.AuthContext)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	if claims.Role != "moderator" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	resolution := r.URL.Query().Get("resolution")
	if resolution != "blocked" && resolution != "published" {
		http.Error(w, "incorrect resolution", http.StatusBadRequest)
		return
	}
	complaintID, err := strconv.ParseInt(r.URL.Query().Get("complaint_id"), 10, 64)
	if err != nil {
		http.Error(w, "incorrect complaint_id", http.StatusBadRequest)
		return
	}
	reviewID, err := repository.GetReviewIDByComplaintID(h.db, complaintID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service_integrations.UpdateReviewStatus(reviewID, resolution)
	if err != nil {
		http.Error(w, "error updating review", http.StatusInternalServerError)
		return
	}

	err = repository.DeleteReviewComplaints(h.db, reviewID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
