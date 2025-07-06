# Docker Compose 使用ガイド

Go Movie ExplorerをDocker Composeで簡単に起動する方法です。

## 🚀 クイックスタート

### 1. 環境変数の設定

```bash
# TMDB APIキーを環境変数として設定
export TMDB_API_KEY=your_tmdb_api_key_here
```

### 2. 本番環境での起動

```bash
# 本番環境用設定で起動
docker-compose up --build

# バックグラウンドで起動
docker-compose up -d --build

# ログを確認
docker-compose logs -f
```

### 3. 開発環境での起動

```bash
# 開発用設定で起動（環境変数ファイルを使用）
cp backend/.env.example backend/.env
cp frontend/.env.example frontend/.env
# .envファイルを編集してTMDB_API_KEYを設定
GO_ENV=development docker-compose up --build
```

## 📍 アクセス先

- **フロントエンド**: http://localhost:3003
- **バックエンドAPI**: http://localhost:8080/api
- **ヘルスチェック**: http://localhost:8080/healthz

## 🛠️ 便利なコマンド

### サービス管理

```bash
# 特定のサービスのみ起動
docker-compose up backend
docker-compose up frontend

# サービス停止
docker-compose stop

# サービス削除（ボリューム保持）
docker-compose down

# サービス削除（ボリューム含む）
docker-compose down -v

# 再ビルド
docker-compose build --no-cache
```

### ログとデバッグ

```bash
# 全サービスのログ
docker-compose logs

# 特定のサービスのログ
docker-compose logs backend
docker-compose logs frontend

# リアルタイムログ
docker-compose logs -f backend

# コンテナに接続
docker-compose exec backend sh
docker-compose exec frontend sh
```

### データベース・ボリューム管理

```bash
# ボリューム一覧
docker volume ls

# ボリューム詳細
docker volume inspect go-movie-explorer_backend_logs

# ボリューム削除
docker volume rm go-movie-explorer_backend_logs
```

## ⚙️ 設定詳細

### 環境変数

| 変数名              | 説明                | デフォルト値                |
| ------------------- | ------------------- | --------------------------- |
| `TMDB_API_KEY`      | TMDB APIキー        | **必須**                    |
| `VITE_API_BASE_URL` | バックエンドAPI URL | `http://localhost:8080/api` |
| `VITE_APP_TITLE`    | アプリケーション名  | `Go Movie Explorer`         |

### ポート設定

| サービス | 内部ポート | 外部ポート | 説明                  |
| -------- | ---------- | ---------- | --------------------- |
| backend  | 8080       | 8080       | Go APIサーバー        |
| frontend | 80         | 3003       | Nginx静的ファイル配信 |

### ヘルスチェック

- **バックエンド**: `/healthz` エンドポイントで TMDB API 接続確認
- **フロントエンド**: `/health` エンドポイントで Nginx 稼働確認

## 🐛 トラブルシューティング

### よくある問題

1. **TMDB_API_KEYが設定されていない**
   ```bash
   # 環境変数を確認
   echo $TMDB_API_KEY
   
   # 設定
   export TMDB_API_KEY=your_api_key
   ```

2. **ポートが既に使用されている**
   ```bash
   # ポート使用状況を確認
   lsof -i :8080
   lsof -i :3003
   
   # docker-compose.ymlでポートを変更
   # "3004:80" など
   ```

3. **ビルドエラー**
   ```bash
   # キャッシュを無効化して再ビルド
   docker-compose build --no-cache
   
   # イメージを削除して再作成
   docker-compose down --rmi all
   docker-compose up --build
   ```

4. **CORS エラー**
   ```bash
   # FRONTEND_URLが正しく設定されているか確認
   docker-compose logs backend | grep CORS
   ```

### パフォーマンス最適化

```bash
# 不要なイメージ・コンテナを削除
docker system prune -a

# ビルドキャッシュを活用
docker-compose build

# リソース使用量を確認
docker-compose top
docker stats
```

## 📝 開発時のヒント

### ホットリロード

開発時はDocker Composeではなく、直接開発サーバーを使用することを推奨：

```bash
# バックエンド
cd backend
go run main.go

# フロントエンド  
cd frontend
npm run dev
```

### デバッグ用設定

`docker-compose.dev.yml` を使用：
- ソースコードのマウント
- 開発用環境変数
- 詳細ログ出力

### 開発・本番環境の差異

#### 開発環境 (`GO_ENV=development`)
- **Dockerfile**: `backend/Dockerfile` を使用（マルチステージビルド）
- **環境変数**: `backend/.env` と `frontend/.env` ファイルを使用
- **ボリューム**: ログボリュームのみ（コンテナ内でビルド）
- **セキュリティ**: 緩い設定（localhost接続許可）
- **ログ**: 詳細ログ出力
- **CSP**: `http://localhost:8080` への接続許可

#### 本番環境 (`GO_ENV=production`)
- **Dockerfile**: `backend/Dockerfile` を使用（マルチステージビルド）
- **環境変数**: 環境変数を直接設定（`.env`ファイル併用可能）
- **ボリューム**: ログボリュームのみ
- **セキュリティ**: 厳格な設定（HTTPS強制、厳格なCSP）
- **ログ**: 最適化されたログ出力
- **CSP**: HTTPS接続のみ許可

#### 環境変数設定方法

**開発環境**:
```bash
# .envファイルを作成
cp backend/.env.example backend/.env
cp frontend/.env.example frontend/.env
# .envファイルを編集してAPIキーを設定
GO_ENV=development docker-compose up -d
```

**本番環境**:
```bash
# 環境変数を直接設定
export TMDB_API_KEY=your_production_api_key
export FRONTEND_URL=https://your-domain.com
export GO_ENV=production
docker-compose up -d
```
