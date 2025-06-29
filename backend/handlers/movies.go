package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

// 映画関連APIのハンドラーをまとめるファイル
//
// - /api/movies/{id}    : 映画詳細取得（今後追加予定）
// - /api/movies/search  : 映画検索（今後追加予定）
// - /api/movies/popular : 人気映画ランキング（今後追加予定）
func PopularMoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// クエリパラメータ取得
	page := 1
	pageStr := r.URL.Query().Get("page")
	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}

	// サービス呼び出し
	resp, err := services.GetPopularMoviesFromTMDB(page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// レスポンス返却
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

// - /api/movies/genre   : ジャンル別映画取得（今後追加予定）
//
// 新しいエンドポイントを追加する場合は、このファイルにハンドラー関数を追記してください。
