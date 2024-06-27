# goバージョン
FROM golang:1.19.1-alpine
# アップデートとgitのインストール
RUN apk update && apk add git
# volumのマウント先のディレクトリの作成
RUN mkdir 
# ワーキングディレクトリの設定
WORKDIR
# ホストのファイルをコンテナの作業ディレクトリに以降
ADD . 
# パッケージのインポート
RUN go get -u golang.org/x/tools/cmd/goimports
