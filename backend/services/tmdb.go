package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"go-movie-explorer/models"
)

const BaseURL = "https://api.themoviedb.org/3"

// TMDBのAPIキーを環境変数から取得
func GetTMDBApiKey() string {
	return os.Getenv("TMDB_API_KEY")
}

// GetMoviesFromTMDBはTMDBの/discover/movieから映画一覧を取得する
func GetMoviesFromTMDB(page int) (*models.MoviesResponse, error) {
	apiKey := GetTMDBApiKey()
	if apiKey == "" {
		return nil, fmt.Errorf("TMDB_API_KEYが設定されていません")
	}

	url := fmt.Sprintf("%s/discover/movie?page=%d", BaseURL, page)
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("リクエスト作成失敗: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("TMDB APIリクエスト失敗: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDB APIエラー: status=%d", resp.StatusCode)
	}

	var tmdbResp models.TmdbDiscoverResponse
	if err := json.NewDecoder(resp.Body).Decode(&tmdbResp); err != nil {
		return nil, fmt.Errorf("TMDBレスポンスのデコード失敗: %w", err)
	}

	// 独自のMoviesResponseに詰め替え
	return &models.MoviesResponse{
		Page:         tmdbResp.Page,
		TotalPages:   tmdbResp.TotalPages,
		TotalResults: tmdbResp.TotalResults,
		Movies:       tmdbResp.Results,
	}, nil
}
