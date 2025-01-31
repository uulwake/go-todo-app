import { useState } from "react";
import { Alert as AlertMui, AlertTitle, Snackbar } from "@mui/material";

const Alert = ({
  title,
  message,
  handleClearMessage,
  type = "error",
  vertical = "top",
  horizontal = "center",
  duration = 5e3,
}) => {
  const [isOpen, setOpen] = useState(true);

  const handleClose = () => {
    setOpen(false);
    handleClearMessage && handleClearMessage();
  };

  return (
    <Snackbar
      open={isOpen}
      autoHideDuration={duration}
      onClose={handleClose}
      anchorOrigin={{ vertical, horizontal }}
    >
      <AlertMui severity={type} onClose={handleClose}>
        <AlertTitle>{title}</AlertTitle>
        {message}
      </AlertMui>
    </Snackbar>
  );
};

export default Alert;
