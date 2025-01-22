# web-tech-dojo

https://web-tech-dojo.benzen-games.com/

<div id="top"></div>

## 使用技術一覧

<p style="display: inline">
  <!-- フロントエンド -->
  <img src="https://img.shields.io/badge/-React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB">
  <img src="https://img.shields.io/badge/-TailwindCSS-000000.svg?logo=tailwindcss&style=for-the-badge">
  <img src="https://img.shields.io/badge/shadcn/ui-000000?style=for-the-badge&logo=shadcn/ui&logoColor=white">
  <img src="https://img.shields.io/badge/TypeScript-3178C6?style=for-the-badge&logo=typescript&logoColor=white">
  <!-- バックエンド -->
  <img src="https://img.shields.io/badge/Go-00ADD8?logo=Go&logoColor=white&style=for-the-badge">
  <!-- インフラ -->
  <img src="https://img.shields.io/badge/-Google%20Cloud%20Platform-4285F4?style=for-the-badge&logo=google%20cloud&logoColor=white">
  <img src="https://img.shields.io/badge/firebase-ffca28?style=for-the-badge&logo=firebase&logoColor=black">
  <!-- AI -->
  <img src="https://img.shields.io/badge/Google%20Gemini-886FBF?style=for-the-badge&logo=googlegemini&logoColor=fff">
</p>

## 目次

1. [プロジェクトについて](#プロジェクトについて)
2. [環境](#環境)
3. [ディレクトリ構成](#ディレクトリ構成)
4. [開発環境構築](#開発環境構築)
5. [環境変数の一覧](#環境変数の一覧)

## プロジェクト名

Web Tech Dojo

## プロジェクトについて

WebTechDojo は、Web 技術に特化した学習をサポートするためのインタラクティブなトレーニングプラットフォームです。このアプリは、ユーザーが技術的な質問に答えたり、理解が浅い箇所を AI が自動的に見つけ出し、深堀りの質問を提示することで、効率的な学習を促します。

## 環境

- フロントエンド: React, Typescript, Tailwind CSS, Shadcn UI
- バックエンド: Go, echo (Cloud Functions)
- AI モデル: Google Cloud's Vertex AI
- インフラ: Google Cloud Functions, Firebase

その他のパッケージのバージョンは package.json を参照してください

## ディレクトリ構成

後で書く

```
.
├── backend
├── frontend
└── README.md
```

## 開発環境構築

後で書く

### 環境変数の一覧

| 変数名                              | 役割                                   | デフォルト値 | DEV 環境での値 |
| ----------------------------------- | -------------------------------------- | ------------ | -------------- |
| SERVICE_ACCOUNT_FOR_CLOUD_FUNCTIONS | Cloud Functions 用のサービスアカウント | secret       |                |
| SERVICE_ACCOUNT_FOR_HOSTING         | Hosting 用のサービスアカウント         | secret       |                |
| WORKLOAD_IDENTITY_PROVIDER          | Workload Identity 用のプロバイダー     | secret       |                |
