import { useEffect, useState } from 'react';
import { getMovieDetail } from '@/api';
import type { MovieDetail } from '@/types/movie';

export function useMovieDetail(id?: string) {
  const [movie, setMovie] = useState<MovieDetail | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!id) return;
    setLoading(true);
    getMovieDetail(Number(id))
      .then((detail) => {
        setMovie(detail);
        setError(null);
      })
      .catch((e) => setError(e.message))
      .finally(() => setLoading(false));
  }, [id]);

  return { movie, loading, error };
}