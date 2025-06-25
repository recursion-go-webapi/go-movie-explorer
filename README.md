# Go Web API Team Project

映画データベース(TMDB)を活用したGo WebAPIプロジェクト

## 📋 プロジェクト概要

**開発期間**: 2025/06/21 〜 2025/07/06 (2週間)  
**チームメンバー**: 5名  

## ✅ 決定事項

### 技術スタック
- **バックエンド**: Go言語 + net/http標準ライブラリ
- **フロントエンド**: HTML, CSS, JavaScript
- **デプロイ**: ローカル環境のみ

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
- ジャンル別映画取得API (ジャンル別)

### データ構造
- TMDB APIと連携した映画情報データ
- JSON形式でのレスポンス
- Go構造体での型安全なデータ管理

### 役割分担
| 担当者  | 役割             | 詳細                                                 |
| ------- | ---------------- | ---------------------------------------------------- |
| Masato  | ドキュメント作成 | APIのドキュメント作成         |
| Arihori | インフラ・基盤   | ディレクトリ構成・ヘルスチェックAPI・映画一覧取得API |
| 全員    | 映画API作成      | 各自エンドポイント分担・TMDB API連携                 |

### レビュー体制
| プッシュする人 | レビュワー |
| -------------- | ---------- |
| Masato         | Hiroki     |
| Hiroki         | Arihori    |
| Arihori        | Masato     |

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

#### マイルストーン
| 期限       | 担当 | タスク                                     |
| ---------- | ---- | ------------------------------------------ |
| 27日まで | 有堀 | ディレクトリ構成                 |
| 27日まで | 有堀 | API healthチェック + 基本エンドポイント1つ |
| 29日まで | 全員 | Go API 各エンドポイント作成                |

### 開発方針
- **API実装**: 1人1エンドポイント以上担当
- **コードレビュー**: プルリクエストでお互いの実装を学び合う
- **詳細なGitHub運用ルールは[README.github.md](./README.github.md)を参照**

## 🛠️ 開発環境セットアップ

### 必要なツール
- Go
- Git

### セットアップ手順
```bash
# リポジトリのクローン
git clone [このリポジトリのURL]
cd [リポジトリ名]

# Goのバージョン確認
go version

# 依存関係の確認（go.modファイル作成後）
cd backend
go mod tidy
```


## 🗂️ ディレクトリ構造

```
.
├── README.md            # プロジェクトの概要・使い方・開発ルール
├── devlog/              # 開発ログ・MTG記録
│   └── 2025-06-25-theme-decision.md # テーマ決定MTG記録
├── backend/             # Go言語のWeb APIサーバー
│   ├── main.go          # アプリケーションのエントリーポイント・HTTPサーバー起動
│   ├── handlers.go      # HTTPリクエストの処理・エンドポイントのロジック
│   ├── models.go        # データ構造の定義・JSON変換用の構造体
│   └── go.mod           # Go言語の依存関係管理ファイル
├── frontend/            # デモWebアプリケーション
│   ├── index.html       # デモアプリケーション・API動作確認用UI
│   ├── css/             # スタイルシート格納ディレクトリ
│   │   └── app.css      # アプリケーションのスタイリング
│   └── js/              # JavaScript格納ディレクトリ
│       └── app.js       # API通信処理・UI制御・イベントハンドリング
└── docs/                # APIドキュメント
    └── api-spec.md      # API仕様書・エンドポイント詳細（Markdown形式）
```

### 📂 各ディレクトリの役割

#### `backend/`
- **Go言語によるWeb APIサーバー**
- HTTPリクエストを受け取りJSONレスポンスを返す
- ポート8080で起動（例：http://localhost:8080）

#### `frontend/`
- **デモWebアプリケーション**
- APIの動作確認・機能デモンストレーション用
- ブラウザで直接HTMLファイルを開いて使用

#### `docs/`
- **APIドキュメント（Markdown形式）**
- エンドポイント仕様・リクエスト/レスポンス例
- 使用方法・サンプルコード

#### `devlog/`
- **開発過程の記録**
- MTG議事録・技術的な決定事項・振り返り

### 🔧 開発時の使い方

```bash
# バックエンド起動
cd backend
go run .
# → http://localhost:8080 でAPI稼働

# デモアプリ確認
cd frontend
# index.htmlをブラウザで開く  
# → デモアプリでAPI動作確認

# APIドキュメント確認
# docs/api-spec.md をエディタまたはGitHubで閲覧
```

### 📋 成果物の構成

1. **Web API** (`backend/`)
2. **デモアプリケーション** (`frontend/index.html`)  
3. **APIドキュメント** (`docs/api-spec.md`)

## 📊 進捗管理

### Issues・Pull Requests
- **Issues**: 機能実装やバグ報告の単位で作成

## 🎯 成果物

### 最終的な提出物
- [ ] Go製Web API
- [ ] デモWebアプリケーション(フロントエンド・バックエンド)
- [ ] APIドキュメント（Markdown形式）
- [ ] README（使用方法・API仕様）

## 👥 チームメンバー

| Name            | GitHub                                                | Role      | 担当エンドポイント |
| --------------- | ----------------------------------------------------- | --------- | ------------------ |
| Takeshi Arihori | [Takeshi Arihori](https://github.com/takeshi-arihori) | Developer | TBD                |

---

**Last Updated**: 2025/06/25  

