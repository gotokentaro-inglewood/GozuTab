# Architecture Overview

このドキュメントでは、本プロジェクトの技術的な全体像と設計方針について説明します。

## 1. システム概要（Bird's Eye View）
ユーザーがブラウザでタブ譜を編集し、GoのAPIを通じてデータベースに保存・閲覧できる構成です。

## 2. 技術スタック
- **Frontend**: TypeScript + Next.js (React)
- **Backend**: Go (Gin or Echo)
- **Database**: PostgreSQL (Dockerで実行)

## 3. ディレクトリ構造（Code Map）
プロジェクトの主要な構成要素とその役割です。

- `client/`: フロントエンド。UIコンポーネントとタブ譜描画ロジック。
- `server/`: バックエンド。APIエンドポイントとDB操作。
- `docs/`: 設計書、ER図、要件定義などのドキュメント類。
- `docker/`: データベース起動用の設定ファイル。

## 4. データフロー（Data Flow）
1. ユーザーがエディタで譜面を編集（JSON形式）。
2. `POST /tabs` でバックエンドに送信。
3. Goがバリデーションを行い、PostgreSQLへ保存。

## 5. 設計上の重要な決定（Design Invariants）
- **タブ譜の保存形式**: 柔軟性を保つため、譜面データは正規化せず、JSON形式のテキストとして一括保存する。
- **認証**: MVPフェーズでは基本的なJWT認証を採用する。