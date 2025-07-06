# Go Movie Explorer

映画データベース(TMDB)を活用したGoバックエンド + Reactフロントエンドアプリケーション

<img width="1720" alt="Screenshot 2025-07-06 at 8 08 24 PM" src="https://github.com/user-attachments/assets/ad38d4b4-ea34-41c5-afeb-f26adb6041e5" />

<img width="1779" alt="Screenshot 2025-07-06 at 8 09 28 PM" src="https://github.com/user-attachments/assets/0a068472-51cd-4516-9325-eb3e784d0199" />

<img width="1776" alt="Screenshot 2025-07-06 at 8 44 58 PM" src="https://github.com/user-attachments/assets/c7520f98-63be-4083-a2e9-966d51b32a80" />


## 📖 概要

Go Movie Explorerは、The Movie Database (TMDB) APIを利用した映画情報検索・閲覧アプリケーションです。

### 主な機能
- 🎬 **映画一覧表示** - 人気映画の一覧表示
- 🔍 **映画検索** - タイトルで映画を検索
- 📝 **映画詳細** - 映画の詳細情報表示
- 🎭 **ジャンル別表示** - ジャンル別の映画一覧
- ⭐ **人気映画ランキング** - 人気映画のランキング表示

### 技術スタック
- **バックエンド**: Go + net/http標準ライブラリ
- **フロントエンド**: React + TypeScript + Vite + Tailwind CSS + shadcn/ui
- **API**: The Movie Database (TMDB) API
- **開発環境**: Docker Compose

---

## 🚀 クイックスタート

### 必要環境
- **Go**: 1.19以上
- **Node.js**: 18以上
- **TMDB API Key**: [TMDB公式サイト](https://www.themoviedb.org/settings/api)から取得

### 1. リポジトリのクローン
```bash
git clone https://github.com/recursion-go-webapi/go-movie-explorer.git
cd go-movie-explorer
```

### 2. TMDB API Keyの設定

#### バックエンド用
```bash
cd backend
cp .env.example .env
```

`.env`ファイルを編集してAPIキーを設定：
```bash
PORT=8080
TMDB_API_KEY=your_tmdb_api_key_here
GO_ENV=development
FRONTEND_URL=http://localhost:3003
```

#### フロントエンド用
```bash
cd frontend
cp .env.example .env # .envを作成
```

`.env`ファイルを編集：
```bash
VITE_API_BASE_URL=http://localhost:8080
```

---

## 💻 ローカル開発環境での起動

### 方法1: 手動起動

#### バックエンド起動
```bash
cd backend

# 依存関係のインストール
go mod tidy

# サーバー起動
go run main.go
```
✅ バックエンド: http://localhost:8080

#### フロントエンド起動
```bash
cd frontend

# 依存関係のインストール
npm install

# 開発サーバー起動
npm run dev
```
✅ フロントエンド: http://localhost:3003

### 方法2: Docker Compose（自動起動）

Docker Composeを使用した詳細な起動方法は [README.docker.md](./README.docker.md) を参照してください。

```bash
# 一括起動
docker compose up -d

# 起動確認
docker compose ps
```

起動後のアクセス先：
- **フロントエンド**: http://localhost:3003
- **バックエンドAPI**: http://localhost:8080
- **Swagger UI**: http://localhost:8081

---

## 🔗 APIエンドポイント

| メソッド | エンドポイント | 説明 |
|---------|---------------|------|
| GET | `/healthz` | ヘルスチェック |
| GET | `/api/movies` | 映画一覧取得 |
| GET | `/api/movie/{id}` | 映画詳細取得 |
| GET | `/api/movies/search` | 映画検索 |
| GET | `/api/movies/popular` | 人気映画ランキング |
| GET | `/api/genres` | ジャンル一覧取得 |
| GET | `/api/movies/genre` | ジャンル別映画取得 |

### API仕様書
- **Swagger UI**: http://localhost:8081 (Docker起動時)
- **OpenAPI仕様**: [docs/openapi.yaml](./docs/openapi.yaml)

---

## 🧪 動作確認

### バックエンドAPI
```bash
# ヘルスチェック
curl http://localhost:8080/healthz

# 映画一覧取得
curl http://localhost:8080/api/movies

# 映画詳細取得（例：Fight Club）
curl http://localhost:8080/api/movie/550

# 映画検索
curl "http://localhost:8080/api/movies/search?query=batman"

# ジャンル一覧取得
curl http://localhost:8080/api/genres

# ジャンル別映画取得（例：Actionジャンル）
curl "http://localhost:8080/api/movies/genre?genre_id=28"

# 人気映画ランキング
curl http://localhost:8080/api/movies/popular

```

### フロントエンド
ブラウザで http://localhost:3003 にアクセスして以下を確認：
- ホームページの表示
- 映画一覧の表示
- 検索機能
- 映画詳細ページ
- 映画カテゴリ一覧
---

## 🗂️ プロジェクト構成

```
go-movie-explorer/
├── README.md                 # このファイル（ローカル起動方法）
├── README.docker.md          # Docker関連の操作方法
├── README.github.md          # GitHub運用ルール
├── compose.yml               # Docker Compose設定
├── backend/                  # Goバックエンド
│   ├── main.go              # メインエントリーポイント
│   ├── handlers/            # APIハンドラー
│   ├── middleware/          # ミドルウェア
│   ├── models/              # データモデル
│   ├── services/            # TMDB APIクライアント
│   └── .env.example         # 環境変数テンプレート
├── frontend/                # Reactフロントエンド
│   ├── src/                 # ソースコード
│   │   ├── pages/           # ページコンポーネント
│   │   ├── components/      # UIコンポーネント
│   │   ├── hooks/           # カスタムフック
│   │   └── api/             # APIクライアント
│   └── .env.example         # 環境変数テンプレート
└── docs/                    # API仕様書
    └── openapi.yaml         # OpenAPI仕様
```

---

## 🛠️ 開発・運用

### 開発フロー
詳細なGitHub運用ルール・ブランチ戦略は [README.github.md](./README.github.md) を参照してください。

### Docker運用
Docker Composeを使用したデプロイ・本番運用については [README.docker.md](./README.docker.md) を参照してください。

### テスト実行
```bash
# バックエンドテスト
cd backend
go test ./...

# フロントエンドテスト
cd frontend
npm test
```

---

## 🆘 トラブルシューティング

### よくある問題

#### 1. TMDB APIキーエラー
```
log.Fatal("TMDB_API_KEYが設定されていません")
```
**解決方法**: `.env`ファイルでTMDB_API_KEYが正しく設定されているか確認

#### 2. ポート番号の競合
```
bind: address already in use
```
**解決方法**: 既に起動中のプロセスを停止するか、別のポートを使用

#### 3. フロントエンドでAPIエラー
```
Failed to fetch
```
**解決方法**: バックエンドが起動しているか確認（http://localhost:8080/healthz）

#### 4. Docker起動エラー
**解決方法**: [README.docker.md](./README.docker.md) のトラブルシューティングセクションを参照

---

## 🤝 コントリビューション

### 👥 チームメンバー

| Name            | GitHub                                                | Role      | 担当エンドポイント |
| --------------- | ----------------------------------------------------- | --------- | ------------------ |
| Takeshi | [Takeshi ](https://github.com/takeshi-arihori) | Developer | TBD                |
| Hiroki | [Hiroki](https://github.com/hiroki-jandararin) | Developer | TBD                |
| Masato | [Masato](https://github.com/iwmstjp) | Developer | TBD                |
| Tomohiko |  [Tomohiko](https://github.com/2017cjx) | Developer | TBD                |

### 開発参加方法
1. このリポジトリをフォーク
2. フィーチャーブランチを作成
3. 変更をコミット
4. プルリクエストを作成

詳細は [README.github.md](./README.github.md) を参照してください。

---

## 🔗 関連リンク

- [TMDB API Documentation](https://developer.themoviedb.org/docs)
- [Go Documentation](https://golang.org/doc/)
- [React Documentation](https://react.dev/)
- [Tailwind CSS](https://tailwindcss.com/)
- [shadcn/ui](https://ui.shadcn.com/)
