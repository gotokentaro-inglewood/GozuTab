- ユーザー管理
    
    
    | **操作 (CRUD)** | **Method** | **Endpoint** | **内容** | **リクエスト(Body)** | **レスポンス** |
    | --- | --- | --- | --- | --- | --- |
    | **Create** | `POST` | `/users/create` | 新規ユーザー登録 | `name`, `email` | 302 (Redirect) |
    | **Read** | `GET` | `/` | ユーザー一覧表示 | なし | 200 OK (HTML) |
    | **Update** | `POST` | `/users/update` | プロフィール編集 | `id`, `name`, `icon_url` | 200 OK |
    | **Delete** | `POST` | `/users/delete` | アカウント削除 | `id` | 302 (Redirect) |
- タブ譜管理
    | **操作 (CRUD)** | **Method** | **Endpoint** | **内容** | **リクエスト(Body)** | **レスポンス** |
    | --- | --- | --- | --- | --- | --- |
    | **Create** | `POST` | `/tabs/create` | タブ譜の新規投稿 | `user_id`, `title`, `content` | 302 (Redirect) |
    | **Read** | `GET` | `/tabs` | タブ譜一覧表示 | なし | 200 OK (HTML) |
    | **Update** | `POST` | `/tabs/update` | タブ譜の編集 | `id`, `title`, `content` | 200 OK |
    | **Delete** | `POST` | `/tabs/delete` | タブ譜の削除 | `id` | 302 (Redirect) |