import { useNavigate } from 'react-router'

import { Button } from './components/ui/button'
import { auth } from './lib/firebase'

const App = () => {
  const navigate = useNavigate()

  return (
    <div className="container mx-auto text-center pt-32">
      <h1 className="text-7xl">WebTech 道場</h1>
      <div className="flex flex-col items-center mt-8 gap-2">
        <p>
          WebTech
          道場は、Webエンジニアを目指すあなたの学びを加速させるインタラクティブな学習ツールです。
        </p>
        <p>
          自由記述で答えた回答にAIがリアルタイムでフィードバックを提供し、深掘り質問で理解をさらに深めます。
        </p>
        <p>学習の進捗を確認しながら、着実に技術力を向上させましょう。</p>
        <p>さあ、一緒に次のステップに進みませんか？</p>
      </div>
      <Button className="mt-6" onClick={() => navigate('/signup')}>
        {auth.currentUser ? '問題一覧へ' : '登録して始める'}
      </Button>
    </div>
  )
}

export default App
