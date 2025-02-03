import { useEffect, useRef, useState } from 'react'

import axios, { AxiosResponse } from 'axios'
import { FaArrowUp } from 'react-icons/fa'
import { useNavigate, useParams } from 'react-router'

import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'
import { Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'
import { Textarea } from '@/components/ui/textarea'
import { toast } from '@/hooks/use-toast'
import { Message } from '@/types/message'
import { Question } from '@/types/question'

const ChatPage = () => {
  const inputRef = useRef<HTMLDivElement>(null)
  const endOfMessagesRef = useRef<HTMLDivElement>(null)

  const [loading, setLoading] = useState<boolean>(false)
  const [waiting, setWaiting] = useState<boolean>(false)
  const [content, setContent] = useState<string>('')
  const [messages, setMessages] = useState<Message[]>([])
  const [input, setInput] = useState<string>('')

  const navigate = useNavigate()
  const { id } = useParams()

  useEffect(() => {
    scrollToBottom()
  }, [messages])

  useEffect(() => {
    let ignore = false
    ;(async () => {
      if (!id) {
        return
      }
      setLoading(true)
      // 2回実行されるのを防ぐ
      // see https://react.dev/learn/synchronizing-with-effects#fetching-data

      try {
        // 問題のタイトルを取得
        const {
          data: questionResponse,
          status: questionResponseStatus,
        }: AxiosResponse<Question> = await axios.get(`/api/question/${id}`)
        if (!ignore && questionResponseStatus === 200) {
          setContent(questionResponse.content)
        }

        // 会話履歴を取得
        const {
          data: messagesResponse,
          status: messagesResponseStatus,
        }: AxiosResponse<{ messages: Message[] }> = await axios.get(
          `/api/question/${id}/answer`,
        )
        if (!ignore && messagesResponseStatus === 200) {
          setMessages(messagesResponse.messages)
        }
      } catch {
        toast({
          variant: 'destructive',
          title: 'Uh oh! Something went wrong.',
          description: 'There was a problem with your request.',
        })
      } finally {
        setLoading(false)
      }
    })()
    return () => {
      ignore = true
    }
  }, [id])

  if (!id) {
    navigate('/questions')
    return null
  }

  const handleSend = async () => {
    if (!input || !input.trim()) {
      return
    }
    setWaiting(true)
    setMessages((prevMessages) => [
      ...prevMessages,
      { text: input, sentByUser: true, createdAt: new Date().toISOString() },
    ])
    setInput('')
    try {
      const { data, status }: AxiosResponse<{ message: string }> =
        await axios.post(`/api/question/${id}/answer`, {
          message: input,
        })
      if (status === 200) {
        setMessages((prevMessages) => [
          ...prevMessages,
          {
            text: data.message,
            sentByUser: false,
            createdAt: new Date().toISOString(),
          },
        ])
      }
    } catch {
      toast({
        variant: 'destructive',
        title: 'Uh oh! Something went wrong.',
        description: 'There was a problem with your request.',
      })
    } finally {
      setWaiting(false)
    }
    scrollToBottom()
  }

  const scrollToBottom = () => {
    if (endOfMessagesRef.current) {
      endOfMessagesRef.current.scrollIntoView({ behavior: 'smooth' })
    }
  }

  if (loading) {
    return (
      <div className='grid grid-rows-[auto_1fr_auto] overflow-hidden h-full'>
        <Skeleton className="mt-4 h-12 w-full rounded-lg" />
        <div className="flex-1 overflow-y-auto p-4">
          <div className="mb-4 w-full flex justify-end">
            <Skeleton className="h-10 w-full rounded-xl max-w-[80%]" />
          </div>
          <div className="w-full text-left flex">
            <Skeleton className="h-10 w-10 rounded-full" />
            <Skeleton className="h-10 inline-block ml-4 rounded-xl w-[80%]" />
          </div>
        </div>
        <div className="mx-auto w-full bg-white pb-2 max-w-[800px]">
          <div className="flex p-4 shadow-md rounded-lg bg-gray-100 items-end cursor-text">
            <Textarea
              className="resize-none max-h-32 border-none shadow-none focus-visible:ring-0"
              disabled
              placeholder="メッセージを入力"
            />
            <Button className="ml-2 rounded-full h-8 w-8" disabled>
              <FaArrowUp />
            </Button>
          </div>
          <p className="text-sm text-gray-500 text-center mt-2">
            AIの回答は必ずしも正しいとは限りません。重要な情報は確認するようにしてください。
          </p>
        </div>
      </div>
    )
  }

  return (
    <div className='grid grid-rows-[auto_1fr_auto] overflow-hidden h-full'>
      <div className="w-full mt-4 p-2 mx-auto bg-gray-100 flex justify-start rounded-lg items-center max-w-[800px]">
        <div className="p-2 bg-white rounded-lg mr-2">Q{id}</div>
        <h1>{content}</h1>
      </div>
      <div className="hidden-scrollbar flex-1 overflow-y-auto p-4">
        {messages.map((message, index) => {
          if (!message.sentByUser) {
            return (
              <div className="mx-auto mb-2 w-full text-left flex max-w-[800px]" key={index}>
                <Avatar>
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
              <div className="mx-auto mb-2 w-full text-right max-w-[800px]" key={index}>
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
        className="mx-auto w-full bg-white pb-2 max-w-[800px]"
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
          {waiting ? (
            <Button className="ml-2 rounded-full h-8 w-8" disabled>
              <FaArrowUp />
            </Button>
          ) : (
            <Button className="ml-2 rounded-full h-8 w-8" onClick={handleSend}>
              <FaArrowUp />
            </Button>
          )}
        </div>
        <p className="text-sm text-gray-500 text-center mt-2">
          AIの回答は必ずしも正しいとは限りません。重要な情報は確認するようにしてください。
        </p>
      </div>
    </div>
  )
}

export default ChatPage
