import { Box, Button, Typography } from "@mui/material";
import { useNavigate } from "react-router-dom";

const NotFound = () => {
  const navigate = useNavigate();

  return (
    <>
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          justifyContent: "center",
          alignItems: "center",
          width: "100%",
          height: "100vh",
        }}
      >
        <Typography variant="h2">Halaman Tidak Ditemukan</Typography>
        <Button
          variant="contained"
          sx={{ marginTop: 3 }}
          onClick={() => navigate("/")}
        >
          Kembali
        </Button>
      </Box>
    </>
  );
};

export default NotFound;
