package middleware

import (
	"log"
	"net/http"
	"time"
)

type AppHandler func(http.ResponseWriter, *http.Request) error

// LoggingHandler はアクセスログとエラーハンドリングを統合したミドルウェア
// アクセスログの出力とエラー処理を一元化し、main.goのlogHandlerとErrorHandlerを統合
func LoggingHandler(h AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// リクエスト開始時刻を記録
		start := time.Now()

		// アクセスログの出力
		log.Printf("[%s] %s - Start", r.Method, r.URL.Path)

		// ハンドラーの実行とエラーハンドリング
		if err := h(w, r); err != nil {
			// エラーログの出力
			log.Printf("[%s] %s - Error: %v", r.Method, r.URL.Path, err)

			// APIErrorの場合は適切なステータスコードでレスポンス
			if apiErr, ok := err.(*APIError); ok {
				http.Error(w, apiErr.Message, apiErr.StatusCode)
			} else {
				// その他のエラーは500 Internal Server Error
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		} else {
			// 成功時のログ出力（実行時間も含む）
			duration := time.Since(start)
			log.Printf("[%s] %s - Success (took %v)", r.Method, r.URL.Path, duration)
		}
	}
}

// SimpleLoggingHandler はシンプルなアクセスログのみを出力するミドルウェア
// 元のlogHandlerと同等の機能（互換性のため残しておく）
func SimpleLoggingHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s", r.Method, r.URL.Path)
		next(w, r)
	}
}

// DetailedLoggingHandler はより詳細なログを出力するミドルウェア
// リクエストヘッダー、リモートアドレス、ユーザーエージェントなども記録
func DetailedLoggingHandler(h AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// 詳細なリクエスト情報をログ出力
		log.Printf("[%s] %s - From: %s, User-Agent: %s",
			r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())

		// ハンドラーの実行とエラーハンドリング
		if err := h(w, r); err != nil {
			duration := time.Since(start)
			log.Printf("[%s] %s - Error after %v: %v",
				r.Method, r.URL.Path, duration, err)

			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		} else {
			duration := time.Since(start)
			log.Printf("[%s] %s - Success in %v",
				r.Method, r.URL.Path, duration)
		}
	}
}

// RequestLoggingConfig はログ設定を定義
type RequestLoggingConfig struct {
	LogLevel      string // "simple", "standard", "detailed"
	LogDuration   bool   // 実行時間をログに含むか
	LogUserAgent  bool   // User-Agentをログに含むか
	LogRemoteAddr bool   // リモートアドレスをログに含むか
}

// DefaultLoggingConfig はデフォルトのログ設定を返す
func DefaultLoggingConfig() *RequestLoggingConfig {
	return &RequestLoggingConfig{
		LogLevel:      "standard",
		LogDuration:   true,
		LogUserAgent:  false,
		LogRemoteAddr: false,
	}
}

// ConfigurableLoggingHandler は設定可能なログミドルウェア
func ConfigurableLoggingHandler(config *RequestLoggingConfig) func(AppHandler) http.HandlerFunc {
	if config == nil {
		config = DefaultLoggingConfig()
	}

	return func(h AppHandler) http.HandlerFunc {
		switch config.LogLevel {
		case "detailed":
			return DetailedLoggingHandler(h)
		case "simple":
			// AppHandlerをhttp.HandlerFuncに変換してSimpleLoggingHandlerで包む
			return SimpleLoggingHandler(func(w http.ResponseWriter, r *http.Request) {
				if err := h(w, r); err != nil {
					log.Printf("handler error: %v", err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			})
		default: // "standard"
			return LoggingHandler(h)
		}
	}
}
