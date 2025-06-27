package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-movie-explorer/handlers"
)

// ログ付きハンドラーラッパー
func logHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s", r.Method, r.URL.Path)
		handler(w, r)
	}
}

func main() {
	// ログファイル設定
	logFile, err := os.OpenFile("logs/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("ログファイルを開けません: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// ルーティング設定（ログ付き）
	http.HandleFunc("/health", logHandler(handlers.HealthHandler))

	// サーバー設定
	// TODO: 後ほど環境変数で設定できるようにする
	port := ":8080"
	fmt.Printf("Server starting on http://localhost%s\n", port)
	log.Printf("Server listening on port %s", port)

	// サーバー起動
	log.Fatal(http.ListenAndServe(port, nil))
}
