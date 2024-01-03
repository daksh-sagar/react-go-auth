import axios from 'axios'
import { createContext, useEffect, useState, useMemo } from 'react'

axios.defaults.baseURL = 'http://localhost:4000/v1'
axios.defaults.withCredentials = true

export const AuthContext = createContext<{
  token: string
  loading: boolean
  setToken: (token: string) => void
}>({
  token: '',
  loading: false,
  setToken: () => {}
})

export function AuthProvider({ children }: {children: React.ReactNode}){
  const [token, setToken] = useState('')
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    // make a call to the refreshToken endpoint with credentials and get the jwt in return (if refresh token is present in cookies)
    (async function refreshToken() {
      try {
        setLoading(true)
        const res = await axios.get('/refreshToken')
  
        const {accessToken} = res.data.tokens
        setToken(accessToken)
      } catch (error) {
        console.error({error})
        console.log('refresh token not found in cookies')
        setToken('')
      } finally {
        setLoading(false)
      }
    })()
  }, [])

  useEffect(() => {
    if(token) {
      axios.defaults.headers.common["Authorization"] = "Bearer " + token
    } else {
      delete axios.defaults.headers.common["Authorization"]
    }
  }, [token])

  const contextValue = useMemo(
    () => ({
      token,
      loading,
      setToken,
    }),
    [token,loading]
  )

  return <AuthContext.Provider value={contextValue}>{children}</AuthContext.Provider>
}