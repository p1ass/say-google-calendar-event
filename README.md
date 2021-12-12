# say-google-calender-event

## Getting Started

### 1. GCPのプロジェクトを作成する

### 2. Google Calendar APIを有効化する

https://console.cloud.google.com/marketplace/product/google/calendar-json.googleapis.com

### 3. サービスアカウントを発行する

このスクリプトは個人が手元で実行することを想定するので、OAuth2を使った認可は行わない。

代わりに、サービスアカウントを発行し、そのメールアドレスをカレンダーの共有に追加することで、カレンダーの情報にアクセスできるようにする。

https://console.cloud.google.com/iam-admin/serviceaccounts

- サービスアカウント名: 任意
- ロール: 不要

作成したらJSONの鍵を保存する。


### 4. 予定を読み取りたいカレンダーの設定からサービスアカウントのメールアドレスを共有に追加する

Google Calendarにアクセスし、カレンダーの設定内にある「特定のユーザーと共有」からサービスアカウントのメールアドレスを追加する。

その際に、「予定の表示（すべての予定の詳細）」を選ぶと非公開の予定の詳細も取得できる。

また、このタイミングで設定画面から「カレンダーID」を控えておく。
（デフォルトのカレンダーはメールアドレスがカレンダーIDになる）

### 5. 実行

```shell
export GOOGLE_APPLICATION_CREDENTIALS=PATH_TO_SERVICE_ACCOUNT
go run main.go -calendarId xxx@example.com
```
