# TMDB API 機能一覧

このドキュメントでは、The Movie Database (TMDB) APIが提供している全機能について説明します。

## 概要

TMDB APIは映画、TV番組、人物に関する豊富な情報を提供するREST APIです。バージョン3を使用しており、映画・TV・俳優・画像APIの決定版となっています。

## API機能一覧

### **ACCOUNT（アカウント）**
- **Details**（アカウント詳細）- ユーザーアカウント情報の取得
- **Add Favorite**（お気に入り追加）- 映画・TV番組をお気に入りに追加
- **Add To Watchlist**（ウォッチリスト追加）- 映画・TV番組をウォッチリストに追加
- **Favorite Movies**（お気に入り映画）- お気に入り映画一覧の取得
- **Favorite TV**（お気に入りTV）- お気に入りTV番組一覧の取得
- **Lists**（リスト）- ユーザーが作成したリストの取得
- **Rated Movies**（評価済み映画）- 評価済み映画一覧の取得
- **Rated TV**（評価済みTV）- 評価済みTV番組一覧の取得
- **Rated TV Episodes**（評価済みTVエピソード）- 評価済みTVエピソード一覧の取得
- **Watchlist Movies**（ウォッチリスト映画）- ウォッチリスト映画一覧の取得
- **Watchlist TV**（ウォッチリストTV）- ウォッチリストTV番組一覧の取得

### **AUTHENTICATION（認証）**
- **Create Guest Session**（ゲストセッション作成）- ゲストセッションの作成
- **Create Request Token**（リクエストトークン作成）- 認証用リクエストトークンの作成
- **Create Session**（セッション作成）- ユーザーセッションの作成
- **Create Session (from v4 token)**（v4トークンからセッション作成）- v4トークンを使用したセッション作成
- **Create Session (with login)**（ログイン付きセッション作成）- ログイン情報を使用したセッション作成
- **Delete Session**（セッション削除）- セッションの削除
- **Validate Key**（キー検証）- APIキーの検証

### **CERTIFICATIONS（認定・レーティング）**
- **Movie Certifications**（映画認定）- 映画のレーティング情報取得
- **TV Certifications**（TV認定）- TV番組のレーティング情報取得

### **CHANGES（変更履歴）**
- **Movie List**（映画リスト）- 映画データの変更履歴
- **People List**（人物リスト）- 人物データの変更履歴
- **TV List**（TVリスト）- TV番組データの変更履歴

### **COLLECTIONS（コレクション）**
- **Details**（詳細）- コレクションの詳細情報
- **Images**（画像）- コレクションの画像（ポスター、背景画像）
- **Translations**（翻訳）- コレクションの翻訳情報

### **COMPANIES（会社・制作会社）**
- **Details**（詳細）- 制作会社の詳細情報
- **Alternative Names**（代替名）- 制作会社の代替名
- **Images**（画像）- 制作会社のロゴ画像

### **CONFIGURATION（設定）**
- **Details**（詳細）- API設定の詳細（画像サイズ、ベースURL等）
- **Countries**（国家）- 利用可能な国家一覧
- **Jobs**（職種）- 映画・TV業界の職種一覧
- **Languages**（言語）- 利用可能な言語一覧
- **Primary Translations**（主要翻訳）- 主要翻訳言語一覧
- **Timezones**（タイムゾーン）- 利用可能なタイムゾーン一覧

### **CREDITS（クレジット）**
- **Details**（詳細）- クレジット情報の詳細

### **DISCOVER（発見・検索）**
- **Movie**（映画）- 30以上のフィルターと並び替えオプションで映画を検索
- **TV**（テレビ）- 30以上のフィルターと並び替えオプションでTV番組を検索

### **FIND（検索）**
- **Find By ID**（ID検索）- 外部ID（IMDb、TVDb等）による検索

### **GENRES（ジャンル）**
- **Movie List**（映画リスト）- 映画の公式ジャンル一覧
- **TV List**（TVリスト）- TV番組の公式ジャンル一覧

### **GUEST SESSIONS（ゲストセッション）**
- **Rated Movies**（評価済み映画）- ゲストセッションで評価した映画
- **Rated TV**（評価済みTV）- ゲストセッションで評価したTV番組
- **Rated TV Episodes**（評価済みTVエピソード）- ゲストセッションで評価したTVエピソード

### **KEYWORDS（キーワード）**
- **Details**（詳細）- キーワードの詳細情報
- **Movies**（映画）- キーワードに関連する映画

### **LISTS（リスト）**
- **Add Movie**（映画追加）- リストに映画を追加
- **Check Item Status**（アイテムステータス確認）- アイテムのリスト登録状況確認
- **Clear**（クリア）- リストのクリア
- **Create**（作成）- リストの作成
- **Delete**（削除）- リストの削除
- **Details**（詳細）- リストの詳細情報
- **Remove Movie**（映画削除）- リストから映画を削除

### **MOVIE LISTS（映画リスト）**
- **Now Playing**（現在上映中）- 現在上映中の映画一覧
- **Popular**（人気）- 人気映画一覧
- **Top Rated**（高評価）- 高評価映画一覧
- **Upcoming**（公開予定）- 公開予定映画一覧

### **MOVIES（映画）**
- **Details**（詳細）- 映画の詳細情報
- **Account States**（アカウント状態）- ユーザーの映画に対する状態
- **Alternative Titles**（代替タイトル）- 映画の代替タイトル
- **Changes**（変更履歴）- 映画データの変更履歴
- **Credits**（クレジット）- 映画のキャスト・クルー情報
- **External IDs**（外部ID）- 映画の外部ID（IMDb、Facebook等）
- **Images**（画像）- 映画の画像（ポスター、背景画像、スチール）
- **Keywords**（キーワード）- 映画に関連するキーワード
- **Latest**（最新）- 最新追加映画
- **Lists**（リスト）- 映画が含まれるリスト
- **Recommendations**（おすすめ）- 映画のおすすめ作品
- **Release Dates**（公開日）- 映画の公開日情報
- **Reviews**（レビュー）- 映画のレビュー
- **Similar**（類似）- 類似映画
- **Translations**（翻訳）- 映画の翻訳情報
- **Videos**（動画）- 映画の動画（トレーラー、クリップ等）
- **Watch Providers**（配信プロバイダー）- 映画の配信プロバイダー情報
- **Add Rating**（評価追加）- 映画への評価追加
- **Delete Rating**（評価削除）- 映画の評価削除

### **NETWORKS（ネットワーク）**
- **Details**（詳細）- ネットワークの詳細情報
- **Alternative Names**（代替名）- ネットワークの代替名
- **Images**（画像）- ネットワークのロゴ画像

### **PEOPLE LISTS（人物リスト）**
- **Popular**（人気）- 人気の人物一覧

### **PEOPLE（人物）**
- **Details**（詳細）- 人物の詳細情報
- **Changes**（変更履歴）- 人物データの変更履歴
- **Combined Credits**（合計クレジット）- 人物の全クレジット情報
- **External IDs**（外部ID）- 人物の外部ID
- **Images**（画像）- 人物の画像
- **Latest**（最新）- 最新追加人物
- **Movie Credits**（映画クレジット）- 人物の映画クレジット
- **TV Credits**（TVクレジット）- 人物のTVクレジット
- **Tagged Images**（タグ付き画像）- 人物がタグ付けされた画像
- **Translations**（翻訳）- 人物の翻訳情報

### **REVIEWS（レビュー）**
- **Details**（詳細）- レビューの詳細情報

### **SEARCH（検索）**
- **Collection**（コレクション）- コレクション検索
- **Company**（会社）- 制作会社検索
- **Keyword**（キーワード）- キーワード検索
- **Movie**（映画）- 映画検索
- **Multi**（マルチ）- 複数メディアタイプの検索
- **Person**（人物）- 人物検索
- **TV**（テレビ）- TV番組検索

### **TRENDING（トレンド）**
- **All**（全て）- 全メディアタイプのトレンド
- **Movies**（映画）- 映画のトレンド
- **People**（人物）- 人物のトレンド
- **TV**（テレビ）- TV番組のトレンド

### **TV SERIES LISTS（TVシリーズリスト）**
- **Airing Today**（本日放送）- 本日放送のTV番組一覧
- **On The Air**（放送中）- 現在放送中のTV番組一覧
- **Popular**（人気）- 人気TV番組一覧
- **Top Rated**（高評価）- 高評価TV番組一覧

### **TV SERIES（TVシリーズ）**
- **Details**（詳細）- TV番組の詳細情報
- **Account States**（アカウント状態）- ユーザーのTV番組に対する状態
- **Aggregate Credits**（総合クレジット）- TV番組の総合クレジット情報
- **Alternative Titles**（代替タイトル）- TV番組の代替タイトル
- **Changes**（変更履歴）- TV番組データの変更履歴
- **Content Ratings**（コンテンツレーティング）- TV番組のコンテンツレーティング
- **Credits**（クレジット）- TV番組のキャスト・クルー情報
- **Episode Groups**（エピソードグループ）- TV番組のエピソードグループ
- **External IDs**（外部ID）- TV番組の外部ID
- **Images**（画像）- TV番組の画像
- **Keywords**（キーワード）- TV番組に関連するキーワード
- **Latest**（最新）- 最新追加TV番組
- **Lists**（リスト）- TV番組が含まれるリスト
- **Recommendations**（おすすめ）- TV番組のおすすめ作品
- **Reviews**（レビュー）- TV番組のレビュー
- **Screened Theatrically**（劇場上映）- 劇場上映されたTV番組情報
- **Similar**（類似）- 類似TV番組
- **Translations**（翻訳）- TV番組の翻訳情報
- **Videos**（動画）- TV番組の動画
- **Watch Providers**（配信プロバイダー）- TV番組の配信プロバイダー情報
- **Add Rating**（評価追加）- TV番組への評価追加
- **Delete Rating**（評価削除）- TV番組の評価削除

### **TV SEASONS（TVシーズン）**
- **Details**（詳細）- TVシーズンの詳細情報
- **Account States**（アカウント状態）- ユーザーのTVシーズンに対する状態
- **Aggregate Credits**（総合クレジット）- TVシーズンの総合クレジット情報
- **Changes**（変更履歴）- TVシーズンデータの変更履歴
- **Credits**（クレジット）- TVシーズンのキャスト・クルー情報
- **External IDs**（外部ID）- TVシーズンの外部ID
- **Images**（画像）- TVシーズンの画像
- **Translations**（翻訳）- TVシーズンの翻訳情報
- **Videos**（動画）- TVシーズンの動画
- **Watch Providers**（配信プロバイダー）- TVシーズンの配信プロバイダー情報

### **TV EPISODES（TVエピソード）**
- **Details**（詳細）- TVエピソードの詳細情報
- **Account States**（アカウント状態）- ユーザーのTVエピソードに対する状態
- **Changes**（変更履歴）- TVエピソードデータの変更履歴
- **Credits**（クレジット）- TVエピソードのキャスト・クルー情報
- **External IDs**（外部ID）- TVエピソードの外部ID
- **Images**（画像）- TVエピソードの画像
- **Translations**（翻訳）- TVエピソードの翻訳情報
- **Videos**（動画）- TVエピソードの動画
- **Add Rating**（評価追加）- TVエピソードへの評価追加
- **Delete Rating**（評価削除）- TVエピソードの評価削除

### **TV EPISODE GROUPS（TVエピソードグループ）**
- **Details**（詳細）- TVエピソードグループの詳細情報

### **WATCH PROVIDERS（配信プロバイダー）**
- **Available Regions**（利用可能地域）- 配信プロバイダーが利用可能な地域
- **Movie Providers**（映画プロバイダー）- 映画の配信プロバイダー一覧
- **TV Providers**（TVプロバイダー）- TV番組の配信プロバイダー一覧

## 主要な使用例

### 映画情報の取得
```
/movie/{movie_id} - 映画の詳細情報
/movie/{movie_id}/images - 映画の画像
/movie/{movie_id}/videos - 映画の動画（トレーラー等）
/movie/{movie_id}/credits - 映画のキャスト・クルー情報
```

### 検索機能
```
/search/movie - 映画検索
/search/tv - TV番組検索
/search/person - 人物検索
/search/multi - 複数メディアタイプの検索
```

### 人気・トレンド情報
```
/movie/popular - 人気映画
/tv/popular - 人気TV番組
/trending/all/day - 1日のトレンド（全メディア）
/trending/movie/week - 1週間の映画トレンド
```

### 配信情報
```
/movie/{movie_id}/watch/providers - 映画の配信プロバイダー
/tv/{tv_id}/watch/providers - TV番組の配信プロバイダー
```

## 注意事項

- APIキーが必要です
- レート制限があります
- 一部の機能（評価、リスト作成等）にはユーザー認証が必要です
- 画像URLは設定APIから取得したベースURLと組み合わせて使用します

## 参考リンク

- [TMDB API Documentation](https://developer.themoviedb.org/reference/intro/getting-started)
- [TMDB API Support Forum](https://www.themoviedb.org/talk/category/5047958519c29526b50017d6)
- [画像の使用方法](https://developer.themoviedb.org/docs/image-basics)
