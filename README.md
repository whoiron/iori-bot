# iori-bot

入室メッセージと特定コメントへ反応するbot。
設定は `config.yml` で行い、CLIパラメータでBot TokenとチャンネルIDを渡します。

## compile

```bash
$ git clone git@github.com:whoiron/iori-bot.git
$ dep ensure
$ go build -o bot
```

## run bot

1. discord app bot作成 [My Apps](https://discordapp.com/developers/applications/me)
2. Bot Token( `XXXXXX` )を取得
3. Client ID(例: `1234567890` )をコピー
4. `https://discordapp.com/oauth2/authorize?client_id=1234567890&scope=bot&permissions=0` にアクセスしてサーバーにBotを招待
5. Botがメッセージを送る & 反応させるチャンネルID( `111111` )を取得

```bash
$ ./bot -t XXXXXX -c 111111
Bot is now running.  Press CTRL-C to exit.

```

## config

```yaml
welcome:
  - "Hey %s, Welcome"
  - "%sさん、ようこそ"
keyword:
  request: "ping"
  response:
    - "pong"
```

- welcome
    - 入室したユーザにメッセージをランダムで送ります。 `%s` にユーザへのメンションが入ります。
- keyword
    - ユーザのコメントがrequestと完全一致した場合、responseのメッセージをランダムで送ります。