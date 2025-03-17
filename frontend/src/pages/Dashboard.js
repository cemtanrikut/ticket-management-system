import React, { useEffect, useState } from "react";
import { getTickets } from "../api/api";
import {
    Container,
    Grid,
    Paper,
    Typography,
    CircularProgress
} from "@mui/material";
import AssignmentTurnedInIcon from "@mui/icons-material/AssignmentTurnedIn";
import HourglassEmptyIcon from "@mui/icons-material/HourglassEmpty";
import SyncIcon from "@mui/icons-material/Sync";
import AssessmentIcon from "@mui/icons-material/Assessment";

const Dashboard = () => {
    const [ticketStats, setTicketStats] = useState({
        total: 0,
        todo: 0,
        inProgress: 0,
        done: 0
    });
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchTicketStats = async () => {
            try {
                const response = await getTickets();
                const tickets = response.data;

                const todo = tickets.filter(ticket => ticket.status === "to-do").length;
                const inProgress = tickets.filter(ticket => ticket.status === "in-progress").length;
                const done = tickets.filter(ticket => ticket.status === "done").length;

                setTicketStats({
                    total: tickets.length,
                    todo,
                    inProgress,
                    done
                });
                setLoading(false);
            } catch (error) {
                console.error("Error fetching ticket stats:", error);
                setLoading(false);
            }
        };

        fetchTicketStats();
    }, []);

    return (
        <Container>
            <Typography variant="h4" sx={{ marginBottom: 2 }}>📊 Dashboard</Typography>

            {loading ? (
                <CircularProgress />
            ) : (
                <Grid container spacing={3}>
                    {/* Total Tickets */}
                    <Grid item xs={12} md={4}>
                        <Paper elevation={3} sx={{ padding: 3, textAlign: "center", backgroundColor: "#f5f5f5" }}>
                            <AssessmentIcon sx={{ fontSize: 50, color: "#1976d2" }} />
                            <Typography variant="h6">Total Tickets</Typography>
                            <Typography variant="h4">{ticketStats.total}</Typography>
                        </Paper>
                    </Grid>

                    {/* To Do Tickets */}
                    <Grid item xs={12} md={4}>
                        <Paper elevation={3} sx={{ padding: 3, textAlign: "center", backgroundColor: "#ffebee" }}>
                            <HourglassEmptyIcon sx={{ fontSize: 50, color: "#d32f2f" }} />
                            <Typography variant="h6">To Do</Typography>
                            <Typography variant="h4">{ticketStats.todo}</Typography>
                        </Paper>
                    </Grid>

                    {/* In Progress Tickets */}
                    <Grid item xs={12} md={4}>
                        <Paper elevation={3} sx={{ padding: 3, textAlign: "center", backgroundColor: "#fff3e0" }}>
                            <SyncIcon sx={{ fontSize: 50, color: "#f57c00" }} />
                            <Typography variant="h6">In Progress</Typography>
                            <Typography variant="h4">{ticketStats.inProgress}</Typography>
                        </Paper>
                    </Grid>

                    {/* Done Tickets */}
                    <Grid item xs={12} md={4}>
                        <Paper elevation={3} sx={{ padding: 3, textAlign: "center", backgroundColor: "#e8f5e9" }}>
                            <AssignmentTurnedInIcon sx={{ fontSize: 50, color: "#388e3c" }} />
                            <Typography variant="h6">Done</Typography>
                            <Typography variant="h4">{ticketStats.done}</Typography>
                        </Paper>
                    </Grid>
                </Grid>
            )}
        </Container>
    );
};

export default Dashboard;
