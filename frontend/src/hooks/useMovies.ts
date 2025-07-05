import { useState, useEffect } from 'react';
import { useSearchParams } from 'react-router-dom';
import type { Movie, MoviesResponse } from '@/types/movie';
import { getMovies, searchMovies as searchMoviesApi } from '@/api';

export function useMovies() {
  const [searchParams, setSearchParams] = useSearchParams();
  const [movies, setMovies] = useState<Movie[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [totalPages, setTotalPages] = useState(0);
  const [totalResults, setTotalResults] = useState(0);
  
  // URLクエリパラメータから現在の状態を取得
  const currentPage = parseInt(searchParams.get('page') || '1');
  const searchQuery = searchParams.get('query') || '';

  const fetchMovies = async (page: number = 1, query?: string) => {
    setLoading(true);
    setError(null);

    try {
      let response: MoviesResponse;

      if (query) {
        response = await searchMoviesApi(query, page);
      } else {
        response = await getMovies(page);
      }

      setMovies(response.results);
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
    const newParams = new URLSearchParams();
    newParams.set('query', query);
    newParams.set('page', '1');
    setSearchParams(newParams);
  };

  const clearSearch = () => {
    const newParams = new URLSearchParams();
    newParams.set('page', '1');
    setSearchParams(newParams);
  };

  const goToPage = (page: number) => {
    const newParams = new URLSearchParams(searchParams);
    newParams.set('page', page.toString());
    setSearchParams(newParams);
  };

  const refresh = () => {
    fetchMovies(currentPage, searchQuery || undefined);
  };

  useEffect(() => {
    // 初回ロード時にページ番号がURLにない場合は page=1 を設定
    if (!searchParams.has('page')) {
      const newParams = new URLSearchParams(searchParams);
      newParams.set('page', '1');
      setSearchParams(newParams, { replace: true });
    } else {
      // URLパラメータがある場合は映画を取得
      fetchMovies(currentPage, searchQuery || undefined);
    }
  }, [searchParams]);

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
