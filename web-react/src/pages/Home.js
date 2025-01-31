import { useContext, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import AuthContext from "../context/AuthContext";
import NavBar from "../components/NavBar";
import { useQuery, useMutation, useQueryClient } from "react-query";
import axios from "../axios";
import LoadingSpinner from "../components/LoadingSpinner";
import LoadingButton from "@mui/lab/LoadingButton";
import {
  Box,
  IconButton,
  Table,
  TableBody,
  TableCell,
  TableRow,
  TextField,
  Typography,
} from "@mui/material";
import CheckIcon from "@mui/icons-material/Check";
import Pagination from "@mui/material/Pagination";

const Home = () => {
  const [taskTitle, setTaskTitle] = useState("");
  const [page, setPage] = useState(1);
  const [size] = useState(5);
  const { accessToken, user, setAccessToken, setUser } =
    useContext(AuthContext);
  const navigate = useNavigate();
  const queryClient = useQueryClient();

  const queryGetTask = "queryGetTask";
  const { data: tasks, isLoading: isLoadingTask } = useQuery(
    [queryGetTask, page, size],
    async () => {
      const { data } = await axios({
        method: "GET",
        url: `/v1/tasks`,
        params: {
          page_size: size,
          page_number: page,
        },
        headers: {
          Authorization: accessToken || localStorage.getItem("accessToken"),
        },
      });

      return data;
    },
    { retry: 0 }
  );

  const mutationCreate = useMutation(
    async (newTask) => {
      await axios({
        method: "POST",
        url: `/v1/tasks`,
        data: newTask,
        headers: {
          Authorization: accessToken || localStorage.getItem("accessToken"),
        },
      });
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries(queryGetTask);
      },
    }
  );

  const mutationComplete = useMutation(
    async (taskId) => {
      await axios({
        method: "PATCH",
        url: `/v1/tasks/${taskId}`,
        headers: {
          Authorization: accessToken || localStorage.getItem("accessToken"),
        },
      });
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries(queryGetTask);
      },
    }
  );

  const handleAddTask = async (e) => {
    e.preventDefault();
    await mutationCreate.mutateAsync({ title: taskTitle });
    setTaskTitle("");
  };

  const handleCompleteTask = async (taskId) => {
    await mutationComplete.mutateAsync(taskId);
  };

  const formatDate = (dateInput, dateStyle = "full") => {
    if (!dateInput) {
      dateInput = new Date();
    } else {
      dateInput = new Date(dateInput);
    }

    const formatter = new Intl.DateTimeFormat("id-ID", { dateStyle });
    return formatter.format(dateInput);
  };

  useEffect(() => {
    if (!accessToken) {
      navigate("/login");
    }
  }, [accessToken, navigate]);

  useEffect(() => {
    if (tasks && !tasks.data.length && tasks.page.total > 0) {
      setPage((page) => page - 1);
    }
  }, [tasks]);

  if (!user || !accessToken || isLoadingTask) {
    return <LoadingSpinner />;
  }

  return (
    <>
      <NavBar user={user} setAccessToken={setAccessToken} setUser={setUser} />
      <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          flexDirection: "column",
        }}
      >
        <Box
          sx={{
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
            flexDirection: "row",
          }}
          component="form"
          onSubmit={handleAddTask}
        >
          <TextField
            id="title"
            name="title"
            type="text"
            label="New Task"
            variant="standard"
            sx={{ marginTop: 2 }}
            autoComplete="off"
            onChange={(e) => setTaskTitle(e.target.value)}
            value={taskTitle}
          />
          <LoadingButton
            variant="contained"
            loading={mutationCreate.isLoading}
            type="submit"
            sx={{ marginTop: 2, marginLeft: 3 }}
          >
            Add
          </LoadingButton>
        </Box>
        <Box
          sx={{
            marginTop: 3,
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
            flexDirection: "column",
          }}
        >
          {tasks.data && !!tasks.data.length ? (
            <>
              <Table sx={{ width: "500px" }} size="small">
                <TableBody>
                  {tasks.data.map((task, idx) => (
                    <TableRow key={task.task_id} sx={{ height: 3 }}>
                      <TableCell>{(page - 1) * size + idx + 1}</TableCell>
                      <TableCell>{task.task_title}</TableCell>
                      <TableCell>{formatDate(task.task_created_at)}</TableCell>
                      <TableCell>
                        <IconButton
                          onClick={() => handleCompleteTask(task.task_id)}
                        >
                          <CheckIcon />
                        </IconButton>
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
              <Pagination
                sx={{ marginTop: 2 }}
                count={Math.ceil(tasks.page.total / tasks.page.size)}
                page={page}
                onChange={(e, value) => setPage(value)}
              />
            </>
          ) : (
            <Typography marginTop={5} variant="h3">
              No Task Today!
            </Typography>
          )}
        </Box>
      </Box>
    </>
  );
};

export default Home;
