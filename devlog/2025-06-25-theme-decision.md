# 2025/06/25 - プロジェクトテーマ決定MTG

## 📅 日時・概要
- **日時**: 2025/06/25
- **種別**: プロジェクトテーマ決定・役割分担MTG
- **参加者**: チームメンバー3名

## 🎯 決定事項

### テーマ選定
**映画API (TMDB連携)**
- [The Movie Database (TMDB) API](https://developer.themoviedb.org/docs/getting-started) を使用
- 映画情報の検索・取得・管理機能を実装

### マイルストーン・スケジュール
| 期限       | 担当 | タスク                                     |
| ---------- | ---- | ------------------------------------------ |
| 金曜日まで | Arihori | ディレクトリ構成 + デプロイ                |
| 金曜日まで | Arihori | API healthチェック + 基本エンドポイント1つ(一覧取得) |
| 日曜日まで | 全員 | Go API 各エンドポイント作成                |

## 👥 役割分担

### レビュー体制
| PR作成者 | レビュワー |
| -------------- | ---------- |
| Masato         | Hiroki     |
| Hiroki         | Arihori    |
| Arihori        | Masato     |

### 担当タスク

#### 📚 ドキュメント作成
**担当**: masato
- API仕様書の詳細化
- エンドポイント設計書
- 使用方法・サンプルコード

#### 🏗️ Go ディレクトリ構成・インフラ
**担当**: Arihori
- プロジェクト構造セットアップ ✅
- デプロイ環境構築
- ヘルスチェックAPI実装

#### 🎬 映画API作成
**担当**: 全員（各自エンドポイント分担）
- TMDB APIキー取得（無料アカウント作成）
- 各映画関連エンドポイントの実装

⚠️ **注意**: TMDB ユーザー登録時は全て英語で入力（日本語だとバリデーションエラー起こります💦）

## 🔗 技術選択・外部API

### TMDB API
- **公式ドキュメント**: https://developer.themoviedb.org/docs/getting-started
- **API Reference**: https://developer.themoviedb.org/reference/intro/getting-started
- **API種別**: REST API
- **認証**: API Key
- **料金**: 無料プラン利用

#### 🧪 API実テスト環境
**重要**: [TMDB API Reference](https://developer.themoviedb.org/reference/intro/getting-started)では、各言語（Shell, Node, Ruby, PHP, Python）でのAPI実テストが可能！
- リアルタイムでAPIレスポンスを確認
- コードサンプル自動生成
- 各エンドポイントのパラメータ詳細確認

### 想定エンドポイント例
- 映画一覧取得: Arihori
- 映画詳細取得
- 映画検索
- 人気映画ランキング
- ジャンル別映画取得

## 📋 次回MTG
- **日時**: 月曜夜 20:00 〜
- **種別**: TEAM MTG
- **目的**: 進捗確認・問題解決・方向性調整

## ✅ 本日完了事項
- [x] プロジェクトテーマ決定（映画API）
- [x] 役割分担・レビュー体制確立
- [x] マイルストーン・スケジュール設定
- [x] ディレクトリ構成セットアップ完了

## 🚀 次のアクション
1. **各自**: TMDB APIキー取得
2. **各自**: [API Reference](https://developer.themoviedb.org/reference/intro/getting-started)を参考にAPIテスト実施
3. **有堀**: 金曜までにデプロイ環境構築
4. **masato**: ドキュメント作成(API仕様書)
5. **全員**: 担当エンドポイント設計開始

---

**記録者**: Takeshi Arihori  
**ブランチ**: `feature/setup-project-structure`
