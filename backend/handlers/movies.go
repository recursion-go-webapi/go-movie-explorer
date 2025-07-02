package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"go-movie-explorer/middleware"
	"go-movie-explorer/services"
)

// 映画一覧取得APIハンドラー /api/movies
func MoviesHandler(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

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
		return middleware.NewInternalServerError(fmt.Sprintf("TMDB API呼び出し失敗: %v", err))
	}

	// レスポンスをJSONで返却（レスポンス形式を変更したい場合はここを編集）
	if err := json.NewEncoder(w).Encode(moviesResp); err != nil {
		return middleware.NewInternalServerError(fmt.Sprintf("JSONレスポンスのエンコードに失敗しました: %v", err))
	}
	return nil
}

// 映画詳細取得ハンドラー /api/movies/{id}
func MovieDetailHandler(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	prefix := "/api/movie/"
	if !strings.HasPrefix(r.URL.Path, prefix) {
		return middleware.NewBadRequestError(fmt.Sprintf("無効なパス: %s", r.URL.Path))
	}
	id := strings.TrimPrefix(r.URL.Path, prefix)
	if id == "" || strings.Contains(id, "/") {
		return middleware.NewBadRequestError(fmt.Sprintf("無効な映画ID: %s", id))
	}
	// IDを整数に変換
	movieID, err := strconv.Atoi(id)
	if err != nil {
		return middleware.NewBadRequestError(fmt.Sprintf("無効な映画ID: %s", id))
	}
	// サービス層でTMDB APIから映画詳細を取得
	movieDetail, err := services.GetMovieDetailFromTMDB(movieID)

	if err != nil {
		return middleware.NewInternalServerError(fmt.Sprintf("映画詳細取得失敗: %v", err))
	}
	
	w.WriteHeader(http.StatusOK)
	// レスポンスをJSONで返却
	if err := json.NewEncoder(w).Encode(movieDetail); err != nil {
		return middleware.NewInternalServerError(fmt.Sprintf("JSONレスポンスのエンコードに失敗しました: %v", err))
	}
	return nil
}

// 映画検索APIハンドラー /api/movies/search
func SearchMoviesHandler(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")

	// クエリパラメータ取得
	query := r.URL.Query().Get("query")
	if query == "" {
		return middleware.NewBadRequestError("検索クエリが指定されていません")
	}

	pageStr := r.URL.Query().Get("page")
	page := 1
	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}

	// サービス層でTMDB APIから映画検索結果を取得
	moviesResp, err := services.SearchMoviesFromTMDB(query, page)
	if err != nil {
		return middleware.NewInternalServerError(fmt.Sprintf("TMDB 検索API呼び出し失敗: %v", err))
	}

	w.WriteHeader(http.StatusOK)
	// レスポンスをJSONで返却
	if err := json.NewEncoder(w).Encode(moviesResp); err != nil {
		return middleware.NewInternalServerError(fmt.Sprintf("JSONレスポンスのエンコードに失敗しました: %v", err))
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
		return middleware.NewBadRequestError("genre_id が指定されていません")
	}
	genreID, err := strconv.Atoi(genreIDStr)
	if err != nil {
		return middleware.NewBadRequestError(fmt.Sprintf("genre_id の変換に失敗しました: %v", err))
	}

	page := 1
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	result, err := services.GetMoviesByGenreFromTMDB(genreID, page)
	if err != nil {
		return middleware.NewInternalServerError(fmt.Sprintf("ジャンルの取得に失敗しました。: %v", err))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
	return nil
}
