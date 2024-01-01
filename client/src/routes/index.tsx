import { BrowserRouter, Route, Routes } from "react-router-dom";
import { ProtectedRoute } from "./ProtectedRoute";
import { LoginForm } from "../components/Login";
import { SignupForm } from "../components/Singup";
import { Home } from "../components/Home";
import { useAuth } from "../contexts/useAuthContext";

export function Router() {
  const {loading} = useAuth()

  if(loading) return <h1>Loading...</h1>
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<LoginForm />} />
        <Route path="/signup" element={<SignupForm />} />
  
        <Route element={<ProtectedRoute />}>
          <Route path="/" element={<Home />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}
