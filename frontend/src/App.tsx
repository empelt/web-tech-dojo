import { useSignInWithGoogle } from 'react-firebase-hooks/auth'

import auth from './libs/firebase'

const App = () => {
  const [signInWithGoogle, user, loading, error] = useSignInWithGoogle(auth)

  const handleLogout = () => {
    auth.signOut()
    window.location.reload()
  }

  const request = async () => {
    const token = await user?.user.getIdToken()
    fetch(import.meta.env.VITE_BACKEND_URL, {
      method: 'GET',
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((res) => {
        console.log(res)
      })
      .catch((error) => {
        console.error('リクエストエラー:えええお', error)
      })
  }

  if (loading) {
    return <div>Loading...</div>
  }

  if (error) {
    return <div>Error: {error.message}</div>
  }

  return user ? (
    <div>
      <p>{user.user.email} でログイン中</p>
      <button onClick={handleLogout}>ログアウト</button>
      <button onClick={request}>リクエスト！！！</button>
    </div>
  ) : (
    <button onClick={() => signInWithGoogle()}>ログイン</button>
  )
}

export default App
