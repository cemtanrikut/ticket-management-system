import React, { useEffect, useState } from "react";
import { getWorkers, createWorker, deleteWorker } from "../api/api";
import {
    Container,
    Typography,
    List,
    ListItem,
    ListItemText,
    Paper,
    Divider,
    TextField,
    InputAdornment,
    IconButton,
    Button
} from "@mui/material";
import PeopleIcon from "@mui/icons-material/People";
import SearchIcon from "@mui/icons-material/Search";
import AddCircleIcon from "@mui/icons-material/AddCircle";
import DeleteIcon from "@mui/icons-material/Delete";
import WorkIcon from "@mui/icons-material/Work";
import PhoneIcon from "@mui/icons-material/Phone";
import EmailIcon from "@mui/icons-material/Email";

const Workers = () => {
    const [workers, setWorkers] = useState([]);
    const [searchQuery, setSearchQuery] = useState("");
    const [newWorker, setNewWorker] = useState({
        name: "",
        email: "",
        phone: "",
        department: "",
    });

    useEffect(() => {
        fetchWorkers();
    }, []);

    const fetchWorkers = async () => {
        try {
            const response = await getWorkers();
            setWorkers(response.data);
        } catch (error) {
            console.error("Error fetching workers:", error);
        }
    };

    const handleAddWorker = async () => {
        if (!newWorker.name.trim() || !newWorker.email.trim() || !newWorker.phone.trim()) return;

        try {
            await createWorker(newWorker);
            setNewWorker({ name: "", email: "", phone: "", department: "" });
            fetchWorkers();
        } catch (error) {
            console.error("Error adding worker:", error);
        }
    };

    const handleDeleteWorker = async (workerId) => {
        try {
            await deleteWorker(workerId);
            fetchWorkers();
        } catch (error) {
            console.error("Error deleting worker:", error);
        }
    };

    const filteredWorkers = workers.filter((worker) =>
        worker.name.toLowerCase().includes(searchQuery.toLowerCase())
    );

    return (
        <Container>
            <Typography variant="h4" sx={{ marginBottom: 2 }}>
                👷 Workers
            </Typography>

            {/* Search Bar */}
            <TextField
                label="Search Workers"
                fullWidth
                variant="outlined"
                margin="normal"
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                InputProps={{
                    startAdornment: (
                        <InputAdornment position="start">
                            <SearchIcon />
                        </InputAdornment>
                    ),
                }}
            />

            {/* New Worker Form */}
            <TextField
                label="Worker Name"
                fullWidth
                variant="outlined"
                margin="normal"
                value={newWorker.name}
                onChange={(e) => setNewWorker({ ...newWorker, name: e.target.value })}
            />
            <TextField
                label="Email"
                fullWidth
                variant="outlined"
                margin="normal"
                value={newWorker.email}
                onChange={(e) => setNewWorker({ ...newWorker, email: e.target.value })}
                InputProps={{
                    startAdornment: (
                        <InputAdornment position="start">
                            <EmailIcon />
                        </InputAdornment>
                    ),
                }}
            />
            <TextField
                label="Phone"
                fullWidth
                variant="outlined"
                margin="normal"
                value={newWorker.phone}
                onChange={(e) => setNewWorker({ ...newWorker, phone: e.target.value })}
                InputProps={{
                    startAdornment: (
                        <InputAdornment position="start">
                            <PhoneIcon />
                        </InputAdornment>
                    ),
                }}
            />
            <TextField
                label="Department"
                fullWidth
                variant="outlined"
                margin="normal"
                value={newWorker.department}
                onChange={(e) => setNewWorker({ ...newWorker, department: e.target.value })}
                InputProps={{
                    startAdornment: (
                        <InputAdornment position="start">
                            <WorkIcon />
                        </InputAdornment>
                    ),
                }}
            />
            <Button
                variant="contained"
                color="primary"
                startIcon={<AddCircleIcon />}
                fullWidth
                onClick={handleAddWorker}
            >
                Add Worker
            </Button>

            {/* Workers List */}
            <Paper elevation={3} sx={{ padding: 2, backgroundColor: "#f5f5f5", marginTop: 2 }}>
                <Typography variant="h6">
                    <PeopleIcon sx={{ verticalAlign: "middle", marginRight: 1 }} />
                    Worker List
                </Typography>
                <Divider sx={{ marginY: 1 }} />
                <List>
                    {filteredWorkers.map((worker) => (
                        <ListItem
                            key={worker.id}
                            sx={{
                                borderRadius: 2,
                                backgroundColor: "#fff",
                                marginBottom: 1,
                                boxShadow: 1,
                                display: "flex",
                                justifyContent: "space-between",
                            }}
                        >
                            <ListItemText
                                primary={worker.name}
                                secondary={
                                    <>
                                        <EmailIcon sx={{ fontSize: 14, marginRight: 1 }} />
                                        {worker.email}{" | "}
                                        <PhoneIcon sx={{ fontSize: 14, marginRight: 1 }} />
                                        {worker.phone}{" | "}
                                        <WorkIcon sx={{ fontSize: 14, marginRight: 1 }} />
                                        {worker.department}
                                    </>
                                }
                            />
                            {/* Delete Button */}
                            <IconButton color="error" onClick={() => handleDeleteWorker(worker.id)}>
                                <DeleteIcon />
                            </IconButton>
                        </ListItem>
                    ))}
                </List>
            </Paper>
        </Container>
    );
};

export default Workers;
