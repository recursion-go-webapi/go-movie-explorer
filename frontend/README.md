# Go Movie Explorer - Frontend

🎬 映画探索アプリケーションのフロントエンド

## 技術スタック

- **React 19** - UIフレームワーク
- **TypeScript** - 型安全性
- **Vite** - 高速ビルドツール
- **Tailwind CSS 4.0** - スタイリング
- **shadcn/ui** - UIコンポーネント

## 主要機能

- 🎬 映画一覧表示
- 🔍 映画検索機能
- 📄 ページネーション
- 📱 レスポンシブデザイン
- ⚡ ローディング・エラーハンドリング

## セットアップ

### 1. 依存関係のインストール

```bash
npm install
```

### 2. 環境変数の設定

`.env.example`をコピーして`.env`ファイルを作成：

```bash
cp .env.example .env
```

`.env`ファイルの内容を必要に応じて編集：

```env
VITE_API_BASE_URL=http://localhost:8080
```

### 3. 開発サーバー起動

```bash
npm run dev
```

アプリケーションは http://localhost:3003 で起動します。

## ビルド

本番用ビルドの作成：

```bash
npm run build
```

ビルド結果のプレビュー：

```bash
npm run preview
```

## プロジェクト構成

```
src/
├── api/           # APIクライアント
├── components/    # Reactコンポーネント
│   └── ui/       # shadcn/uiコンポーネント
├── hooks/        # カスタムフック
├── lib/          # ユーティリティ関数
└── types/        # TypeScript型定義
```

## バックエンドとの連携

このフロントエンドはGo Movie Explorer APIと連携します：

- APIベースURL: `http://localhost:8080`
- フロントエンド: `http://localhost:3003`

バックエンドが起動していることを確認してください。

## 開発ガイド

### 新しいコンポーネントの追加

shadcn/uiコンポーネントを追加：

```bash
npx shadcn@latest add [component-name]
```

### リント実行

```bash
npm run lint
```
