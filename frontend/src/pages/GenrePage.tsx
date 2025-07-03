import { useParams } from 'react-router-dom';

const genres: Record<string, { name: string; description: string }> = {
  '28': { name: 'アクション', description: 'スリル満点のアクション映画' },
  '35': { name: 'コメディ', description: '笑いあふれるコメディ映画' },
  '18': { name: 'ドラマ', description: '心に響くドラマ映画' },
  '27': { name: 'ホラー', description: '恐怖と緊張のホラー映画' },
  '10749': { name: 'ロマンス', description: '愛と感動のロマンス映画' },
  '878': { name: 'SF', description: '未来とテクノロジーのSF映画' }
};

export function GenrePage() {
  const { id } = useParams<{ id: string }>();
  const genre = id ? genres[id] : null;

  return (
    <div className="min-h-screen py-8">
      <div className="bg-white/95 backdrop-blur-md rounded-xl p-8 shadow-lg text-center mb-8">
        <h1 className="text-3xl font-bold text-indigo-600 mb-4">
          🏷️ {genre?.name || 'ジャンル別映画'}
        </h1>
        <p className="text-gray-600">
          {genre?.description || '選択されたジャンルの映画一覧'}
        </p>
      </div>
      
      <div className="bg-white/95 backdrop-blur-md rounded-xl p-8 shadow-lg">
        <div className="text-center text-gray-600">
          <p className="text-lg mb-4">準備中...</p>
          <p>ジャンル ID: {id || '未指定'}</p>
          <p className="mt-2">ジャンル別映画一覧表示機能を実装予定です</p>
        </div>
      </div>
      
      {/* ページネーション */}
      <div className="flex justify-center mt-8">
        <div className="flex gap-2">
          <button className="px-4 py-2 border-2 border-indigo-600 text-indigo-600 rounded hover:bg-indigo-600 hover:text-white transition-colors">
            ‹ 前
          </button>
          <button className="px-4 py-2 bg-indigo-600 text-white rounded">
            1
          </button>
          <button className="px-4 py-2 border-2 border-indigo-600 text-indigo-600 rounded hover:bg-indigo-600 hover:text-white transition-colors">
            次 ›
          </button>
        </div>
      </div>
    </div>
  );
}