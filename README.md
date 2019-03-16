# LINE Notity API

## 概要

LINE Notifyアカウントからメッセージ送信するサービスです。

無料ですが、メッセージを送信する数に上限(1000/H)があります。



## フロー

開発者の場合、5のアクセストークンまでスキップできますが、一般にユーザーに使ってもらうサービスの場合、各ユーザーに1〜5までの処理を実行してもらう必要があります。

> API は OAuth2 による認証と、LINEへの通知用のものとで構成されます。全体として以下のような流れになります。
>
> 1. ユーザ：LINEの通知を設定しようとする
> 2. 連携サービス： OAuth2 authorization endpoint へのリダイレクトを行う
> 3. LINE：通知先チャンネルの選択及びユーザへの認可確認を行い、連携サービスにリダイレクト
> 4. 連携サービス：リダイレクト時に付与されたパラメータを用いて OAuth2 token endpoint にアクセスし、アクセストークンを取得する
> 5. 連携サービス：アクセストークンを保存する
> 6. (通知時) 連携サービス：保存したアクセストークンを用いて通知APIをコールする
> 7. (通知設定確認時) 連携サービス：連携状態確認APIをコールしてユーザに連携状態を表示する
> 8. (通知解除時) 連携サービス：連携解除APIをコールする
>
> 以上の流れのうち連携サービスで必要と思われる実装箇所は以下の通りです
>
> - OAuth2 のアクセストークンと、ユーザを関連づけて保存する部分
> - 通知のタイミングで通知APIを呼ぶ部分
> - (連携状態を確認するページがある場合) 連携状態APIをコールして連携状態を表示する部分
> - (連携サービス側から通知解除を行う場合) 通知解除APIをコールする部分
>
> 通知設定確認や通知解除の機能はこちら側のウェブページからも提供しますので、APIでの実装は必須ではありません。



## 開発者向け

開発者登録するとマイページからアクセストークン(以降[access_token])を取得出来ます。

アクセストークンを使ってメッセージを送信します。

公式ドキュメントにあるとおり、POSTでメッセージを送信できます。

```bash
$ curl -X POST -H 'Authorization: Bearer [access_token]' -F 'message=foobar' https://  
notify-api.line.me/api/notify
```

curlコマンドがない場合、ブラウザの拡張機能から送信することもできます。

例：Restlet Client on Google Chrome

[Restlet Client \- REST API Testing \- Chrome ウェブストア](https://chrome.google.com/webstore/detail/restlet-client-rest-api-t/aejoelaoggembcahagimdiliamlcdmfm)



時間を決めて定期的にメッセージを送信するためには、サーバーが必要です。

- [LINE Notifyを使ってごみ収集日を自分にお知らせしてみた \- takapiのブログ](http://takapi86.hatenablog.com/entry/2017/01/02/224244) (Ruby)
- [モテる男は LINE Notify を使って 2タップ で LINE を送信できるらしい \- Qiita](https://qiita.com/RyoAbe/items/dd969935f3267d0ad8c2) (iOS:Firebase)
- [LINE Notify \+ GoogleAppsScript \+ Googleカレンダーで明日の予定を絶対忘れない \- Qiita](https://qiita.com/imajoriri/items/e211547438967827661f) (GoogleAppsScript)



## GC_Alert

LINE Notifyを使ったゴミ収集曜日通知サービスです。

登録したユーザーにゴミ収集曜日を通知します。

現在、作成しているサービスはログイン機能付きですが、この機能無しで動作するバージョンを現在作成中です。

### テーブル構成

| テーブル   |                  |                |
| ---------- | ---------------- | -------------- |
| ユーザー   | ユーザーID       | AUTO INCREMENT |
|            | アクセストークン |                |
|            | 地域ID(※)        | 703825601      |
| ゴミ収集日 | 地域ID           | 703825601      |
|            | 週               | 3(第3週)       |
|            | 曜日             | 4(金曜日)      |
|            | メッセージ       | 燃えるゴミ     |

(*)地域IDは、[住所データCSV \- 仕様【住所\.jp】](http://jusyo.jp/csv/document.html)のデータを使っています。住所検索の予定でしたが郵便番号検索に変更しようとしています。

# LINE Message

## 概要

LINE Notifyと仕組みは同じですが、LINE Notifyアカウントからのメッセージではなくサービスからメッセージを通知することができます。また、LINE Notifyは一方向の通知ですが、Messaging APIは、あなたのサービスとLINEユーザーの双方向コミュニケーションが可能です。

有料です。

ただし、自治体は無料です。



## リンク

- https://developers.line.biz/ja/services/messaging-api/
- [欲しい情報だけLINEで届く！福岡市LINE公式アカウント、開設 : LINE Fukuoka公式ブログ](http://linefukuoka.blog.jp/archives/69828494.html)
- [LINE Messaging API でできることまとめ【送信編】 \- Qiita](https://qiita.com/kakakaori830/items/52e52d969800de61ce28)



## ビジネスアイディア

ゴミ収集曜日通知サービスを作成し、自治体に提案したいと考えています。

- 通知サービスを提供する
  - 対象：自治体、学校(PTA)
  - ツール：LINE Notify
  - 通知種類
    - ゴミ収集日(自治体)
    - 休校(学校：PTA)
    - スケジュール一般
    - 予定
    - リマインダー
- お客様とのコミュニケーション環境構築サービスを提供
  - 対象：企業
  - ツール：LINE Message