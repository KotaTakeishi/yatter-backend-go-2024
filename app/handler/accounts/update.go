package accounts

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/domain/auth"
)

// Request body for `POST /v1/accounts/update_credentials`
type PutRequest struct {
	DisplayName *string `json:"display_name"`
	Note        *string `json:"note"`
	Avatar      *string `json:"avatar"`
	Header      *string `json:"header"`
}

// Handle request for `PUT /v1/accounts/update_credentials`
func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	var req PutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	account_info := auth.AccountOf(ctx)

	dto, err := h.accountUsecase.Update(ctx, account_info.ID, req.DisplayName, req.Note, req.Avatar, req.Header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.Account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
