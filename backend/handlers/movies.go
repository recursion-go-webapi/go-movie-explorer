package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"go-movie-explorer/services"
)

// 映画一覧取得APIハンドラー /api/movies
func MoviesHandler(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")

	// クエリパラメータ取得（今後、検索条件やソート条件などを追加する場合はここを編集）
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	page := 1
	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		_ = l // 仮実装なので未使用（ページネーションや件数制限を実装する場合はここを編集）
	}

	// サービス層でTMDB APIから映画一覧を取得（API仕様変更や他サービス連携時はここを編集）
	moviesResp, err := services.GetMoviesFromTMDB(page)
	if err != nil {
		return fmt.Errorf("TMDB API呼び出し失敗: %w", err)
	}

	// レスポンスをJSONで返却（レスポンス形式を変更したい場合はここを編集）
	if err := json.NewEncoder(w).Encode(moviesResp); err != nil {
		return fmt.Errorf("failed to encode response: %w", err)
	}
	return nil
}

// 映画詳細取得ハンドラー /api/movies/{id}
func MovieDetailHandler(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	prefix := "/api/movie/"
	if !strings.HasPrefix(r.URL.Path, prefix) {
        http.NotFound(w, r)
        return fmt.Errorf("無効なパス: %s", r.URL.Path)
    }
    id := strings.TrimPrefix(r.URL.Path, prefix)
    if id == "" || strings.Contains(id, "/") {
        http.NotFound(w, r)
        return fmt.Errorf("無効な映画ID: %s", id)
    }
	// IDを整数に変換
	movieID, err := strconv.Atoi(id)
	if err != nil {
		http.NotFound(w, r)
		return fmt.Errorf("無効な映画ID: %s", id)
	}
	// サービス層でTMDB APIから映画詳細を取得
	movieDetail, err := services.GetMovieDetailFromTMDB(movieID)

	if err != nil {
		http.Error(w, fmt.Sprintf("映画詳細取得失敗: %v", err), http.StatusInternalServerError)
		return fmt.Errorf("映画詳細取得失敗: %w", err)
	}
	// レスポンスをJSONで返却
	if err := json.NewEncoder(w).Encode(movieDetail); err != nil {
		return fmt.Errorf("failed to encode response: %w", err)
	}
	return nil
	
}

// 映画検索APIハンドラー /api/movies/search
func SearchMoviesHandler(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")

	// クエリパラメータ取得
	query := r.URL.Query().Get("query")
	if query == "" {
		return fmt.Errorf("検索クエリが指定されていません")
	}

	pageStr := r.URL.Query().Get("page")
	page := 1
	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}

	// サービス層でTMDB APIから映画検索結果を取得
	moviesResp, err := services.SearchMoviesFromTMDB(query, page)
	if err != nil {
		return fmt.Errorf("TMDB 検索API呼び出し失敗: %w", err)
	}

	// レスポンスをJSONで返却
	if err := json.NewEncoder(w).Encode(moviesResp); err != nil {
		return fmt.Errorf("failed to encode response: %w", err)
	}
	return nil
}
// - /api/movies/popular : 人気映画ランキング（今後追加予定）
//
// 新しいエンドポイントを追加する場合は、このファイルにハンドラー関数を追記してください。

func ListMoviesByGenreHandler(w http.ResponseWriter, r *http.Request) error {
	genreIDStr := r.URL.Query().Get("genre_id")
	pageStr := r.URL.Query().Get("page")

	if genreIDStr == "" {
		http.Error(w, "Missing required parameter: genre_id", http.StatusBadRequest)
		return fmt.Errorf("genre_id が指定されていません")
	}
	genreID, err := strconv.Atoi(genreIDStr)
	if err != nil {
		http.Error(w, "Invalid parameter: genre_id must be an integer", http.StatusBadRequest)
		return fmt.Errorf("genre_id の変換に失敗しました: %w", err)
	}

	page := 1
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	result, err := services.GetMoviesByGenre(genreID, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return fmt.Errorf("ジャンルの取得に失敗しました。: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	return nil
}
