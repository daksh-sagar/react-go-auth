import { BrowserRouter, Route, Routes } from "react-router-dom";
import { ProtectedRoute } from "./ProtectedRoute";
import { LoginForm } from "../pages/Login";
import { SignupForm } from "../pages/Singup";
import { Home } from "../pages/Home";
import { useAuth } from "../contexts/useAuthContext";
import { Protected } from "../pages/Protected";

export function Router() {
  const {loading} = useAuth()

  if(loading) return <h1>Loading...</h1>
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<LoginForm />} />
        <Route path="/signup" element={<SignupForm />} />
        <Route path="/" element={<Home />} />
  
        <Route element={<ProtectedRoute />}>
          <Route path="/protected" element={<Protected />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}
