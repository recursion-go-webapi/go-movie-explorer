import { MovieGrid } from '@/components/MovieGrid';
import { SearchBar } from '@/components/SearchBar';
import { Pagination } from '@/components/Pagination';
import { useMovies } from '@/hooks/useMovies';
import type { Movie } from '@/types/movie';

function App() {
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
  } = useMovies();

  const handleMovieClick = (movie: Movie) => {
    console.log('Movie clicked:', movie);
  };

  return (
    <div className="min-h-screen bg-background">
      <div className="container mx-auto px-4 py-8">
        <header className="text-center mb-8">
          <h1 className="text-4xl font-bold mb-4">🎬 Go Movie Explorer</h1>
          <p className="text-muted-foreground mb-6">
            最新の映画を検索・閲覧できるアプリケーション
          </p>
          <SearchBar
            onSearch={searchMovies}
            onClear={clearSearch}
            loading={loading}
          />
        </header>

        <main>
          {error && (
            <div className="bg-destructive/10 border border-destructive text-destructive px-4 py-3 rounded mb-6">
              <p className="font-semibold">エラーが発生しました</p>
              <p>{error}</p>
            </div>
          )}

          <div className="mb-6">
            {searchQuery ? (
              <p className="text-sm text-muted-foreground">
                「{searchQuery}」の検索結果: {totalResults.toLocaleString()}件
              </p>
            ) : (
              <p className="text-sm text-muted-foreground">
                人気の映画: {totalResults.toLocaleString()}件
              </p>
            )}
          </div>

          <MovieGrid
            movies={movies}
            onMovieClick={handleMovieClick}
            loading={loading}
          />

          <Pagination
            currentPage={currentPage}
            totalPages={totalPages}
            onPageChange={goToPage}
            loading={loading}
          />
        </main>
      </div>
    </div>
  );
}

export default App
