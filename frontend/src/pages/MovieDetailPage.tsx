import { useParams } from 'react-router-dom';
import { useMovieDetail } from '@/hooks/useMovieDetail';
import { useNavigate } from 'react-router-dom';

export function MovieDetailPage() {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const { movie, loading, error } = useMovieDetail(id);

  if (loading) {
    return <div className="min-h-screen flex items-center justify-center">読み込み中...</div>;
  }
  if (error) {
    return <div className="min-h-screen flex items-center justify-center">エラー: {error}</div>;
  }
  if (!movie) {
    return <div className="min-h-screen flex items-center justify-center">映画が見つかりません</div>;
  }

  return (
    <div className="min-h-screen py-8">
      <button
        className="mb-6 bg-indigo-600 hover:bg-indigo-700 text-white px-4 py-2 rounded-lg transition-colors"
        onClick={() => navigate('/')}
      >
        ← 戻る
      </button>
      <div className="bg-white/95 backdrop-blur-md rounded-xl p-8 shadow-lg">
        <h1 className="text-3xl font-bold text-indigo-600 mb-6 text-center">
          🎦 {movie.title}
        </h1>
        <div className="flex flex-col md:flex-row gap-8 items-center">
          {movie.poster_path && (
            <img
              src={`https://image.tmdb.org/t/p/w300${movie.poster_path}`}
              alt={movie.title}
              className="rounded-lg shadow-md mb-4 md:mb-0"
            />
          )}
          <div className="flex-1 text-gray-700">
            <p className="mb-2"><span className="font-semibold">原題:</span> {movie.original_title}</p>
            <p className="mb-2"><span className="font-semibold">公開日:</span> {movie.release_date}</p>
            <p className="mb-2"><span className="font-semibold">ジャンル:</span> {movie.genres.map(g => g.name).join(', ')}</p>
            <p className="mb-2"><span className="font-semibold">言語:</span> {movie.original_language}</p>
            <p className="mb-2"><span className="font-semibold">人気度:</span> {movie.popularity}</p>
            <p className="mb-2"><span className="font-semibold">予算:</span> {movie.budget ? `${movie.budget.toLocaleString()} USD` : '不明'}</p>
            <p className="mb-2"><span className="font-semibold">IMDB:</span> {movie.imdb_id ? (
              <a href={`https://www.imdb.com/title/${movie.imdb_id}`} target="_blank" rel="noopener noreferrer" className="text-blue-600 underline">{movie.imdb_id}</a>
            ) : 'なし'}</p>
            <p className="mb-4"><span className="font-semibold">概要:</span> {movie.overview || '説明なし'}</p>
            {movie.homepage && (
              <a href={movie.homepage} target="_blank" rel="noopener noreferrer" className="text-indigo-600 underline">公式サイト</a>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}