export function SearchPage() {
  return (
    <div className="min-h-screen py-8">
      <div className="bg-white/95 backdrop-blur-md rounded-xl p-8 shadow-lg mb-8">
        <h1 className="text-3xl font-bold text-indigo-600 mb-6 text-center">
          🔍 リアルタイム検索
        </h1>
        
        <div className="space-y-6">
          <div className="flex gap-4">
            <input
              type="text"
              placeholder="映画タイトルを入力..."
              className="flex-1 px-4 py-3 border-2 border-gray-200 rounded-full focus:border-indigo-500 focus:outline-none transition-colors"
            />
            <button className="bg-indigo-600 hover:bg-indigo-700 text-white px-6 py-3 rounded-full font-semibold transition-colors">
              詳細検索へ
            </button>
          </div>
          
          <div className="text-center text-gray-600">
            映画タイトルを入力してください
          </div>
        </div>
      </div>

      {/* 検索結果表示エリア */}
      <div className="bg-white/95 backdrop-blur-md rounded-xl p-8 shadow-lg">
        <div className="text-center text-gray-600">
          <p className="text-lg mb-4">準備中...</p>
          <p>リアルタイム検索機能を実装予定です</p>
        </div>
      </div>
    </div>
  );
}