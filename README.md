# VVGo

GCP の Cloud Run で実行する Discord Bot コンテナ．

# Deploy

Github Actions でデプロイ．トリガーは main への push．

# ローカル実行

環境変数 DEV_TOKEN をセットし，以下を実行

```shell
$ go run bot.go dev
```

※`dev`オプションをつけない場合，`TOKEN`（本番 Bot のトークン）を読む
