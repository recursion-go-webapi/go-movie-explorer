export function MoviesPage() {
  return (
    <div className="min-h-screen py-8">
      <div className="bg-white/95 backdrop-blur-md rounded-xl p-8 shadow-lg mb-8">
        <h1 className="text-3xl font-bold text-indigo-600 mb-6 text-center">
          🎦 映画一覧
        </h1>
        
        {/* 検索・フィルターセクション */}
        <div className="space-y-6">
          <div className="flex gap-4">
            <input
              type="text"
              placeholder="映画タイトルを入力..."
              className="flex-1 px-4 py-3 border-2 border-gray-200 rounded-full focus:border-indigo-500 focus:outline-none transition-colors"
            />
            <button className="bg-indigo-600 hover:bg-indigo-700 text-white px-6 py-3 rounded-full font-semibold transition-colors">
              検索
            </button>
          </div>
          
          <div className="flex flex-wrap gap-4">
            <select className="px-4 py-2 border-2 border-gray-200 rounded-full focus:border-indigo-500 focus:outline-none bg-white">
              <option value="">すべてのジャンル</option>
              <option value="28">アクション</option>
              <option value="35">コメディ</option>
              <option value="18">ドラマ</option>
              <option value="27">ホラー</option>
              <option value="10749">ロマンス</option>
              <option value="878">SF</option>
            </select>
            
            <select className="px-4 py-2 border-2 border-gray-200 rounded-full focus:border-indigo-500 focus:outline-none bg-white">
              <option value="">すべての年代</option>
              <option value="2024">2024年</option>
              <option value="2023">2023年</option>
              <option value="2022">2022年</option>
              <option value="2021">2021年</option>
              <option value="2020">2020年</option>
            </select>
            
            <select className="px-4 py-2 border-2 border-gray-200 rounded-full focus:border-indigo-500 focus:outline-none bg-white">
              <option value="popularity.desc">人気順</option>
              <option value="release_date.desc">最新順</option>
              <option value="vote_average.desc">評価順</option>
            </select>
          </div>
        </div>
      </div>

      {/* 映画一覧表示エリア */}
      <div className="bg-white/95 backdrop-blur-md rounded-xl p-8 shadow-lg">
        <div className="text-center text-gray-600">
          <p className="text-lg mb-4">準備中...</p>
          <p>既存の映画一覧機能を移行予定です</p>
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
            2
          </button>
          <button className="px-4 py-2 border-2 border-indigo-600 text-indigo-600 rounded hover:bg-indigo-600 hover:text-white transition-colors">
            3
          </button>
          <button className="px-4 py-2 border-2 border-indigo-600 text-indigo-600 rounded hover:bg-indigo-600 hover:text-white transition-colors">
            次 ›
          </button>
        </div>
      </div>
    </div>
  );
}