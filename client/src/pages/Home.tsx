import { Link } from "react-router-dom"
import { useAuth } from "../contexts/useAuthContext"
import axios from "axios"

export function Home() {
  const {token, setToken} = useAuth()

  async function handleLogoutClick() {
    await axios.get('/logout')
    setToken('')
  }
  return (
    <div style={{}}>
      {
         token ? 
        <>
         <Link to='/protected'>
          Protected Route
        </Link>
        <button onClick={handleLogoutClick}>Logout</button>
       </> : <>
         <Link to='/login'>Login</Link>
         <Link to='/signup'>Signup</Link>
       </>
      }
    </div>
   
  )
}
