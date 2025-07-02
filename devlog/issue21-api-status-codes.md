# Issue #21: APIエラーレスポンスの適切なHTTPステータスコード対応

**日付**: 2025/07/02  
**ブランチ**: `feature/issue-21-api-status-codes`  
**担当者**: takeshi-arihori

## 概要

現在のAPIでは全てのエラーが500 Internal Server Errorとして返されており、クライアント側で適切なエラーハンドリングができない問題がありました。この問題を解決するため、カスタムAPIError型を実装し、適切なHTTPステータスコードでエラーレスポンスを返すように改善しました。

## 問題点

- 検索クエリ未指定: `500 Internal Server Error` → 本来は `400 Bad Request`
- 無効な映画ID: `500 Internal Server Error` → 本来は `400 Bad Request`  
- TMDB API失敗: `500 Internal Server Error` → これは正しい

## 解決策

### 1. 新規作成: `middleware/error.go`

```go
package middleware

import "net/http"

// APIError はHTTPステータスコード付きのカスタムエラー型
type APIError struct {
    StatusCode int    `json:"statusCode"`
    Message    string `json:"message"`
}

// Error はerrorインターフェースを実装
func (e *APIError) Error() string {
    return e.Message
}

// NewBadRequestError は400 Bad Requestエラーを作成
func NewBadRequestError(message string) *APIError {
    return &APIError{
        StatusCode: http.StatusBadRequest,
        Message:    message,
    }
}

// NewInternalServerError は500 Internal Server Errorを作成
func NewInternalServerError(message string) *APIError {
    return &APIError{
        StatusCode: http.StatusInternalServerError,
        Message:    message,
    }
}
```

### 2. 修正: `middleware/logging.go`

```go
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
}
```

### 3. 修正: `handlers/movies.go`

**主な変更点**:

1. **インポート追加**:
```go
import (
    // ...
    "go-movie-explorer/middleware"
    // ...
)
```

2. **適切なAPIError使用**:
```go
// 修正前
return fmt.Errorf("検索クエリが指定されていません")

// 修正後  
return middleware.NewBadRequestError("検索クエリが指定されていません")
```

3. **正常レスポンスに200 OK明示**:
```go
func MoviesHandler(w http.ResponseWriter, r *http.Request) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)  // ← 追加
    // ...
}
```

## 対応ステータスコード

| ステータスコード | 使用ケース | 例 |
|------------------|------------|-----|
| **200 OK** | 正常なレスポンス | 映画一覧取得成功、検索結果取得成功 |
| **400 Bad Request** | 不正なパラメータ | 検索クエリ未指定、無効な映画ID、ジャンルID未指定 |
| **500 Internal Server Error** | サーバー内部エラー | TMDB API接続失敗、JSONエンコードエラー |

## テスト方法

### curlによるAPIステータスコードテスト

**1. 400 Bad Request テスト**:
```bash
# 検索クエリ未指定
curl -v "http://localhost:8080/api/movies/search"
# 期待結果: HTTP/1.1 400 Bad Request

# 無効な映画ID
curl -v "http://localhost:8080/api/movie/abc"  
# 期待結果: HTTP/1.1 400 Bad Request

# ジャンルID未指定
curl -v "http://localhost:8080/api/movies/genre"
# 期待結果: HTTP/1.1 400 Bad Request
```

**2. 200 OK テスト**:
```bash
# 正常な検索クエリ
curl -v "http://localhost:8080/api/movies/search?query=batman"
# 期待結果: HTTP/1.1 200 OK

# 映画一覧取得
curl -v "http://localhost:8080/api/movies"
# 期待結果: HTTP/1.1 200 OK
```

### 実際のテスト結果

```bash
# ✅ 400 Bad Request - クエリ未指定
$ curl -v "http://localhost:8080/api/movies/search"
> GET /api/movies/search HTTP/1.1
< HTTP/1.1 400 Bad Request
< Content-Type: text/plain; charset=utf-8
検索クエリが指定されていません

# ✅ 200 OK - 正常なクエリ
$ curl -v "http://localhost:8080/api/movies/search?query=batman"  
> GET /api/movies/search?query=batman HTTP/1.1
< HTTP/1.1 200 OK
< Content-Type: application/json
{"page":1,"total_pages":...}

# ✅ 400 Bad Request - 無効なID
$ curl -v "http://localhost:8080/api/movie/abc"
> GET /api/movie/abc HTTP/1.1
< HTTP/1.1 400 Bad Request
< Content-Type: text/plain; charset=utf-8
無効な映画ID: abc
```

### ログ出力確認

```
2025/07/02 23:38:41 [GET] /api/movies/search - Start
2025/07/02 23:38:41 [GET] /api/movies/search - Error: 検索クエリが指定されていません
2025/07/02 23:38:57 [GET] /api/movies/search - Start  
2025/07/02 23:38:58 [GET] /api/movies/search - Success (took 445.567319ms)
2025/07/02 23:39:08 [GET] /api/movie/abc - Start
2025/07/02 23:39:08 [GET] /api/movie/abc - Error: 無効な映画ID: abc
```

## 関連Issue

- GitHub Issue #21: APIエラーレスポンスの適切なHTTPステータスコード対応

## 参考資料

- [Go HTTP Status Codes](https://go.dev/src/net/http/status.go)
- [REST API設計のベストプラクティス](https://restfulapi.net/http-status-codes/)
