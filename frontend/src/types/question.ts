export type Question = {
  content: string
  id: string
  isBookmarked: boolean
  progress: number
  tags: string[]
  title: string
}

export const tags = [
  {
    value: 'Web開発',
    label: 'Web開発',
  },
  {
    value: 'HTTP',
    label: 'HTTP',
  },
  {
    value: 'リクエストメソッド',
    label: 'リクエストメソッド',
  },
  {
    value: 'API',
    label: 'API',
  },
  {
    value: 'REST',
    label: 'REST',
  },
  {
    value: 'GraphQL',
    label: 'GraphQL',
  },
  {
    value: 'HTTPS',
    label: 'HTTPS',
  },
  {
    value: 'セキュリティ',
    label: 'セキュリティ',
  },
  {
    value: 'CORS',
    label: 'CORS',
  },
  {
    value: 'フロントエンド',
    label: 'フロントエンド',
  },
  {
    value: 'レンダリング',
    label: 'レンダリング',
  },
  {
    value: 'React',
    label: 'React',
  },
  {
    value: 'Vue.js',
    label: 'Vue.js',
  },
  {
    value: 'Virtual DOM',
    label: 'Virtual DOM',
  },
  {
    value: 'Web Components',
    label: 'Web Components',
  },
  {
    value: 'コンポーネント',
    label: 'コンポーネント',
  },
  {
    value: 'ブラウザ',
    label: 'ブラウザ',
  },
  {
    value: 'JavaScript',
    label: 'JavaScript',
  },
  {
    value: 'イベントループ',
    label: 'イベントループ',
  },
  {
    value: '非同期処理',
    label: '非同期処理',
  },
  {
    value: 'CSS',
    label: 'CSS',
  },
  {
    value: 'レイアウト',
    label: 'レイアウト',
  },
  {
    value: 'バックエンド',
    label: 'バックエンド',
  },
  {
    value: 'Node.js',
    label: 'Node.js',
  },
  {
    value: 'WebSocket',
    label: 'WebSocket',
  },
  {
    value: 'リアルタイム',
    label: 'リアルタイム',
  },
  {
    value: '認証',
    label: '認証',
  },
  {
    value: 'JWT',
    label: 'JWT',
  },
  {
    value: 'データベース',
    label: 'データベース',
  },
  {
    value: 'パフォーマンス',
    label: 'パフォーマンス',
  },
  {
    value: 'ミドルウェア',
    label: 'ミドルウェア',
  },
  {
    value: 'RDBMS',
    label: 'RDBMS',
  },
  {
    value: 'NoSQL',
    label: 'NoSQL',
  },
  {
    value: 'インデックス',
    label: 'インデックス',
  },
  {
    value: 'トランザクション',
    label: 'トランザクション',
  },
  {
    value: 'ACID',
    label: 'ACID',
  },
  {
    value: '正規化',
    label: '正規化',
  },
  {
    value: '非正規化',
    label: '非正規化',
  },
  {
    value: 'MongoDB',
    label: 'MongoDB',
  },
  {
    value: 'CSRF',
    label: 'CSRF',
  },
  {
    value: 'XSS',
    label: 'XSS',
  },
  {
    value: 'SQLインジェクション',
    label: 'SQLインジェクション',
  },
  {
    value: 'SSL',
    label: 'SSL',
  },
  {
    value: 'Cookie',
    label: 'Cookie',
  },
  {
    value: '属性',
    label: '属性',
  },
  {
    value: 'デプロイ',
    label: 'デプロイ',
  },
  {
    value: 'CI/CD',
    label: 'CI/CD',
  },
  {
    value: 'インフラ',
    label: 'インフラ',
  },
  {
    value: 'Docker',
    label: 'Docker',
  },
  {
    value: 'コンテナ',
    label: 'コンテナ',
  },
  {
    value: 'Webサーバー',
    label: 'Webサーバー',
  },
  {
    value: 'Nginx',
    label: 'Nginx',
  },
  {
    value: '負荷分散',
    label: '負荷分散',
  },
  {
    value: 'スケーリング',
    label: 'スケーリング',
  },
  {
    value: 'CDN',
    label: 'CDN',
  },
  {
    value: 'IaC',
    label: 'IaC',
  },
  {
    value: '自動化',
    label: '自動化',
  },
  {
    value: '最適化',
    label: '最適化',
  },
  {
    value: 'TTFB',
    label: 'TTFB',
  },
  {
    value: 'キャッシュ',
    label: 'キャッシュ',
  },
  {
    value: '画像',
    label: '画像',
  },
]

export const bookmarkStatus = [
  {
    value: 'isBookmarked',
    label: 'Is Bookmarked',
  },
  {
    value: 'isNotBookmarked',
    label: 'Is Not Bookmarked',
  },
]

export const progressStatus = [
  {
    value: 'todo',
    label: 'To Do',
  },
  {
    value: 'inProgress',
    label: 'In Progress',
  },
  {
    value: 'completed',
    label: 'Completed',
  },
]
