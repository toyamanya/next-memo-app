# ベースイメージ
FROM golang:1.20

# 作業ディレクトリの設定
WORKDIR /app

# ソースコードのコピー
COPY . .

# 必要なモジュールをダウンロード
RUN go mod init memo-app && go mod tidy

