# backend

Web Tech Dojo のバックエンド(Echo)です。

## 環境構築

docker で起動する場合は，プロジェクトルートの README を参照してください。

```
$ go build
$ go run .
```

## ディレクトリ構成と実装方針

本プロジェクトは Clean Architecture を簡略化したアーキテクチャで実装を行います。

```
backend
├── handlers          // HTTPリクエストを処理し、サービスやロジックに委譲する
|   ├── mock/
|   ├── ...
|   └── interface.go
├── infrastructures   // 外部リソースや環境設定を管理する
|   ├── mock/
|   ├── ...
|   └── interface.go
├── models            // プリミティブな値を管理する
|   └── ...
├── services          // ビジネスロジックを管理する
|   ├── port/
|   ├── mock/
|   └── ...
├── validator         // 入力データのバリデーションを管理する
|   └── validator.go
├── Dockerfils
├── go.mod
├── go.sum
├── main.go
└── README.md
```

main.go → handlers → services → infrastructures の一方通行の package 依存関係です。

各 package は interface.go のみで通信し、interface.go から mock を生成します。

## 環境変数の一覧

以下の環境変数を設定してください

| 変数名         | 役割                           | デフォルト値 | DEV 環境での値 |
| -------------- | ------------------------------ | ------------ | -------------- |
| PORT           | バックエンドサーバのポート番号 | 8080         |                |
| GCP_PROJECT_ID | GCP のプロジェクト ID          |              |                |
| AUTHORIZATION_HEADER_TAG_NAME | ヘッダのAuthorizationのタグ名 | | Authorization |
| ALLOW_ORIGINS | CORSのAllowOrigins | | http://localhost:5173 |
