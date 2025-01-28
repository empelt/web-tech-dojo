import { useRef, useState } from 'react'

import { useNavigate, useParams } from 'react-router'

import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'
import { Button } from '@/components/ui/button'
import { Textarea } from '@/components/ui/textarea'
import { Message } from '@/types/message'

const sampleMessages: Message[] = [
  {
    text: 'こんにちは！どのようなお悩みがありますか？',
    sentByAI: true,
    createdAt: new Date().toISOString(),
  },
  {
    text: 'ネットワークの設定について教えてくださいネットワークの設定について教えてくださいネットワークの設定について教えてください',
    sentByAI: false,
    createdAt: new Date().toISOString(),
  },
  {
    text: 'ネットワークの設定についてですね。どのようなことについてお知りになりたいですか？ネットワークの設定についてですね。どのようなことについてお知りになりたいですか？',
    sentByAI: true,
    createdAt: new Date().toISOString(),
  },
  {
    text: 'HTTPとHTTPSの違いはなんですか？',
    sentByAI: false,
    createdAt: new Date().toISOString(),
  },
  {
    text: 'HTTPとHTTPSの違いは暗号化の有無です。HTTPは暗号化されていないのに対して、HTTPSは暗号化されています。',
    sentByAI: true,
    createdAt: new Date().toISOString(),
  },
  {
    text: 'ありがとうございます！',
    sentByAI: false,
    createdAt: new Date().toISOString(),
  },
  {
    text: 'どういたしまして！',
    sentByAI: true,
    createdAt: new Date().toISOString(),
  },
  {
    text: '他に質問はありますか？',
    sentByAI: true,
    createdAt: new Date().toISOString(),
  },
  {
    text: 'いいえ、大丈夫です。',
    sentByAI: false,
    createdAt: new Date().toISOString(),
  },
  {
    text: 'それでは、またのご利用をお待ちしております。',
    sentByAI: true,
    createdAt: new Date().toISOString(),
  },
  {
    text: 'こんにちは！どのようなお悩みがありますか？',
    sentByAI: true,
    createdAt: new Date().toISOString(),
  },
  {
    text: 'ネットワークの設定について教えてくださいネットワークの設定について教えてくださいネットワークの設定について教えてください',
    sentByAI: false,
    createdAt: new Date().toISOString(),
  },
  {
    text: 'ネットワークの設定についてですね。どのようなことについてお知りになりたいですか？ネットワークの設定についてですね。どのようなことについてお知りになりたいですか？',
    sentByAI: true,
    createdAt: new Date().toISOString(),
  },
  {
    text: 'HTTPとHTTPSの違いはなんですか？',
    sentByAI: false,
    createdAt: new Date().toISOString(),
  },
  {
    text: 'HTTPとHTTPSの違いは暗号化の有無です。HTTPは暗号化されていないのに対して、HTTPSは暗号化されています。',
    sentByAI: true,
    createdAt: new Date().toISOString(),
  },
  {
    text: 'ありがとうございます！',
    sentByAI: false,
    createdAt: new Date().toISOString(),
  },
  {
    text: 'どういたしまして！',
    sentByAI: true,
    createdAt: new Date().toISOString(),
  },
  {
    text: '他に質問はありますか？',
    sentByAI: true,
    createdAt: new Date().toISOString(),
  },
  {
    text: 'いいえ、大丈夫です。',
    sentByAI: false,
    createdAt: new Date().toISOString(),
  },
  {
    text: 'それでは、またのご利用をお待ちしております。',
    sentByAI: true,
    createdAt: new Date().toISOString(),
  },
]

const ChatPage = () => {
  const inputRef = useRef<HTMLDivElement>(null)
  const endOfMessagesRef = useRef<HTMLDivElement>(null)

  const [messages, setMessages] = useState<string[]>([])
  const [input, setInput] = useState<string>('')

  const navigate = useNavigate()
  const { id } = useParams()
  if (!id) {
    navigate('/questions')
    return null
  }

  const handleSend = () => {
    if (input.trim()) {
      setMessages([...messages, input])
      setInput('')
    }
  }

  // TODO: バックエンドとの繋ぎこみ時に使う予定
  //   const scrollToBottom = () => {
  //     if (ref.current) {
  //       ref.current.scrollIntoView({ behavior: 'smooth' })
  //     }
  //   }

  return (
    <div className="container mx-auto pt-4">
      <div className="w-full p-2 bg-gray-100 flex justify-start rounded-lg items-center">
        <div className="p-2 bg-white rounded-lg mr-2">Q23</div>
        <h1>HTTPとHTTPSの違いはなんですか？</h1>
      </div>
      <div
        className="flex-1 overflow-y-auto p-4"
        style={{ marginBottom: `${inputRef.current?.offsetHeight ?? 120}px` }}>
        {sampleMessages.map((message, index) => {
          if (message.sentByAI) {
            return (
              <div className="mb-2 w-full text-left flex" key={index}>
                <Avatar className="cursor-pointer">
                  <AvatarImage
                    alt="avatar image"
                    src="https://github.com/shadcn.png"
                  />
                  <AvatarFallback>w</AvatarFallback>
                  <span className="sr-only">avatar icon</span>
                </Avatar>
                <div className="inline-block p-2 rounded">{message.text}</div>
              </div>
            )
          } else {
            return (
              <div className="mb-2 w-full text-right" key={index}>
                <div className="inline-block px-3 py-2 bg-gray-200 rounded-xl text-left max-w-[80%]">
                  {message.text}
                </div>
              </div>
            )
          }
        })}
        <div ref={endOfMessagesRef} />
      </div>
      <div
        className="container mx-auto fixed bottom-2 left-1/2 transform -translate-x-1/2"
        ref={inputRef}>
        <div
          className="flex p-4 shadow-md rounded-lg bg-gray-100 items-end cursor-text"
          onClick={() => document.getElementById('chat-input')?.focus()}>
          <Textarea
            className="resize-none max-h-32 border-none shadow-none focus-visible:ring-0"
            id="chat-input"
            onChange={(e) => setInput(e.target.value)}
            onInput={(e) => {
              const target = e.target as HTMLTextAreaElement
              target.style.height = 'auto'
              target.style.height = `${target.scrollHeight}px`
            }}
            placeholder="メッセージを入力"
            value={input}
          />
          <Button className="ml-2" onClick={handleSend}>
            送信
          </Button>
        </div>
        <p className="text-sm text-gray-500 text-center mt-2">
          AIの回答は必ずしも正しいとは限りません。重要な情報は確認するようにしてください。
        </p>
      </div>
    </div>
  )
}

export default ChatPage
