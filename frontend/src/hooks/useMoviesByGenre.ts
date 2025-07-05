import { useState, useEffect, useRef } from 'react';
import { useSearchParams } from 'react-router-dom';
import { getMoviesByGenre } from '@/api';
import type { Movie, GenreMovieListResponse } from '@/types/movie';

export function useMoviesByGenre(genreId: number) {
  const [searchParams, setSearchParams] = useSearchParams();
  const [movies, setMovies] = useState<Movie[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [perPage, setPerPage] = useState(20);
  const [totalPages, setTotalPages] = useState(0);
  const [totalResults, setTotalResults] = useState(0);
  
  // URLクエリパラメータから現在のページを取得
  const currentPage = parseInt(searchParams.get('page') || '1');

  const fetchMoviesByGenre = async (genreId: number, page: number) => {
    setLoading(true);
    setError(null);

    try {
      const response: GenreMovieListResponse = await getMoviesByGenre(genreId, page);
      setMovies(response.results);
      setPerPage(response.per_page);
      setTotalPages(response.total_pages);
      setTotalResults(response.total_results);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An error occurred');
      setMovies([]);
    } finally {
      setLoading(false);
    }
  };

  const prevGenreRef = useRef<number | null>(null);
  useEffect(() => {
    if (!genreId) return;
    
    // ジャンルが変わった場合はページを1にリセット
    if (prevGenreRef.current !== genreId) {
      prevGenreRef.current = genreId;
      const newParams = new URLSearchParams(searchParams);
      newParams.set('page', '1');
      setSearchParams(newParams, { replace: true });
      return;
    }
    
    // 初回ロード時にページ番号がURLにない場合は page=1 を設定
    if (!searchParams.has('page')) {
      const newParams = new URLSearchParams(searchParams);
      newParams.set('page', '1');
      setSearchParams(newParams, { replace: true });
    } else {
      // URLパラメータがある場合は映画を取得
      fetchMoviesByGenre(genreId, currentPage);
    }
  }, [genreId, searchParams]);

  const goToPage = (page: number) => {
    const newParams = new URLSearchParams(searchParams);
    newParams.set('page', page.toString());
    setSearchParams(newParams);
  };
  
  const refresh = () => {
    fetchMoviesByGenre(genreId, currentPage);
  };

  return {
    movies,
    loading,
    error,
    currentPage,
    perPage,
    totalPages,
    totalResults,
    goToPage,
    refresh,
  };
}