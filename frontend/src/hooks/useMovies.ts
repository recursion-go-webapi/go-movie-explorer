import { useState, useEffect } from 'react';
import type { Movie, MoviesResponse } from '@/types/movie';
import { apiClient } from '@/api';

export function useMovies() {
  const [movies, setMovies] = useState<Movie[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(0);
  const [totalResults, setTotalResults] = useState(0);
  const [searchQuery, setSearchQuery] = useState<string>('');

  const fetchMovies = async (page: number = 1, query?: string) => {
    setLoading(true);
    setError(null);

    try {
      let response: MoviesResponse;
      
      if (query) {
        response = await apiClient.searchMovies(query, page);
      } else {
        response = await apiClient.getMovies(page);
      }

      setMovies(response.results);
      setCurrentPage(response.page);
      setTotalPages(response.total_pages);
      setTotalResults(response.total_results);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An error occurred');
      setMovies([]);
    } finally {
      setLoading(false);
    }
  };

  const searchMovies = (query: string) => {
    setSearchQuery(query);
    setCurrentPage(1);
    fetchMovies(1, query);
  };

  const clearSearch = () => {
    setSearchQuery('');
    setCurrentPage(1);
    fetchMovies(1);
  };

  const goToPage = (page: number) => {
    setCurrentPage(page);
    fetchMovies(page, searchQuery || undefined);
  };

  const refresh = () => {
    fetchMovies(currentPage, searchQuery || undefined);
  };

  useEffect(() => {
    fetchMovies();
  }, []);

  return {
    movies,
    loading,
    error,
    currentPage,
    totalPages,
    totalResults,
    searchQuery,
    searchMovies,
    clearSearch,
    goToPage,
    refresh,
  };
}