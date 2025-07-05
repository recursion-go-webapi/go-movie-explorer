import { useParams } from "react-router-dom";
import { useNavigate } from "react-router-dom";
import { useMoviesByGenre } from "@/hooks/useMoviesByGenre";
import { useGenres } from "@/hooks/useGenres";
import { MovieGrid } from "@/components/MovieGrid";
import { Pagination } from "@/components/Pagination";
import type { Movie } from "@/types/movie";

export function GenrePage() {
  const { genres, loading: genresLoading, error: genresError } = useGenres();
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();

  // デフォルトでid=28を使用（idパラメータがない場合）
  const genreId = id || '28';

  const genreMap = genres.reduce((acc, genre) => {
    acc[genre.id.toString()] = genre;
    return acc;
  }, {} as Record<string, { id: number; name: string }>);
  const genre = genreMap[genreId] || null;

  const { movies, loading, error, currentPage, totalPages, goToPage } =
    useMoviesByGenre(genre ? genre.id : parseInt(genreId));
  const handleMovieClick = (movie: Movie) => {
    navigate(`/movie/${movie.id}`);
  };

  if (genresLoading) {
    return (
      <p className="text-center text-indigo-600 font-semibold mt-8">
        ジャンルを読み込み中...
      </p>
    );
  }

  if (genresError) {
    return (
      <p className="text-center text-red-500 font-semibold mt-8">
        ジャンルの取得中にエラーが発生しました: {genresError}
      </p>
    );
  }

  return (
    <div className="min-h-screen py-8">
      <div className="bg-white/95 backdrop-blur-md rounded-xl p-8 shadow-lg mb-8 relative">
        <h1 className="text-3xl font-bold text-indigo-600 text-center">
          🏷️ {genre?.name || "ジャンル別映画"}
        </h1>
        {genres.length > 0 && (
          <div className="absolute top-1/2 right-4 transform -translate-y-1/2">
            <select
              className="border border-gray-300 rounded px-4 py-2 text-gray-700"
              value={genreId}
              onChange={(e) => navigate(`/genre/${e.target.value}`)}
            >
              {genres.map((g) => (
                <option key={g.id} value={g.id}>
                  {g.name}
                </option>
              ))}
            </select>
          </div>
        )}
      </div>

      <div className="bg-white/95 backdrop-blur-md rounded-xl p-8 shadow-lg">
        <div className="text-center text-gray-600">
          <p className="text-lg mb-2 font-semibold">
            {genre?.name ? `${genre.name}映画一覧` : "ジャンル別映画一覧"}
          </p>
          <p className="text-sm text-gray-500 mb-4">
            ジャンル ID: {genreId}
          </p>

          {error && (
            <p className="text-red-500 font-semibold mb-4">
              エラーが発生しました: {error}
            </p>
          )}
          {loading && movies.length === 0 ? (
            <p className="text-indigo-600 font-semibold">読み込み中...</p>
          ) : (
            <MovieGrid
              movies={movies}
              loading={loading}
              onMovieClick={handleMovieClick}
            />
          )}
        </div>
      </div>

      {/* ページネーション */}
      <div className="flex justify-center mt-8">
        <Pagination
          currentPage={currentPage}
          totalPages={totalPages}
          onPageChange={goToPage}
          loading={loading}
        />
      </div>
    </div>
  );
}
