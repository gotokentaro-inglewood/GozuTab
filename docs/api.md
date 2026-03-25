# API仕様

## ユーザー管理

| **操作 (CRUD)** | **Method** | **Endpoint** | **内容** | **リクエスト** | **レスポンス** |
| --- | --- | --- | --- | --- | --- |
| **Create** | `POST` | `/users/create` | 新規ユーザー登録 | Body (JSON): `name`, `email` | 201 Created (JSON: User) |
| **Read** | `GET` | `/` | ユーザー・タブ一覧取得 | なし | 200 OK (JSON: `{users, tabs}`) |
| **Update** | `POST` | `/users/update?id={id}` | プロフィール編集 | Query: `id` / Body (JSON): `name`, `icon_url` | 200 OK (JSON: `{message}`) |
| **Delete** | `POST` | `/users/delete?id={id}` | アカウント削除 | Query: `id` | 200 OK (JSON: `{message}`) |

## タブ譜管理

| **操作 (CRUD)** | **Method** | **Endpoint** | **内容** | **リクエスト** | **レスポンス** |
| --- | --- | --- | --- | --- | --- |
| **Create** | `POST` | `/tabs/create` | タブ譜の新規投稿 | Body (JSON): `user_id`, `title`, `content` | 201 Created (JSON: Tab) |
| **Read** | `GET` | `/tabs` | タブ譜一覧取得 | なし | 200 OK (JSON: `[Tab]`) |
| **Update** | `POST` | `/tabs/update?id={id}` | タブ譜の編集 | Query: `id` / Body (JSON): `title`, `content` | 200 OK (JSON: `{message}`) |
| **Delete** | `POST` | `/tabs/delete?id={id}` | タブ譜の削除 | Query: `id` | 200 OK (JSON: `{message}`) |

## スキーマ

### User
```json
{
  "id": 1,
  "name": "Kentaro",
  "email": "kentaro@example.com",
  "icon_url": ""
}
```

### Tab
```json
{
  "id": 1,
  "user_id": 1,
  "title": "Wonderwall",
  "artist": "",
  "content": "Capo 2\nEm G D A",
  "audio_url": "",
  "status": ""
}
```
