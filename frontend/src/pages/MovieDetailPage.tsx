import { useParams } from 'react-router-dom';

export function MovieDetailPage() {
  const { id } = useParams<{ id: string }>();

  return (
    <div className="min-h-screen py-8">
      <button className="mb-6 bg-indigo-600 hover:bg-indigo-700 text-white px-4 py-2 rounded-lg transition-colors">
        ← 戻る
      </button>
      
      <div className="bg-white/95 backdrop-blur-md rounded-xl p-8 shadow-lg">
        <h1 className="text-3xl font-bold text-indigo-600 mb-6 text-center">
          🎦 映画詳細
        </h1>
        
        <div className="text-center text-gray-600">
          <p className="text-lg mb-4">準備中...</p>
          <p>映画 ID: {id || '未指定'}</p>
          <p className="mt-2">映画詳細情報と関連映画表示機能を実装予定です</p>
        </div>
      </div>
    </div>
  );
}