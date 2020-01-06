# go_lineBot_translate
日本語を英語に翻訳するLINE_Bot

## 使用方法
[こちら](https://github.com/Hiroya3/LINE_Bot_parrot_return)を参考にしてください。

### ただ、デプロイするコマンドは違います
以下のコマンドを流してください。

```
gcloud functions deploy ReplayEnglish --runtime go111 --set-env-vars APIKEY={APIKey} --trigger-http
```

APIKEYの取得方法は[こちら](https://cloud.google.com/docs/authentication/api-keys?hl=ja)。
