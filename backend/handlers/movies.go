package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go-movie-explorer/models"
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

	// サービス層でTMDB APIから映画一覧を取得（仮実装）
	// TODO: services.GetMoviesFromTMDB(page, limit) などに置き換え
	resp := models.MoviesResponse{
		Page:         page,
		TotalPages:   1,
		TotalResults: 1,
		Movies: []models.Movie{
			{
				ID:          1,
				Title:       "サンプル映画",
				Overview:    "これはサンプルです",
				ReleaseDate: "2025-01-01",
				PosterPath:  "/sample.jpg",
				VoteAverage: 8.5,
				Popularity:  123.4,
			},
		},
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return fmt.Errorf("failed to encode response: %w", err)
	}
	return nil
}
