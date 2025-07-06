import { useEffect, useState } from 'react';
import { getPopularMovies } from '@/api/index';
import type { Movie, MoviesResponse } from '@/types/movie';

export function useMoviesPopular() {
  const [movies, setMovies] = useState<Movie[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    setLoading(true);
    getPopularMovies(1)
      .then((res: MoviesResponse) => {
        setMovies(res.results.slice(0, 20));
        setError(null);
      })
      .catch((e: any) => setError(e.message))
      .finally(() => setLoading(false));
  }, []);

  return { movies, loading, error };
}
