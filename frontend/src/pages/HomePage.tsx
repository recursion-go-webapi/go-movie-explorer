export function HomePage() {
  return (
    <div className="min-h-screen">
      {/* ヒーローセクション */}
      <section className="text-center py-16 text-white">
        <h1 className="text-5xl font-bold mb-6 drop-shadow-lg">
          🎬 Go Movie Explorer
        </h1>
        <p className="text-xl mb-8 opacity-90 max-w-2xl mx-auto">
          最新映画から名作まで、あなたの次の映画体験を見つけよう
        </p>
      </section>

      {/* 機能紹介セクション */}
      <section className="py-16">
        <div className="grid md:grid-cols-3 gap-8">
          <div className="bg-white/95 backdrop-blur-md p-8 rounded-xl shadow-lg hover:transform hover:scale-105 transition-all duration-300">
            <div className="text-5xl mb-6 text-center">🎦</div>
            <h3 className="text-xl font-bold text-indigo-600 mb-4 text-center">
              豊富な映画データベース
            </h3>
            <p className="text-gray-600 text-center">
              最新作から古典まで、幅広いジャンルの映画情報を網羅
            </p>
          </div>

          <div className="bg-white/95 backdrop-blur-md p-8 rounded-xl shadow-lg hover:transform hover:scale-105 transition-all duration-300">
            <div className="text-5xl mb-6 text-center">⚡</div>
            <h3 className="text-xl font-bold text-indigo-600 mb-4 text-center">
              リアルタイム検索
            </h3>
            <p className="text-gray-600 text-center">
              素早い検索機能で、お探しの映画を瞬時に見つけられます
            </p>
          </div>

          <div className="bg-white/95 backdrop-blur-md p-8 rounded-xl shadow-lg hover:transform hover:scale-105 transition-all duration-300">
            <div className="text-5xl mb-6 text-center">🏷️</div>
            <h3 className="text-xl font-bold text-indigo-600 mb-4 text-center">
              ジャンル別検索
            </h3>
            <p className="text-gray-600 text-center">
              アクション、ドラマ、コメディなど、好みのジャンルで絞り込み
            </p>
          </div>
        </div>
      </section>

      {/* 人気映画ランキング */}
      <section className="py-16">
        <div className="bg-white/95 backdrop-blur-md rounded-xl p-8 shadow-lg">
          <h2 className="text-3xl font-bold text-center text-indigo-600 mb-8">
            🔥 人気映画ランキング
          </h2>
          <div className="text-center text-gray-600">
            <p className="text-lg">準備中...</p>
            <p className="mt-2">人気映画ランキングを実装予定です</p>
          </div>
        </div>
      </section>
    </div>
  );
}
