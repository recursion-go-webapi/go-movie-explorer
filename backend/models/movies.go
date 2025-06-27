package models

// 後ほど作成予定

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
// TMDBのレスポンスは"results"フィールドに映画配列が入っているため、
// それをパースしてMoviesResponseに詰め替える想定です。
type TmdbDiscoverResponse struct {
	Page         int     `json:"page"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
	Results      []Movie `json:"results"`
}
