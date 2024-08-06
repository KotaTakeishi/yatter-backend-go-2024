package auth

import (
	"encoding/json"
	"net/http"
)

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var reqBody requestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := h.ar.FindByUsername(r.Context(), reqBody.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type responseBody struct {
		Username string `json:"username,omitempty"`
		// エラー用
		Message string `json:"message,omitempty"`
	}
	w.Header().Set("Content-Type", "application/json")
	var respBody responseBody
	if account == nil {
		w.WriteHeader(http.StatusUnauthorized)
		respBody = responseBody{Message: "Invalid username or password"}
	} else if isValid := account.CheckPassword(reqBody.Password); !isValid {
		w.WriteHeader(http.StatusUnauthorized)
		respBody = responseBody{Message: "Invalid username or password"}
	} else {
		w.WriteHeader(http.StatusOK)
		respBody = responseBody{Username: account.Username}
	}
	if err := json.NewEncoder(w).Encode(respBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
