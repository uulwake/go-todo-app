import "./App.css";
import { createTheme, ThemeProvider, CssBaseline } from "@mui/material";
import AuthContext from "./context/AuthContext";
import { Suspense, useEffect, useState } from "react";
import { Routes, Route } from "react-router-dom";
import { QueryClient, QueryClientProvider } from "react-query";

import Home from "./pages/Home";
import Login from "./pages/Login";
import Register from "./pages/Register";
import NotFound from "./pages/NotFound";

const queryClient = new QueryClient();

function App() {
  const theme = createTheme();
  const [accessToken, setAccessToken] = useState("");
  const [user, setUser] = useState(null);

  useEffect(() => {
    const accessToken = localStorage.getItem("accessToken");
    const user = localStorage.getItem("user");

    if (accessToken && user) {
      setAccessToken(accessToken);
      setUser(JSON.parse(user));
      return;
    }
  }, []);

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <QueryClientProvider client={queryClient}>
        <AuthContext.Provider
          value={{
            accessToken,
            setAccessToken,
            user,
            setUser,
          }}
        >
          <Routes>
            <Route
              path="/"
              element={
                <Suspense fallback={<></>}>
                  <Home />
                </Suspense>
              }
            />
            <Route
              path="/login"
              element={
                <Suspense fallback={<></>}>
                  <Login />
                </Suspense>
              }
            />
            <Route
              path="/register"
              element={
                <Suspense fallback={<></>}>
                  <Register />
                </Suspense>
              }
            />
            <Route
              path="/*"
              element={
                <Suspense fallback={<></>}>
                  <NotFound />
                </Suspense>
              }
            />
          </Routes>
        </AuthContext.Provider>
      </QueryClientProvider>
    </ThemeProvider>
  );
}

export default App;
