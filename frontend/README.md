# frontend

web-tech-dojo のフロントエンド(React)です。

## 環境構築

dockerで起動する場合は，プロジェクトルートのREADMEを参照してください。

```
$ npm i
$ npm run dev
```

## ディレクトリ構成と実装方針

本プロジェクトは初期開発のスピードとシンプルさを重視し、軽量なディレクトリ構成を採用しています。

```
frontend
├── public
├── src
|   ├── components  // 再利用可能なUIコンポーネントを管理するディレクトリ
|   ├── hooks       // フックを管理するディレクトリ
|   ├── lib         // 外部ライブラリやカスタムライブラリを管理するディレクトリ
|   ├── pages       // ページ単位のコンポーネントを管理するディレクトリ
|   └── types       // 型を管理するディレクトリ
└── ...
```

## 環境変数の一覧

.env.localを作成して以下の環境変数を設定してください

| 変数名                            | 役割                                                        | デフォルト値 | DEV 環境での値        |
| --------------------------------- | ----------------------------------------------------------- | ------------ | --------------------- |
| VITE_FIREBASE_API_KEY             | FirebaseプロジェクトのAPIキー                               |              |                       |
| VITE_FIREBASE_AUTH_DOMAIN         | Firebase Authenticationで使用するドメイン                   |              |                       |
| VITE_FIREBASE_PROJECT_ID          | FirebaseプロジェクトID                                      |              |                       |
| VITE_FIREBASE_STORAGE_BUCKET      | Firebase Cloud Storageのバケット名                          |              |                       |
| VITE_FIREBASE_MESSAGING_SENDER_ID | Firebase Cloud Messagingで使用する送信者ID                  |              |                       |
| VITE_FIREBASE_APP_ID              | FirebaseプロジェクトのアプリケーションID                    |              |                       |
| VITE_FIREBASE_MEASUREMENT_ID      | Google Analyticsの計測ID                                    |              |                       |
| VITE_BACKEND_URL                  | アプリケーションがアクセスするバックエンドURL               |              | http://localhost:8080 |
| VITE_FORM_URL                     | Google Formにリクエストを送信するためのURL                  |              |                       |
| VITE_FORM_NAME                    | Google Formでユーザー名を指定するフィールド名               |              |                       |
| VITE_FORM_EMAIL                   | Google Formでユーザーのメールアドレスを指定するフィールド名 |              |                       |
| VITE_FORM_CONTENT                 | Google Formで問い合わせ内容を指定するフィールド名           |              |                       |

環境変数をbase64エンコードしてsecretsに持たせるときのコマンド（改行を消してsecretsに登録）

```
# windows
知らん

# ubuntu
$ base64 -i .env.local | xsel --clipboard --input

# mac
$ base64 -i .env.local | pbcopy
```

## 開発方法

UIテンプレートとして [shadcn/ui](https://ui.shadcn.com/) を利用しています。

以下のように，CLIでコンポーネントの追加ができます。

```
$ npx shadcn@latest add button
```
