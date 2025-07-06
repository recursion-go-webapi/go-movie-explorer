package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"go-movie-explorer/handlers"   // ハンドラー
	"go-movie-explorer/middleware" // ミドルウェア

	"github.com/joho/godotenv" // .envファイルの読み込み
)

func main() {
	// .env読み込み（ファイルが存在しない場合は無視）
	if err := godotenv.Load(".env"); err != nil {
		log.Printf(".envファイルが見つかりません（環境変数が設定されていれば正常動作します）: %v", err)
	}

	// logsディレクトリ自動作成
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatalf("logsディレクトリの作成に失敗: %v", err)
	}

	// ログファイル設定（コンソールとファイル両方に出力）
	logFile, err := os.OpenFile("logs/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("ログファイルを開けません: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

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

	// セキュリティミドルウェアの設定
	securityConfig := middleware.DefaultSecurityConfig()

	// 本番環境の場合はより厳格な設定を使用
	if os.Getenv("GO_ENV") == "production" {
		frontendURL := os.Getenv("FRONTEND_URL")
		securityConfig = middleware.ProductionSecurityConfig(frontendURL)
	}

	// ルートマルチプレクサーを作成
	mux := http.NewServeMux()

	// セキュリティミドルウェアを全体に適用
	securedHandler := middleware.SecurityMiddleware(securityConfig)(mux)

	// ヘルスチェックエンドポイント
	mux.HandleFunc("/healthz", handlers.HealthHandler)

	// 映画一覧取得
	mux.HandleFunc("/api/movies", middleware.LoggingHandler(handlers.MoviesHandler))

	// - /api/genres : ジャンル一覧取得
	mux.HandleFunc("/api/genres", middleware.LoggingHandler(handlers.GenresHandler))

	// 映画ジャンル別取得
	mux.HandleFunc("/api/movies/genre", middleware.LoggingHandler(handlers.ListMoviesByGenreHandler))

	// - /api/movies/{id} : 映画詳細取得APIエンドポイント
	mux.HandleFunc("/api/movie/", middleware.LoggingHandler(handlers.MovieDetailHandler))

	// - /api/movies/search：映画検索APIエンドポイント
	mux.HandleFunc("/api/movies/search", middleware.LoggingHandler(handlers.SearchMoviesHandler))

	// - /api/movies/popular : 人気映画ランキング（今後追加予定）
	mux.HandleFunc("/api/movies/popular", middleware.LoggingHandler(handlers.PopularMoviesHandler))

	log.Printf("Server starting on http://localhost%s\n", port)
	log.Printf("Server listening on port %s", port)
	log.Printf("Security middleware enabled with CORS origins: %v", securityConfig.AllowedOrigins)

	// サーバー起動（セキュリティミドルウェア適用済み）
	log.Fatal(http.ListenAndServe(port, securedHandler))
}
