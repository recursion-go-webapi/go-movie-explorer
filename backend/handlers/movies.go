package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go-movie-explorer/services"
)

// MoviesHandler handles /api/movies requests
func MoviesHandler(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")

	// クエリパラメータ取得
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	page := 1
	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		_ = l // 仮実装なので未使用
	}

	// サービス層でTMDB APIから映画一覧を取得
	moviesResp, err := services.GetMoviesFromTMDB(page)
	if err != nil {
		return fmt.Errorf("TMDB API呼び出し失敗: %w", err)
	}

	if err := json.NewEncoder(w).Encode(moviesResp); err != nil {
		return fmt.Errorf("failed to encode response: %w", err)
	}
	return nil
}
