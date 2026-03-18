# 開発ドキュメント

## 技術スタック

| レイヤー | 技術 |
|---------|------|
| フロントエンド | React + TypeScript（実装予定） |
| バックエンド | Go（標準ライブラリ `net/http`） |
| データベース | PostgreSQL 16 |
| 開発環境 | Docker, Air（ホットリロード） |
| DB可視化 | Adminer（http://localhost:8080） |

## ディレクトリ構成

```
GozuTab/
├── client/             # フロントエンド（React + TypeScript）
├── server/
│   ├── main.go         # エントリーポイント・HTTPハンドラー
│   ├── database/       # DB接続
│   ├── models/         # 構造体定義
│   └── templates/      # HTMLテンプレート
├── docs/               # ドキュメント
└── docker/             # PostgreSQLデータボリューム
```

## DBスキーマ

### users

| カラム | 型 | 備考 |
|-------|-----|------|
| id | SERIAL | PRIMARY KEY |
| name | TEXT | NOT NULL |
| email | TEXT | UNIQUE NOT NULL |
| icon_url | TEXT | プロフィール画像URL |
| created_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP |

### tabs

| カラム | 型 | 備考 |
|-------|-----|------|
| id | SERIAL | PRIMARY KEY |
| user_id | INTEGER | REFERENCES users(id) |
| title | TEXT | NOT NULL |
| artist | TEXT | アーティスト名 |
| content | TEXT | タブ譜データ（JSON形式） |
| audio_url | TEXT | 音源URL |
| status | TEXT | public / private / draft |
| created_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP |
| updated_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP |

## API設計

### ユーザー

| メソッド | パス | 説明 | 状態 |
|---------|------|------|------|
| POST | /users/create | ユーザー作成 | 実装中 |
| GET | /users | ユーザー一覧 | 予定 |
| GET | /users/{id} | ユーザー取得 | 予定 |
| PUT | /users/{id} | ユーザー更新 | 予定 |
| DELETE | /users/{id} | ユーザー削除 | 予定 |

### タブ譜

| メソッド | パス | 説明 | 状態 |
|---------|------|------|------|
| POST | /tabs | タブ譜作成 | 予定 |
| GET | /tabs | タブ譜一覧 | 予定 |
| GET | /tabs/{id} | タブ譜取得 | 予定 |
| PUT | /tabs/{id} | タブ譜更新 | 予定 |
| DELETE | /tabs/{id} | タブ譜削除 | 予定 |

## 開発進捗

- [x] 環境構築（Docker / PostgreSQL / Air）
- [x] DB設計・テーブル作成
- [x] API設計
- [ ] ユーザー作成（POST /users/create）← **現在ここ**
- [ ] ユーザーCRUD（Read / Update / Delete）
- [ ] タブ譜CRUD
- [ ] 認証（JWT）
- [ ] フロントエンド（React + TypeScript）
