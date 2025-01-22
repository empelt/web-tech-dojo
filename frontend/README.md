# frontend

Web Tech Dojoのフロントエンド(React)です。

## 環境構築

後で書く

## ディレクトリ構成と実装方針

本プロジェクトは初期開発のスピードとシンプルさを重視し、軽量なディレクトリ構成を採用しています。

```
frontend
├── src
|   ├── components  // 再利用可能なUIコンポーネントを管理するディレクトリ
|   ├── libs        // 外部ライブラリやカスタムライブラリを管理するディレクトリ
|   ├── pages       // ページ単位のコンポーネントを管理するディレクトリ
|   └── utils       // ユーティリティ関数やヘルパー関数を管理するディレクトリ
├── .firebaserc
├── .gitignore
├── eslint.config.js
├── firebase.json
├── package-lock.json
├── package.json
├── README.md
├── tsconfig.app.json
├── tsconfig.json
├── tsconfig.node.json
└── vite.config.ts
```

### 環境変数の一覧

.env.localを作成して以下の環境変数を設定してください

| 変数名                            | 役割                                          | デフォルト値 | DEV 環境での値 |
| --------------------------------- | --------------------------------------------- | ------------ | -------------- |
| VITE_FIREBASE_API_KEY             | FirebaseプロジェクトのAPIキー                 |              |                |
| VITE_FIREBASE_AUTH_DOMAIN         | Firebase Authenticationで使用するドメイン     |              |                |
| VITE_FIREBASE_PROJECT_ID          | FirebaseプロジェクトID                        |              |                |
| VITE_FIREBASE_STORAGE_BUCKET      | Firebase Cloud Storageのバケット名            |              |                |
| VITE_FIREBASE_MESSAGING_SENDER_ID | Firebase Cloud Messagingで使用する送信者ID    |              |                |
| VITE_FIREBASE_APP_ID              | FirebaseプロジェクトのアプリケーションID      |              |                |
| VITE_FIREBASE_MEASUREMENT_ID      | Google Analyticsの計測ID                      |              |                |
| VITE_BACKEND_URL                  | アプリケーションがアクセスするバックエンドURL |              |                |
