import { useState, useEffect, useRef } from 'react';
import { getMoviesByGenre } from '@/api';
import type { Movie, GenreMovieListResponse } from '@/types/movie';

export function useMoviesByGenre(genreId: number) {
  const [movies, setMovies] = useState<Movie[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [currentPage, setCurrentPage] = useState(1);
  const [perPage, setPerPage] = useState(20);
  const [totalPages, setTotalPages] = useState(0);
  const [totalResults, setTotalResults] = useState(0);

  const fetchMoviesByGenre = async (genreId: number, page: number) => {
    setLoading(true);
    setError(null);

    try {
      const response: GenreMovieListResponse = await getMoviesByGenre(genreId, page);
      setMovies(response.results);
      setCurrentPage(response.page);
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
    const pageToFetch = prevGenreRef.current !== genreId ? 1 : currentPage;
    prevGenreRef.current = genreId;
    setCurrentPage(pageToFetch);
    fetchMoviesByGenre(genreId, pageToFetch);
  }, [genreId, currentPage]);

  const goToPage = (page: number) => {
    setCurrentPage(page);
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