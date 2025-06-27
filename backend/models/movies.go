package models

// 映画一覧取得用モデル
// TMDB APIの生レスポンス（resultsフィールド）→ アプリのレスポンス（moviesフィールド）に詰め替えて返す設計です。
// 例：
// TMDB:   { ..., "results": [ { ... }, ... ] }
// アプリ: { ..., "movies":  [ { ... }, ... ] }

// 映画一覧取得
type Movie struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Overview    string  `json:"overview"`
	ReleaseDate string  `json:"release_date"`
	PosterPath  string  `json:"poster_path"`
	VoteAverage float64 `json:"vote_average"`
	Popularity  float64 `json:"popularity"`
}

type MoviesResponse struct {
	Page         int     `json:"page"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
	Movies       []Movie `json:"movies"`
}

// TMDB APIのレスポンスに合わせた構造体
type TmdbDiscoverResponse struct {
	Page         int     `json:"page"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
	Results      []Movie `json:"results"`
}

// --- 今後追加予定のエンドポイント用モデル ---
// 映画詳細取得API用モデル（/movie/{id}）
// type MovieDetail struct {
//     // TODO: 映画詳細情報のフィールドを定義
// }

// 映画検索API用モデル（/search/movie）
// type MovieSearchResponse struct {
//     // TODO: 検索結果用のフィールドを定義
// }

// 人気映画ランキングAPI用モデル（/movie/popular）
// type PopularMoviesResponse struct {
//     // TODO: 人気映画ランキング用のフィールドを定義
// }

// ジャンル別映画取得API用モデル（/discover/movie?with_genres=）
// type GenreMoviesResponse struct {
//     // TODO: ジャンル別映画取得用のフィールドを定義
// }
