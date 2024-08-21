package timeline

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *handler) FindPublicTimelines(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	limit := 40 // Maximum number of followings to get (Default 40, Max 80)
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if li, err := strconv.Atoi(limitStr); err == nil {
			limit = min(li, 80)
		}
	}

	dto, err := h.timelineUsecase.FindPublicTimelines(ctx, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.Timeline); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
