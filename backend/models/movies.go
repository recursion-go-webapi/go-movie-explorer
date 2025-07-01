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
type BelongsToCollection struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PosterPath   string `json:"poster_path"`
	BackdropPath string `json:"backdrop_path"`
}

type TmdbMovieDetailResponse struct {
	Adult               bool                 `json:"adult"`
	BackdropPath        string               `json:"backdrop_path"`
	BelongsToCollection *BelongsToCollection `json:"belongs_to_collection,omitempty"`
	Budget              int                  `json:"budget"`
	Genres              []Genre              `json:"genres"`
	Homepage            string               `json:"homepage"`
	ID                  int                  `json:"id"`
	IMDBID              string               `json:"imdb_id"`
	OriginCountry       []string             `json:"origin_country"`
	OriginalLanguage    string               `json:"original_language"`
	OriginalTitle       string               `json:"original_title"`
	Overview            string               `json:"overview"`
	Popularity          float64              `json:"popularity"`
	PosterPath          string               `json:"poster_path"`
	ReleaseDate         string               `json:"release_date"`
	Title               string               `json:"title"`
}

type MovieDetail struct {
	ID               int      `json:"id"`
	Title            string   `json:"title"`
	OriginalTitle    string   `json:"original_title"`
	Overview         string   `json:"overview"`
	ReleaseDate      string   `json:"release_date"`
	PosterPath       string   `json:"poster_path"`
	BackdropPath     string   `json:"backdrop_path"`
	Genres           []Genre  `json:"genres"`
	Homepage         string   `json:"homepage"`
	IMDBID           string   `json:"imdb_id"`
	Popularity       float64  `json:"popularity"`
	Budget           int      `json:"budget"`
	OriginCountry    []string `json:"origin_country"`
	OriginalLanguage string   `json:"original_language"`
}

// ジャンル用モデル
type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// 映画検索API用モデル（/search/movie）
// type MovieSearchResponse struct {
//     // TODO: 検索結果用のフィールドを定義
// }

// 人気映画ランキングAPI用モデル（/movie/popular）
// type PopularMoviesResponse struct {
//     // TODO: 人気映画ランキング用のフィールドを定義
// }

// ジャンル別映画取得API用モデル（/discover/movie?with_genres=）
// TMDBのジャンル別映画リストAPIの全体レスポンスを表す構造体
type TMDBGenreMovieList struct {
	Page         int                   `json:"page"`
	Results      []GenreMoviesResponse `json:"results"`
	TotalPages   int                   `json:"total_pages"`
	TotalResults int                   `json:"total_results"`
}

// TMDBの1件分の映画データ（ジャンル検索時のフォーマット）
type GenreMoviesResponse struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Overview    string  `json:"overview"`
	ReleaseDate string  `json:"release_date"`
	GenreIDs    []int   `json:"genre_ids"`
	PosterPath  string  `json:"poster_path"`
	VoteAverage float64 `json:"vote_average"`
	Popularity  float64 `json:"popularity"`
	VoteCount   int     `json:"vote_count"`
}

// ジャンル別映画リストのレスポンス構造体
type GenreMovieListResponse struct {
	GenreID      int            `json:"genre_id"`
	Page         int            `json:"page"`
	PerPage      int            `json:"per_page"`
	TotalPages   int            `json:"total_pages"`
	TotalResults int            `json:"total_results"`
	Results      []MovieByGenre `json:"results"`
}

type MovieByGenre = GenreMoviesResponse

type GenreListResponse struct {
	Genres []Genre `json:"genres"`
}
