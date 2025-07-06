package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"go-movie-explorer/services"
)

var startTime = time.Now()

// HealthHandler はTMDB APIバージョンを含むヘルスチェックハンドラー
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	// TMDBのPingを実行
	pinger := &services.TmdbPinger{}
	err := pinger.Ping(r.Context())
	
	// レスポンス作成
	var tmdbStatus string
	if err != nil {
		tmdbStatus = "CONNECTION_FAILED"
	} else {
		tmdbStatus = "SUCCESS"
	}
	
	response := map[string]interface{}{
		"uptime":  int(time.Since(startTime).Seconds()),
		"version": "TMDB-" + services.GetTMDBAPIVersion(),
		"status": map[string]string{
			"TMDB_API_CONNECTION": tmdbStatus,
		},
	}
	
	w.Header().Set("Content-Type", "application/json")
	
	// ヘルスチェックが失敗した場合は500を返す
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode health status", http.StatusInternalServerError)
	}
}