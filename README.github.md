# GitHub 運用ルール

本プロジェクトのGitHub運用ルール・開発フロー・コミット規約をまとめます。

---

## 📝 ブランチ戦略：GitHub Flow
- **mainブランチ**: 常にデプロイ可能な状態を保つ
- **featureブランチ**: mainから作成し、新機能実装後にmainへマージ
- **fixブランチ**: mainから作成し、バグ修正後にmainへマージ
- **シンプルなブランチ構成**でチーム開発に最適

### ブランチ運用ルール
- mainブランチへの直接pushは禁止
- 必ずfeature/fixブランチを作成してPull Requestでマージ
- マージ後はfeature/fixブランチを削除

### ブランチ命名規則
```
feature/api-endpoint-name    # 新機能実装
feature/user-authentication  # ユーザー認証機能
fix/json-response-format     # バグ修正
docs/api-documentation       # ドキュメント更新
```

---

## 🗂️ GitHubフローの基本
1. **Issues作成**: 実装する機能や修正内容をIssueで管理
2. **ブランチ作成**: `feature/機能名` または `fix/修正内容` 形式
3. **実装・コミット**: [コミットメッセージ規約](#コミットメッセージ規約)に従う
4. **プルリクエスト**: レビュー依頼とコードレビュー
5. **マージ**: レビュー完了後にmainブランチへマージ

---

## 📝 コミットメッセージ規約

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
- 変更内容を現在形で記載
- 20〜30文字以内が適切

### コミットの粒度
- **1 Commit = 1つの意味あるまとまり**であるべき
- レビュアーがPull Requestを見たときに"ストーリー"が分かることを意識
- **1 Issue、1 Pull Request、1 Commit が理想**

### 例
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

---

## 🛠️ プルリクエスト・レビュー運用
- Pull Requestは必ずレビュワーをアサイン
- レビュー後、LGTMならマージ
- マージ後は不要なfeature/fixブランチを削除
- Whyや詳細な背景はPR本文で記載

---

## ❓ Issue運用
- 機能追加・バグ修正・ドキュメント更新など全てIssueで管理
- Issueには目的・要件・完了定義を明記
- 優先順位は番号の小さい順
- 日々の作業記録やバグ報告もIssueで管理

---

## 📚 参考
- [Gitのコミットメッセージの書き方（2023年ver.）](https://zenn.dev/itosho/articles/git-commit-message-2023)
- [Gitブランチ戦略について](https://qiita.com/ucan-lab/items/371cdbaa2490817a6e2a) 
