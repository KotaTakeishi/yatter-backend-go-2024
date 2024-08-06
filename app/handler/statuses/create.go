package statuses

import (
	"fmt"
	"net/http"
	"yatter-backend-go/app/domain/auth"
)

// Request body for `POST /v1/statuses`
type AddRequest struct {
	Status string
}

// Handle request for `POST /v1/statuses`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	account_info := auth.AccountOf(r.Context()) // 認証情報を取得する

	panic(fmt.Sprintf("Must Implement Status Creation And Check Acount Info %v", account_info))

}
