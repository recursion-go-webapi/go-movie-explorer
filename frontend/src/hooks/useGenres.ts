import { useState, useEffect } from 'react';
import { getGenres } from '@/api/index';
import type { Genre } from '@/types/movie';

export function useGenres() {
  const [genres, setGenres] = useState<Genre[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchGenres = async () => {
    setLoading(true);
    setError(null);
    try {
      const response = await getGenres();
      setGenres(response.genres);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'エラーが発生しました');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchGenres();
  }, []);

  return {
    genres,
    loading,
    error,
    refresh: fetchGenres,
  };
}