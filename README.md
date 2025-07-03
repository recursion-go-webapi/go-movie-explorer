# Go Web API Team Project

映画データベース(TMDB)を活用したGo WebAPIプロジェクト

## 📋 プロジェクト概要

**開発期間**: 2025/06/21 〜 2025/07/06 (2週間)  

## ✅ 決定事項

### 技術スタック
- **バックエンド**: Go言語 + net/http標準ライブラリ
- **フロントエンド**: React + TypeScript + Vite + Tailwind CSS 4.0 + shadcn/ui

### プロジェクトテーマ
**映画API (TMDB連携)**
- [The Movie Database (TMDB) API](https://developer.themoviedb.org/docs/getting-started) を使用
- 映画情報の検索・取得・管理機能を実装
- [API Reference & Live Testing](https://developer.themoviedb.org/reference/account-details)

### エンドポイント設計
- 映画一覧取得API (リスト表示)
- 映画詳細取得API (詳細表示)
- 映画検索API (検索)
- 人気映画ランキングAPI (データの集計)
- ジャンル一覧取得API (ジャンル一覧)
- ジャンル別映画取得API (ジャンル別)

### データ構造
- TMDB APIと連携した映画情報データ
- JSON形式でのレスポンス
- Go構造体での型安全なデータ管理


### レビュー体制
| プッシュする人 | レビュワー |
| -------------- | ---------- |
| Masato         | Hiroki     |
| Hiroki         | Takeshi    |
| Takeshi        | Masato     |

### スケジュール
- **〜6/25**: Go言語学習期間 ✅
- **6/25**: プロジェクトテーマ決定・役割分担MTG ✅
- **6/26〜7/6**: 開発期間（実質10日間）

#### 開発スケジュール詳細
**第1週（6/26〜6/29）: エンドポイント実装週**
- 目標：全エンドポイントの基本実装完了
- 各メンバーが担当エンドポイントを実装
- 基本的なJSON レスポンスが返却される状態

**第2週（6/30〜7/6）: 統合・仕上げ週**
- 目標：各タスクの完成・統合テスト
- フロントエンド（デモアプリ）実装
- APIドキュメント作成・整備
- 全体統合テスト・バグ修正


### 開発方針
- **API実装**: 1人1エンドポイント以上担当
- **コードレビュー**: プルリクエストでお互いの実装を学び合う
- **詳細なGitHub運用ルールは[README.github.md](./README.github.md)を参照**

## 🛠️ 開発環境セットアップ

### 必要なツール
- Go (1.19以上)
- Node.js (18以上)
- npm
- Git

### セットアップ手順

#### 1. リポジトリのクローン
```bash
git clone [このリポジトリのURL]
cd go-movie-explorer
```

#### 2. バックエンドセットアップ
```bash
cd backend

# Goのバージョン確認
go version

# 依存関係のインストール
go mod tidy

# .envファイルの作成とAPIキー設定
cp .env.example .env
# .envをエディタで開き、TMDB_API_KEYを設定

# サーバーの起動
go run main.go
# http://localhost:8080 で起動

# ヘルスチェック
curl http://localhost:8080/healthz
```

#### 3. フロントエンドセットアップ
```bash
cd frontend

# 依存関係のインストール
npm install

# 環境変数の設定
cp .env.example .env

# 開発サーバー起動
npm run dev
# http://localhost:3003 で起動
```

### 🚀 起動方法

**バックエンド起動:**
```bash
cd backend && go run main.go
```

**フロントエンド起動:**
```bash
cd frontend && npm run dev
```

- バックエンド: http://localhost:8080
- フロントエンド: http://localhost:3003


## 🗂️ ディレクトリ構造

```
.
├── README.md            # プロジェクトの概要・使い方・開発ルール
├── devlog/              # 開発ログ・日々の作業記録・MTG記録
│   ├── mtg/            # MTG記録
├── backend/             # Go言語のWeb APIサーバー
│   ├── main.go          # アプリケーションのエントリーポイント・HTTPサーバー起動
│   ├── handlers/        # 各APIエンドポイントのハンドラー群
│   │   ├── movies.go    # 映画関連APIハンドラー
│   │   └── movies_test.go # 映画ハンドラーのテスト
│   ├── middleware/      # ミドルウェア（ログ・セキュリティ・エラーハンドリング）
│   │   ├── error.go     # エラーハンドリングミドルウェア
│   │   ├── logging.go   # ログミドルウェア
│   │   └── security.go  # セキュリティミドルウェア（CORS・ヘッダー設定）
│   ├── models/          # データ構造体定義
│   │   ├── movies.go    # 映画レスポンス構造体
│   │   └── movies_test.go # 映画モデルのテスト
│   ├── services/        # TMDB APIクライアント等のサービス層
│   │   ├── tmdb.go      # TMDB API呼び出しロジック
│   │   └── tmdb_test.go # TMDB サービスのテスト
│   ├── logs/            # ログファイル出力用ディレクトリ
│   │   └── server.log   # サーバーログファイル
│   ├── .env.example     # 環境変数テンプレート
│   ├── go.mod           # Go言語の依存関係管理ファイル
│   ├── go.sum           # 依存関係の検証用ファイル
│   ├── coverage.html    # テストカバレッジレポート（HTML）
│   ├── coverage.out     # テストカバレッジデータ
│   ├── CLAUDE.md        # 開発時のメモ・設定情報
│   └── CODE_REVIEW.md   # コードレビューガイドライン
├── frontend/            # React フロントエンドアプリケーション
│   ├── src/             # ソースコード
│   │   ├── api/         # APIクライアント
│   │   ├── components/  # Reactコンポーネント
│   │   │   ├── Layout.tsx      # 共通レイアウト
│   │   │   ├── Navigation.tsx  # ナビゲーションバー
│   │   │   └── ui/            # shadcn/uiコンポーネント
│   │   ├── pages/       # ページコンポーネント
│   │   │   ├── HomePage.tsx        # ホームページ
│   │   │   ├── MoviesPage.tsx      # 映画一覧ページ
│   │   │   ├── SearchPage.tsx      # 検索ページ
│   │   │   ├── MovieDetailPage.tsx # 映画詳細ページ
│   │   │   └── GenrePage.tsx       # ジャンル別ページ
│   │   ├── hooks/      # カスタムフック
│   │   ├── lib/        # ユーティリティ関数
│   │   └── types/      # TypeScript型定義
│   ├── public/         # 静的ファイル
│   ├── package.json    # npm依存関係
│   ├── vite.config.ts  # Vite設定
│   ├── components.json # shadcn/ui設定
│   ├── eslint.config.js # ESLint設定
│   ├── tsconfig*.json  # TypeScript設定
│   ├── .env.example    # 環境変数テンプレート
│   ├── .gitignore      # Git除外設定
│   └── README.md       # フロントエンド用README
└── docs/                # APIドキュメント
    └── api-spec.md      # API仕様書・エンドポイント詳細（Markdown形式）
```

### 📂 各ディレクトリの役割

#### `backend/`
- **Go言語によるWeb APIサーバー**
- HTTPリクエストを受け取りJSONレスポンスを返す
- ポート8080で起動（例：http://localhost:8080）

#### `frontend/`
- **React + TypeScript マルチページアプリケーション**
- React Router を使用したSPA（Single Page Application）
- 映画検索・一覧表示・詳細表示機能
- モダンなUIライブラリ（Tailwind CSS + shadcn/ui）
- ポート3003で起動（http://localhost:3003）

**ページ構成:**
- ホームページ (`/`) - 人気映画ランキング・アプリ紹介
- 映画一覧ページ (`/movies`) - 映画一覧・検索・フィルター機能
- 検索ページ (`/search`) - リアルタイム検索機能
- 映画詳細ページ (`/movie/:id`) - 映画詳細情報・関連映画
- ジャンル別ページ (`/genre/:id`) - ジャンル別映画一覧

#### `docs/`
- **APIドキュメント（Markdown形式）**
- エンドポイント仕様・リクエスト/レスポンス例
- 使用方法・サンプルコード

#### `devlog/`
- **開発過程の記録**
- MTG議事録・技術的な決定事項・振り返り


### 📋 成果物の構成

1. **Web API** (`backend/`) - Go言語製RESTful API
2. **マルチページWebアプリケーション** (`frontend/`) - React製SPA
3. **APIドキュメント** (`docs/`) - Markdown形式の仕様書

## 📊 進捗管理

### Issues・Pull Requests
- **Issues**: 機能実装やバグ報告の単位で作成

## 🎯 成果物

### 最終的な提出物
- [x] Go製Web API
- [x] React製フロントエンドアプリケーション
- [ ] APIドキュメント（OpenAPI形式）
- [x] README（使用方法・API仕様）


## 👥 チームメンバー

| Name            | GitHub                                                | Role      | 担当エンドポイント |
| --------------- | ----------------------------------------------------- | --------- | ------------------ |
| Takeshi | [Takeshi ](https://github.com/takeshi-arihori) | Developer | TBD                |
| Hiroki | [Hiroki](https://github.com/hiroki-jandararin) | Developer | TBD                |
| Masato | [Masato](https://github.com/iwmstjp) | Developer | TBD                |
