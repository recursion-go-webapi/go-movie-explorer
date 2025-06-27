package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-movie-explorer/handlers"
	"go-movie-explorer/middleware"

	"github.com/joho/godotenv"
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

	// 環境変数取得
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORTが設定されていません")
	}
	port = ":" + port

	tmdbApiKey := os.Getenv("TMDB_API_KEY")
	if tmdbApiKey == "" {
		log.Fatal("TMDB_API_KEYが設定されていません")
	}

	// ルーティング設定（ログ付き）
	http.HandleFunc("/health", logHandler(middleware.ErrorHandler(handlers.HealthHandler)))
	http.HandleFunc("/api/movies", logHandler(middleware.ErrorHandler(handlers.MoviesHandler)))

	fmt.Printf("Server starting on http://localhost%s\n", port)
	log.Printf("Server listening on port %s", port)

	// サーバー起動
	log.Fatal(http.ListenAndServe(port, nil))
}
