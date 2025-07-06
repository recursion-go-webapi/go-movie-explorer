import { useMoviesPopular } from "@/hooks/useMoviesPopular";
import { MovieGrid } from "@/components/MovieGrid";

export function MoviePopular() {
  const { movies, loading, error } = useMoviesPopular();

  return (
    <div className="min-h-screen py-8">
      <div className="bg-white/95 backdrop-blur-md rounded-xl p-8 shadow-lg mb-8">
        <h1 className="text-3xl font-bold text-indigo-600 text-center mb-6">
          🌟 人気映画
        </h1>
        <MovieGrid 
          movies={movies} 
          loading={loading} 
          onMovieClick={(movie) => {
            window.location.href = `/movie/${movie.id}`;
          }}
        />
        {loading && (
          <div className="text-center text-gray-600 mt-4">
            <p>読み込み中...</p> 
          </div>
        )}
        {error && (
          <p className="text-red-500 text-center mt-4">
            エラーが発生しました: {error}
          </p>
        )}
        {!loading && !error && movies.length === 0 && (
          <p className="text-gray-600 text-center mt-4">
            人気映画はまだありません    
          </p>
        )}
      </div>
    </div>
  );
}