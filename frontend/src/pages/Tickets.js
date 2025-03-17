import React, { useEffect, useState } from "react";
import { getTickets, getTicketById, getMessages, sendMessage } from "../api/api";
import {
    Container,
    Typography,
    List,
    ListItem,
    ListItemText,
    Grid,
    Tabs,
    Tab,
    Paper,
    TextField,
    Button,
    Divider,
    IconButton
} from "@mui/material";
import ChatIcon from "@mui/icons-material/Chat";
import SendIcon from "@mui/icons-material/Send";
import AssignmentIcon from "@mui/icons-material/Assignment";
import AccountCircleIcon from "@mui/icons-material/AccountCircle";
import ApartmentIcon from "@mui/icons-material/Apartment";

const Tickets = () => {
    const [tickets, setTickets] = useState([]);
    const [selectedTicket, setSelectedTicket] = useState(null);
    const [messages, setMessages] = useState([]);
    const [newMessage, setNewMessage] = useState("");
    const [statusFilter, setStatusFilter] = useState("to-do");

    useEffect(() => {
        fetchTickets();
    }, [statusFilter]);

    const fetchTickets = async () => {
        try {
            const response = await getTickets();
            setTickets(response.data);
        } catch (error) {
            console.error("Error fetching tickets:", error);
        }
    };

    const handleTicketClick = async (ticketId) => {
        try {
            const response = await getTicketById(ticketId);
            setSelectedTicket(response.data);
            fetchMessages(ticketId);
        } catch (error) {
            console.error("Error fetching ticket details:", error);
        }
    };

    const fetchMessages = async (ticketId) => {
        try {
            const response = await getMessages(ticketId);
            setMessages(response.data);
        } catch (error) {
            console.error("Error fetching messages:", error);
        }
    };

    const handleSendMessage = async () => {
        if (!newMessage.trim() || !selectedTicket) return;

        try {
            await sendMessage(selectedTicket.id, {
                text: newMessage,
                senderId: "user123", // TODO: Kullanıcı kimliği buraya dinamik olarak eklenecek
            });

            setNewMessage("");
            fetchMessages(selectedTicket.id); // Yeni mesajı aldıktan sonra mesajları güncelle
        } catch (error) {
            console.error("Error sending message:", error);
        }
    };

    return (
        <Container>
            <Typography variant="h4" sx={{ marginBottom: 2 }}>🎫 Ticket Management</Typography>

            {/* Status Tabs */}
            <Tabs value={statusFilter} onChange={(e, newValue) => setStatusFilter(newValue)}>
                <Tab label="📝 To Do" value="to-do" />
                <Tab label="🚀 In Progress" value="in-progress" />
                <Tab label="✅ Done" value="done" />
            </Tabs>

            <Grid container spacing={2} sx={{ marginTop: 2 }}>
                {/* Left Panel: Ticket List */}
                <Grid item xs={4}>
                    <Paper elevation={3} sx={{ padding: 2, backgroundColor: "#f5f5f5" }}>
                        <Typography variant="h6"><AssignmentIcon sx={{ verticalAlign: "middle", marginRight: 1 }} /> Tickets</Typography>
                        <Divider sx={{ marginY: 1 }} />
                        <List>
                            {tickets
                                .filter(ticket => ticket.status === statusFilter)
                                .map((ticket) => (
                                    <ListItem
                                        button
                                        key={ticket.id}
                                        onClick={() => handleTicketClick(ticket.id)}
                                        sx={{ borderRadius: 2, backgroundColor: "#fff", marginBottom: 1, boxShadow: 1 }}
                                    >
                                        <ListItemText
                                            primary={ticket.title}
                                            secondary={`Status: ${ticket.status}`}
                                        />
                                    </ListItem>
                                ))}
                        </List>
                    </Paper>
                </Grid>

                {/* Right Panel: Ticket Details & Messaging */}
                <Grid item xs={8}>
                    {selectedTicket ? (
                        <Paper elevation={3} sx={{ padding: 2, backgroundColor: "#e3f2fd" }}>
                            <Typography variant="h6">{selectedTicket.title}</Typography>
                            <Typography variant="body1"><strong>Description:</strong> {selectedTicket.description}</Typography>
                            <Typography variant="body2"><strong>Status:</strong> {selectedTicket.status}</Typography>
                            <Typography variant="body2"><AccountCircleIcon sx={{ fontSize: 18, marginRight: 1 }} /><strong>Assigned Worker:</strong> {selectedTicket.workerId}</Typography>
                            <Typography variant="body2"><ApartmentIcon sx={{ fontSize: 18, marginRight: 1 }} /><strong>Building:</strong> {selectedTicket.buildingId}</Typography>
                            <Typography variant="body2"><strong>Customer:</strong> {selectedTicket.customerId}</Typography>
                            <Typography variant="body2"><strong>Created At:</strong> {selectedTicket.createdAt}</Typography>

                            {/* Messaging Area */}
                            <Typography variant="h6" sx={{ marginTop: 2 }}><ChatIcon sx={{ marginRight: 1 }} /> Messages</Typography>
                            <Divider sx={{ marginY: 1 }} />
                            <List sx={{ maxHeight: 200, overflowY: "auto", backgroundColor: "#fff", borderRadius: 2 }}>
                                {messages.map((msg, index) => (
                                    <ListItem key={index}>
                                        <ListItemText primary={msg.text} secondary={`Sent by: ${msg.senderId}`} />
                                    </ListItem>
                                ))}
                            </List>

                            {/* Message Input */}
                            <Grid container spacing={1} alignItems="center" sx={{ marginTop: 2 }}>
                                <Grid item xs={10}>
                                    <TextField
                                        label="Type a message..."
                                        fullWidth
                                        margin="normal"
                                        value={newMessage}
                                        onChange={(e) => setNewMessage(e.target.value)}
                                    />
                                </Grid>
                                <Grid item xs={2}>
                                    <IconButton color="primary" onClick={handleSendMessage}>
                                        <SendIcon />
                                    </IconButton>
                                </Grid>
                            </Grid>
                        </Paper>
                    ) : (
                        <Typography variant="h6">Select a ticket to view details</Typography>
                    )}
                </Grid>
            </Grid>
        </Container>
    );
};

export default Tickets;
