package models

import "time"

// ヘルスチェックレスポンス構造体
type HealthResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// NewHealthResponse はヘルスチェックレスポンスを生成する
func NewHealthResponse() *HealthResponse {
	return &HealthResponse{
		Status:    "healthy",
		Message:   "Go Web API is running",
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}
}
