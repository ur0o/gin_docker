# gin_docker
## 概要
ginをdocker上で動かす雛形

## ディレクトリ構成
root/
├ docker/
└ src/


## 利用モジュール
- gin
    - Go言語の軽量フレームワーク。
    - https://gin-gonic.com/ja/
    - https://github.com/gin-gonic/gin
- air
    - ホットリロードしてくれるやつ。これがあるとコード書き換えた時にdockerコンテナを立て直す必要がない。
    - https://github.com/cosmtrek/air

## 使い方
- `git clone https://github.com/ur0o/gin_docker && cd gin_docker`
- `docker-compose up`
- `localhost:8000`にアクセス。`{message: hello world}`してくれれば成功。
- main.goの13行目をいじってレスポンスを変えてみる。docker-composeのlogで`main.go has changed.`の後に`building...`とか出てきたらairがホットリロードしてくれてる。

## うまく動かない時
### airがホットリロードしてくれない
dockerコンテナでマウントしてるとairがファイルの変更を検知してくれないことがある。

1. `docker-compose exec api /bin/sh`でapiコンテナの中に入り、viエディタとかでmain.goを編集。airが動いてくれるかを確かめる。動かなかったらなんか設定ミスってそう。動いてたら次に進む。
1. vscodeの拡張のRemote-Containersをインストール。これを使うと、dockerコンテナ内の環境でvscodeのエディタを開ける。
1. vscodeの左側のリモートエクスプローラーを開く。既に`docker-compose up`済みなら、Dev ContainersかOther Containersにgin_dockerのコンテナが2つ表示されている。それぞれapiとnginxなので、apiの方を右クリックなりして、Attach to Containerをクリックする。
1. 開いたvscodeはコンテナ内の環境なので、main.goを編集してairが動くか確認する。
