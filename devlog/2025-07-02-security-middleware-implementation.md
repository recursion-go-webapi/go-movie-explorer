# セキュリティミドルウェア実装 - Issue #14

**日付**: 2025-07-02  
**担当者**: takeshi-arihori  
**Issue**: #14 セキュリティヘッダーミドルウェアの実装  
**ブランチ**: feat/issue-14-security-middleware

## 📋 概要

Webアプリケーションのセキュリティ強化のため、包括的なセキュリティヘッダーミドルウェアを実装しました。

## 🔒 実装目標

- [x] CORS（Cross-Origin Resource Sharing）設定
- [x] XSS Protection ヘッダー
- [x] Content Security Policy（CSP）
- [x] X-Frame-Options（クリックジャッキング防止）
- [x] その他のセキュリティヘッダー

## 📂 実装したファイル

### 1. `backend/middleware/security.go`

**目的**: セキュリティヘッダーを自動付与するミドルウェア

**主要機能**:
- `SecurityConfig`: セキュリティ設定の構造体
- `DefaultSecurityConfig()`: 開発環境用のデフォルト設定
- `ProductionSecurityConfig()`: 本番環境用の厳格な設定
- `SecurityMiddleware()`: メインのミドルウェア関数

**実装されたセキュリティヘッダー**:
- **CORS設定**: オリジン制御、プリフライトリクエスト対応
- **X-XSS-Protection**: `1; mode=block` - XSS攻撃防止
- **X-Content-Type-Options**: `nosniff` - MIMEタイプスニッフィング防止
- **X-Frame-Options**: `DENY` - クリックジャッキング防止
- **Content-Security-Policy**: React・TMDB用に最適化
- **Referrer-Policy**: `strict-origin-when-cross-origin`
- **Permissions-Policy**: 位置情報、マイク、カメラ無効化

### 2. `backend/main.go` の変更

**変更点**:
- マルチプレクサー(`http.NewServeMux`)の導入
- セキュリティミドルウェアを全ルートに適用
- 環境変数による設定切り替え（開発・本番）

## 🔧 CORS設定詳細

### 許可されたオリジン
```go
AllowedOrigins: []string{
    frontendURL,                 // 環境変数から取得
}
```

### 許可されたメソッド
```go
AllowedMethods: []string{
    "GET", "POST", "PUT", "DELETE", "OPTIONS",
}
```

### 許可されたヘッダー
```go
AllowedHeaders: []string{
    "Origin", "Content-Type", "Accept", "Authorization",
    "X-Requested-With", "X-HTTP-Method-Override",
}
```

## 🛡️ Content Security Policy (CSP)

### 開発環境用CSP
```go
CSPDirectives: map[string]string{
    "default-src": "'self'",
    "script-src":  "'self' 'unsafe-inline' 'unsafe-eval'", // React開発用
    "style-src":   "'self' 'unsafe-inline'",               // React開発用
    "img-src":     "'self' data: https://image.tmdb.org",  // TMDB画像
    "connect-src": "'self' https://api.themoviedb.org",    // TMDB API
    "font-src":    "'self' data:",
    "object-src":  "'none'",
    "base-uri":    "'self'",
    "form-action": "'self'",
}
```

### 本番環境用CSP（より厳格）
```go
CSPDirectives: map[string]string{
    "default-src": "'self'",
    "script-src":  "'self'",                               // unsafe-inlineを削除
    "style-src":   "'self' 'unsafe-inline'",               // CSSは許可
    "img-src":     "'self' data: https://image.tmdb.org",  // TMDB画像
    "connect-src": "'self' https://api.themoviedb.org",    // TMDB API
    "font-src":    "'self'",
    "object-src":  "'none'",
    "base-uri":    "'self'",
    "form-action": "'self'",
    "upgrade-insecure-requests": "",                       // HTTPSにアップグレード
}
```

## 🧪 テスト結果

### 基本セキュリティヘッダー確認
```bash
curl -I http://localhost:8080/api/movies
```

**確認されたヘッダー**:
```
X-XSS-Protection: 1; mode=block
X-Content-Type-Options: nosniff
X-Frame-Options: DENY
Content-Security-Policy: [設定済み]
Referrer-Policy: strict-origin-when-cross-origin
Permissions-Policy: geolocation=(), microphone=(), camera=()
```

### CORS動作確認
```bash
curl -I -H "Origin: .env.localのFRONTEND_URL" http://localhost:8080/api/movies
```

**確認されたCORSヘッダー**:
```
Access-Control-Allow-Origin: .env.localのFRONTEND_URL
Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
Access-Control-Allow-Headers: Origin, Content-Type, Accept, Authorization, X-Requested-With, X-HTTP-Method-Override
Access-Control-Allow-Credentials: true
Access-Control-Max-Age: 86400
```

### プリフライトリクエスト確認
```bash
curl -I -X OPTIONS -H "Origin: .env.localのFRONTEND_URL" -H "Access-Control-Request-Method: GET" http://localhost:8080/api/movies
```

**結果**: `204 No Content` - プリフライトリクエスト成功

## 🌟 セキュリティ効果

### 1. **XSS攻撃防止**
- `X-XSS-Protection: 1; mode=block`
- ブラウザのXSSフィルターを有効化
- 悪意のあるスクリプト実行をブロック

### 2. **クリックジャッキング防止**
- `X-Frame-Options: DENY`
- iframe内での表示を完全に禁止
- UI偽装攻撃を防止

### 3. **CSRF攻撃軽減**
- オリジン制御による厳格なCORS設定
- 許可されたオリジンからのリクエストのみ受け入れ
- プリフライトリクエストによる事前検証

### 4. **コンテンツ改ざん防止**
- Content Security Policyによるリソース制御
- 許可されたソースからのみリソース読み込み
- インラインスクリプトの制限（本番環境）

### 5. **情報漏洩防止**
- `Referrer-Policy: strict-origin-when-cross-origin`
- 外部サイトへの情報漏洩を最小化
- `Permissions-Policy`によるAPI使用制限

## 🔄 環境別設定

### 開発環境
- `unsafe-inline`、`unsafe-eval`を許可（React Hot Reload対応）
- HTTP通信許可
- 複数のローカルオリジン許可

### 本番環境
- HTTPS強制（HSTS有効化）
- 厳格なCSP設定
- 特定のフロントエンドURLのみ許可

## 🚀 使用方法

### 環境変数設定
```bash
# 開発環境
FRONTEND_URL=http://localhost:3003

# 本番環境
GO_ENV=production
FRONTEND_URL=https://yourdomain.com
```

### サーバー起動時のログ
```
Server starting on http://localhost:8080
Server listening on port :8080
Security middleware enabled with CORS origins: [http://localhost:3003]
```

## 💡 学習事項

### セキュリティベストプラクティス
- **Defense in Depth**: 多層防御の重要性
- **OWASP Top 10**: 主要な脆弱性への対策
- **CORS理解**: プリフライトリクエストの仕組み

### Go言語でのミドルウェア実装
- `http.Handler`インターフェースの活用
- ミドルウェアチェーンの構築
- 設定の外部化と環境分離

### セキュリティヘッダーの効果的な使用
- CSPディレクティブの適切な設定
- 開発環境と本番環境の設定分離
- ブラウザ互換性の考慮

