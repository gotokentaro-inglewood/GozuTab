# 1. ベースイメージの選択
# golang:<バージョン>-alpine を選ぶのが実務のスタンダード
FROM golang:1.26-alpine3.23

# 2. コンテナ内の作業ディレクトリを決める
WORKDIR /app

# 3. ライブリロード（コードを書き換えたら自動反映）ツールを入れる
# これを入れないと、コードを変えるたびにコンテナを作り直す必要があり大変です
RUN go install github.com/air-verse/air@v1.64.5

# 4. 依存関係ファイルをコピー
COPY go.mod go.sum* ./
RUN go mod download

# 5. 全てのファイルをコピー
COPY . .

# 6. airコマンド（自動再起動モード）で起動
CMD ["air", "-c", ".air.toml"]