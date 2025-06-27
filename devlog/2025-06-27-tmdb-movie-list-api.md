# 2025-06-27 - TMDB映画一覧API・共通エラーハンドラ・ログ出力 実装記録

## 実装内容

### 1. TMDB映画一覧APIの実装
- `/api/movies` エンドポイントでTMDBの映画一覧を取得し返却
- サービス層（services/tmdb.go）でTMDB API `/discover/movie` を呼び出し
- models/movies.goでレスポンス構造体を定義
- handlers/movies.goでサービス層を呼び出してレスポンス返却

### 2. 共通エラーハンドラの導入
- middleware/error_handler.goでAppHandler型・ErrorHandlerラッパーを実装
- すべてのハンドラーで発生したエラーをlogs/server.logに記録
- main.goでErrorHandlerをルーティングに適用

### 3. ログディレクトリの自動作成
- サーバー起動時にlogs/ディレクトリがなければ自動作成
- .gitignoreでlogs/*.logを除外

### 4. 動作確認
- `curl http://localhost:8080/api/movies` でTMDBから映画一覧が取得できることをcurlで確認
- エラー発生時はlogs/server.logに記録されることを確認
