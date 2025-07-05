import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useMovies } from '@/hooks/useMovies';
import { MovieGrid } from '@/components/MovieGrid';
import { Pagination } from '@/components/Pagination';
import type { Movie } from '@/types/movie';

export function MoviesPage() {
  const navigate = useNavigate();
  const [inputValue, setInputValue] = useState('');
  const {
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
  } = useMovies();

  // URLから検索クエリが変わった時に入力フィールドを同期
  useEffect(() => {
    setInputValue(searchQuery);
  }, [searchQuery]);

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault();
    if (inputValue.trim()) {
      searchMovies(inputValue.trim());
    }
  };

  const handleClearSearch = () => {
    setInputValue('');
    clearSearch();
  };

  const handleMovieClick = (movie: Movie) => {
    navigate(`/movie/${movie.id}`);
  };

  return (
    <div className="min-h-screen py-8">
      <div className="bg-white/95 backdrop-blur-md rounded-xl p-8 shadow-lg mb-8">
        <h1 className="text-3xl font-bold text-indigo-600 mb-6 text-center">
          🎦 映画一覧
        </h1>
        
        {/* 検索セクション */}
        <form onSubmit={handleSearch} className="space-y-6">
          <div className="flex gap-4">
            <input
              type="text"
              value={inputValue}
              onChange={(e) => setInputValue(e.target.value)}
              placeholder="映画タイトルを入力..."
              className="flex-1 px-4 py-3 border-2 border-gray-200 rounded-full focus:border-indigo-500 focus:outline-none transition-colors"
            />
            <button 
              type="submit"
              className="bg-indigo-600 hover:bg-indigo-700 text-white px-6 py-3 rounded-full font-semibold transition-colors"
            >
              検索
            </button>
            {searchQuery && (
              <button 
                type="button"
                onClick={handleClearSearch}
                className="bg-gray-500 hover:bg-gray-600 text-white px-6 py-3 rounded-full font-semibold transition-colors"
              >
                クリア
              </button>
            )}
          </div>
        </form>

        {/* 検索結果情報 */}
        {searchQuery && (
          <div className="mt-4 text-center text-gray-600">
            <p>"{searchQuery}" の検索結果: {totalResults}件</p>
          </div>
        )}
      </div>

      {/* エラー表示 */}
      {error && (
        <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-8">
          <p>エラーが発生しました: {error}</p>
          <button 
            onClick={refresh}
            className="mt-2 bg-red-600 hover:bg-red-700 text-white px-4 py-2 rounded transition-colors"
          >
            再試行
          </button>
        </div>
      )}

      {/* 映画一覧表示エリア */}
      <div className="bg-white/95 backdrop-blur-md rounded-xl p-8 shadow-lg">
        <MovieGrid 
          movies={movies}
          loading={loading}
          onMovieClick={handleMovieClick}
        />
      </div>

      {/* ページネーション */}
      {totalPages > 1 && (
        <Pagination
          currentPage={currentPage}
          totalPages={totalPages}
          onPageChange={goToPage}
          loading={loading}
        />
      )}
    </div>
  );
}