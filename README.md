# Go Web API Team Project

Go言語を使用したWeb API開発のチームプロジェクトです。

## 📋 プロジェクト概要

**開発期間**: 2025/06/21 〜 2025/07/05 (2週間)  
**チームメンバー**: 2名  
**目標**: Go言語でテーマ性のあるWeb APIを開発し、ドキュメントとデモアプリケーションを作成

[GoLang Web API](https://docs.google.com/document/d/109lAgPlqe-A-gEmL_E1rnRr0w2NHcBNEi8BpxeYxgCw/edit?tab=t.0)  

## ✅ 決定事項

### 技術スタック
- **バックエンド**: Go言語 + net/http標準ライブラリ
- **フロントエンド**: HTML, CSS, JavaScript
- **デプロイ**: ローカル環境のみ

### スケジュール
- **〜6/25**: Go言語学習期間（各自でGo Tour完了）
- **6/25**: プロジェクトテーマ決定・役割分担MTG
- **6/26〜7/5**: 開発期間（実質10日間）

### 開発方針
- **API実装**: 1人1エンドポイント以上担当
- **コードレビュー**: プルリクエストでお互いの実装を学び合う
- **Git戦略**: GitHubフローで開発

## ❓ 未決定事項

### プロジェクト仕様
- [ ] APIのテーマ選定
- [ ] 具体的なエンドポイント設計
- [ ] データ構造定義
- [ ] 詳細な役割分担

### 技術選択
- [ ] 外部ライブラリの使用範囲
- [ ] データ管理方法（メモリ/ファイル/DB）
- [ ] フロントエンドフレームワークの使用可否

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
go mod tidy
```

## 📝 開発フロー

### GitHubフローの基本
1. **Issues作成**: 実装する機能や修正内容をIssueで管理
2. **ブランチ作成**: `feature/機能名` または `fix/修正内容` 形式
3. **実装・コミット**: [コミットメッセージ規約](##コミットメッセージ規約)に従う
4. **プルリクエスト**: レビュー依頼とコードレビュー
5. **マージ**: レビュー完了後にmainブランチへマージ

### ブランチ命名規則
```
feature/api-endpoint-name    # 新機能実装
feature/user-authentication  # ユーザー認証機能
fix/json-response-format     # バグ修正
docs/api-documentation       # ドキュメント更新
```

## 📋 コミットメッセージ規約

以下の記事を参考にした規約を採用します：
- [Gitのコミットメッセージの書き方（2023年ver.）](https://zenn.dev/itosho/articles/git-commit-message-2023)
- [Gitブランチ戦略について](https://qiita.com/ucan-lab/items/371cdbaa2490817a6e2a)

### 基本フォーマット
```
<Type>: <Emoji> #<Issue Number> <Title>
```

### Type（必須）
- **feat**: ユーザー向けの機能の追加や変更
- **fix**: ユーザー向けの不具合の修正
- **docs**: ドキュメントの更新
- **style**: フォーマットなどのスタイルに関する修正
- **refactor**: リファクタリングを目的とした修正
- **test**: テストコードの追加や修正
- **chore**: タスクファイルなどプロダクションに影響のない修正

### Emoji（任意）
- [gitmoji](https://gitmoji.dev)から選択
- Typeをよりカラフルにするため（必須ではない）

### Issue Number（強く推奨）
- そのコミットに紐づくIssue番号を記載
- GitHub上でリンクされ、トラッキングしやすくなる

### Title（必須・日本語でOK）
- 20〜30文字以内が適切

### コミットの粒度
- **1 Commit = 1つの意味あるまとまり**であるべき
- レビュアーがPull Requestを見たときに"ストーリー"が分かることを意識
- **1 Issue、1 Pull Request、1 Commit が理想**

### 例（日本語推奨）
```bash
# 良い例
feat: ✨ #123 ユーザー登録エンドポイントを追加
fix: 🐛 #124 セッションタイムアウトの問題を修正
docs: 📝 #125 API仕様書を更新
feat: #126 映画一覧取得APIを実装
refactor: #127 ハンドラー関数をリファクタリング

# 避けるべき例
update code
fix bug
add feature
とりあえず保存
ログイン機能
```

### Whyについて
- **Whyはコミットメッセージではなく、IssueやPull Requestで説明**
- コミットメッセージにはIssue番号を必ず紐付ける
- Subject は What に寄った書き方でOK

## 🗂️ ディレクトリ構造

```
.
├── README.md            # プロジェクトの概要・使い方・開発ルール
├── devlog/              # 開発ログ・MTG記録
│   ├── 2025-06-21-kickoff.md       # キックオフMTG記録
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
# → デモアプリでAPIが取得できているか確認

# APIドキュメント確認
# docs/api-spec.md をエディタまたはGitHubで閲覧
```

### 📋 成果物の構成

1. **Web API** (`backend/`)
2. **デモアプリケーション** (`frontend/index.html`)  
3. **APIドキュメント** (`docs/api-spec.md`)
4. **開発ログ** (`devlog/`)

## 📊 進捗管理

### Issues・Pull Requests
- **Issues**: 機能実装やバグ報告の単位で作成
- **Projects**: GitHub Projectsでカンバン管理
- **Milestones**: 週次での進捗管理

### コミュニケーション
- **定期的なMTG**: 予めMTG日程を決めておこうかと思いま
- **困った時**: いつでもDiscordで相談

## 📚 学習リソース

### Go言語
- [Go Tour (日本語)](https://go-tour-jp.appspot.com/) - **必須学習教材**
- [RecursionCS Go](https://recursionist.io/learn/languages/go/) - **推奨学習教材**
- [Go公式ドキュメント](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go.html)

### Web API開発
- [net/httpパッケージ](https://pkg.go.dev/net/http)
- [JSONの扱い方](https://pkg.go.dev/encoding/json)

## 🎯 成果物

### 最終的な提出物
- [ ] Go製Web API
- [ ] APIドキュメントWebサイト
- [ ] デモWebアプリケーション
- [ ] README（使用方法・API仕様）
- [ ] 開発ログ（devlog/内のMarkdownファイル）

## 👥 チームメンバー

| Name | GitHub | Role | 担当エンドポイント |
|------|--------|------|-------------------|
| 有堀 | [@takeshi-arihori] | Team Leader | 例: GET /api/movies, POST /api/movies |
|  | [] | Developer | TBD |
|  | [] | Developer | TBD |
|  | [] | Developer | TBD |
|  | [] | Developer | TBD |


---

**Last Updated**: 2025/06/22  
**Next Milestone**: 6/25 テーマ決定MTG
