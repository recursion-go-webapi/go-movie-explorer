package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-movie-explorer/handlers"
	"go-movie-explorer/middleware"
	"go-movie-explorer/services"

	"github.com/joho/godotenv"
	"goa.design/clue/health"
)

// ログ付きハンドラーラッパー
func logHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s", r.Method, r.URL.Path)
		handler(w, r)
	}
}

func main() {
	// .env読み込み
	_ = godotenv.Load(".env")

	// logsディレクトリ自動作成
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatalf("logsディレクトリの作成に失敗: %v", err)
	}

	// ログファイル設定
	logFile, err := os.OpenFile("logs/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("ログファイルを開けません: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// 環境変数取得(.envファイルに記載したPORTを取得)
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORTが設定されていません")
	}
	port = ":" + port

	// .envファイルに記載したTMDB_API_KEYを取得
	tmdbApiKey := os.Getenv("TMDB_API_KEY")
	if tmdbApiKey == "" {
		log.Fatal("TMDB_API_KEYが設定されていません")
	}

	// clue/healthによるhealthチェックエンドポイント
	checker := health.NewChecker(&services.TmdbPinger{})
	http.Handle("/healthz", health.Handler(checker))

	// 映画一覧取得
	http.HandleFunc("/api/movies", logHandler(middleware.ErrorHandler(handlers.MoviesHandler)))

	// ルーティング設定（新しいAPIエンドポイントを追加する場合はここに追記）
	//
	// - /api/movies/{id}    : 映画詳細取得（今後追加予定）
	// - /api/movies/search  : 映画検索（今後追加予定）
	// - /api/movies/popular : 人気映画ランキング（今後追加予定）
	// - /api/movies/genre   : ジャンル別映画取得（今後追加予定）
	//
	// 新しいエンドポイントを追加する場合は、ここにルーティングを追記してください。

	fmt.Printf("Server starting on http://localhost%s\n", port)
	log.Printf("Server listening on port %s", port)

	// サーバー起動
	log.Fatal(http.ListenAndServe(port, nil))
}
