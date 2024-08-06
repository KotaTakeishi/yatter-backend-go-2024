package health

import (
	"log"
	"net/http"
)

// Handle health check request
func NewRouter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		_, err := w.Write([]byte("OK"))
		if err != nil {
			// レスポンスボディの書き込みに失敗している かつ ステータスコードはレスポンスボディ書き込み後に変更できないのでログにエラーを出す
			log.Printf("health check: failed to write response body: %v\n", err)
		}
	}
}
