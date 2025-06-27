package handlers

import (
	"encoding/json"
	"net/http"

	"go-movie-explorer/models"
)

// HealthHandler はヘルスチェックエンドポイントのハンドラー
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	// JSONレスポンス用ヘッダー設定
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// ヘルスチェックレスポンス生成
	response := models.NewHealthResponse()

	// JSON形式でレスポンス返却
	json.NewEncoder(w).Encode(response)
}
