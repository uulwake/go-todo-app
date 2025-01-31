import { useContext, useEffect, useState } from "react";
import Box from "@mui/material/Box";
import LoadingButton from "@mui/lab/LoadingButton";
import Divider from "@mui/material/Divider";
import FormLabel from "@mui/material/FormLabel";
import FormControl from "@mui/material/FormControl";
import Link from "@mui/material/Link";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import Container from "../components/Auth/Container";
import Card from "../components/Auth/Card";
import axios from "../axios";
import AuthContext from "../context/AuthContext";
import { useNavigate } from "react-router-dom";
import Alert from "../components/Alert";

const Login = () => {
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);
  const { setAccessToken, setUser, accessToken } = useContext(AuthContext);
  const navigate = useNavigate();

  useEffect(() => {
    if (accessToken) {
      navigate("/", { replace: true });
    }

    return () => {
      setIsLoading(false);
      setError(null);
    };
  }, [accessToken, navigate]);

  const handleSubmit = async (event) => {
    event.preventDefault();
    setError(null);
    setIsLoading(true);

    const data = new FormData(event.currentTarget);
    try {
      const { data: userData } = await axios({
        method: "POST",
        url: "/v1/users/login",
        data: {
          email: data.get("email"),
          password: data.get("password"),
        },
      });

      const token = userData.data.jwt_token;
      const userId = userData.data.user_id;
      const userName = userData.data.user_name;
      setAccessToken(token);
      setUser({ id: userId, name: userName });
      localStorage.setItem("accessToken", token);
      localStorage.setItem(
        "user",
        JSON.stringify({ id: userId, name: userName })
      );

      setIsLoading(false);
      navigate("/", { replace: true });
    } catch (err) {
      setError(err.response.data.message);
      setIsLoading(false);
    }
  };

  return (
    <>
      <Container direction="column" justifyContent="space-between">
        {error && <Alert title="Gagal Masuk" message={error} />}
        <Card variant="outlined">
          <Typography
            component="h1"
            variant="h4"
            sx={{ width: "100%", fontSize: "clamp(2rem, 10vw, 2.15rem)" }}
          >
            Login
          </Typography>
          <Box
            component="form"
            onSubmit={handleSubmit}
            noValidate
            sx={{
              display: "flex",
              flexDirection: "column",
              width: "100%",
              gap: 2,
            }}
          >
            <FormControl>
              <FormLabel htmlFor="email">Email</FormLabel>
              <TextField
                id="email"
                type="email"
                name="email"
                placeholder="your@email.com"
                autoComplete="email"
                autoFocus
                required
                fullWidth
                variant="outlined"
              />
            </FormControl>
            <FormControl>
              <FormLabel htmlFor="password">Password</FormLabel>
              <TextField
                name="password"
                placeholder="••••••"
                type="password"
                id="password"
                autoComplete="current-password"
                autoFocus
                required
                fullWidth
                variant="outlined"
              />
            </FormControl>
            <LoadingButton
              type="submit"
              fullWidth
              variant="contained"
              loading={isLoading}
            >
              Login
            </LoadingButton>
          </Box>
          <Divider />
          <Box sx={{ display: "flex", flexDirection: "column", gap: 2 }}>
            <Typography sx={{ textAlign: "center" }}>
              Don&apos;t have an account?{" "}
              <Link
                href="/register"
                variant="body2"
                sx={{ alignSelf: "center" }}
              >
                Register
              </Link>
            </Typography>
          </Box>
        </Card>
      </Container>
    </>
  );
};

export default Login;
