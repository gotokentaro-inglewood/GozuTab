# GozuTab

ギター・ベースなどのタブ譜を共有・編集できるWebアプリです。

## セットアップ

```bash
docker compose up -d
```

## アクセス

| サービス | URL |
|---------|-----|
| アプリ | http://localhost:8000 |
| DB管理UI (Adminer) | http://localhost:8080 |

## ディレクトリ構成

```
GozuTab/
├── client/     # フロントエンド（React + TypeScript）
├── server/     # バックエンド（Go）
├── docs/       # ドキュメント
└── docker/     # Docker関連ファイル
```

## 開発ドキュメント

[docs/development.md](docs/development.md) を参照してください。
