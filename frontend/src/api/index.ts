import type { MoviesResponse, MovieDetail, GenreMovieListResponse, APIError } from '@/types/movie';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

class APIClient {
  private baseURL: string;

  constructor(baseURL: string = API_BASE_URL) {
    this.baseURL = baseURL;
  }

  private async request<T>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<T> {
    const url = `${this.baseURL}${endpoint}`;
    
    try {
      const response = await fetch(url, {
        headers: {
          'Content-Type': 'application/json',
          ...options.headers,
        },
        ...options,
      });

      if (!response.ok) {
        const errorData: APIError = await response.json();
        throw new Error(errorData.message || `HTTP ${response.status}`);
      }

      return await response.json();
    } catch (error) {
      if (error instanceof Error) {
        throw error;
      }
      throw new Error('Network error occurred');
    }
  }

  async getMovies(page: number = 1): Promise<MoviesResponse> {
    return this.request<MoviesResponse>(`/api/movies?page=${page}`);
  }

  async getMovieDetail(id: number): Promise<MovieDetail> {
    return this.request<MovieDetail>(`/api/movie/${id}`);
  }

  async searchMovies(query: string, page: number = 1): Promise<MoviesResponse> {
    const encodedQuery = encodeURIComponent(query);
    return this.request<MoviesResponse>(`/api/movies/search?query=${encodedQuery}&page=${page}`);
  }

  async getMoviesByGenre(genreId: number, page: number = 1): Promise<GenreMovieListResponse> {
    return this.request<GenreMovieListResponse>(`/api/movies/genre?genre_id=${genreId}&page=${page}`);
  }

  async healthCheck(): Promise<{ status: string }> {
    return this.request<{ status: string }>('/healthz');
  }
}

export const apiClient = new APIClient();
export default apiClient;