# Todo
- ドメインの対応
- httpsの対応
- `/`のページやギルドに入っていなかったときに表示するページの有効or無効化の設定追加
- ユーザーログの保存回数削減

---

## ドメインの対応
- 前提としてドメインの対応
- 出来そう（気力があれば）ドメインレベルの認証機能を追加
- ドメインレベルの認証に必要なコンフィグの追加
```yaml
#通常の場所
    host: xx...
    port: 8080
    domain: example.com
    ...
```
```yaml
#ドメインレベルに伴った
    server:
        domain:
            enable: true
            addr: hide.example.com

    domain_privart:
        - hide.example.com
```

---

## httpsの対応
- サーバースタートは`startTLS`を使用する
- httpsに関するコンフィグの追加
```yaml
#Example
    https:
        enable: true
        port: 443
        cert: /etc/ssl/cert.pem
        key: /etc/ssl/key.pem
```
- httpサーバーも同時に指定ポートで実行しhttpsが有効ならhttpsにリダイレクト

---

## /やギルド云々
- /やギルドに入っていなかった場合に表示するページをホストするサーバーを設定できるようにする。
- 設定がなかったらstringを返すのみ
- これに伴ったコンフィグの追加
```yaml
extarnal_page:
    slashaccess: 
        enable: true
        addr: 127.0.0.1:9000
    NotGuilds: 
        enable: true
        addr: 127.0.0.1:9001
    NotLogin: 
        enable: true
        addr: 127.0.0.1:9002
```

---

## ユーザーログの保存回数削減
- 現在はログインするたびに記録するがファイル容量が増えるのが問題
- 初回のみ[NEW]が先頭につき、それ以降はNEWがついていない場合は記録・ついていたら２つ目まで記録しあとは削除して追加（最新のみ）
```log
[new] ~~~~~~
~~~~~~~
```
のように１ユーザー２つのみログを残す

## redisの完全対応
- 対応する。