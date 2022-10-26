# VVGo

GCP の Cloud Run で実行する Discord Bot コンテナ．

# Deploy

Github Actions でデプロイ．トリガーは main への push．

# ローカル実行

.env に`TOKEN`, `DEV_TOKEN`, `VOICEVOX_KEY`を設定する

```shell
$ go run bot.go dev
```

※`dev`オプションをつけない場合，`TOKEN`（本番 Bot のトークン）を読む
