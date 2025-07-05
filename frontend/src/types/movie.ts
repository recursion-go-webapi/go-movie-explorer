export interface Movie {
  id: number;
  title: string;
  overview: string;
  release_date: string;
  poster_path: string;
  vote_average: number;
  popularity: number;
  genre_ids?: number[];
  vote_count?: number;
}

export interface MovieDetail extends Movie {
  original_title: string;
  backdrop_path: string;
  genres: Genre[];
  homepage: string;
  imdb_id: string;
  budget: number;
  origin_country: string[];
  original_language: string;
}

export interface Genre {
  id: number;
  name: string;
}

export interface GenreListResponse {
  genres: Genre[];
}

export interface MoviesResponse {
  page: number;
  total_pages: number;
  total_results: number;
  results: Movie[];
}

export interface GenreMovieListResponse {
  genre_id: number;
  page: number;
  per_page: number;
  total_pages: number;
  total_results: number;
  results: Movie[];
}

export interface APIError {
  statusCode: number;
  message: string;
}