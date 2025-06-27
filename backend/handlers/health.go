package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go-movie-explorer/models"
)

// HealthHandler はヘルスチェックエンドポイントのハンドラー
func HealthHandler(w http.ResponseWriter, r *http.Request) error {
	// JSONレスポンス用ヘッダー設定
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// ヘルスチェックレスポンス生成
	response := models.NewHealthResponse()

	// JSON形式でレスポンス返却
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return fmt.Errorf("failed to encode health response: %w", err)
	}
	return nil
}
