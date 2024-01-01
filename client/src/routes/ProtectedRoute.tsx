import { Navigate, Outlet } from "react-router-dom"
import { useAuth } from "../contexts/useAuthContext"


export function ProtectedRoute(){
  const { token, loading } = useAuth()

  console.log({token, loading})

  if (!token) {
    return <Navigate to="/login" />
  }

  return <Outlet />
}