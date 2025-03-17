import React, { useEffect, useState } from "react";
import { getCustomers } from "../api/api";
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
    IconButton
} from "@mui/material";
import BusinessIcon from "@mui/icons-material/Business";
import SearchIcon from "@mui/icons-material/Search";

const Customers = () => {
    const [customers, setCustomers] = useState([]);
    const [searchQuery, setSearchQuery] = useState("");

    useEffect(() => {
        const fetchCustomers = async () => {
            try {
                const response = await getCustomers();
                setCustomers(response.data);
            } catch (error) {
                console.error("Error fetching customers:", error);
            }
        };

        fetchCustomers();
    }, []);

    const filteredCustomers = customers.filter((customer) =>
        customer.name.toLowerCase().includes(searchQuery.toLowerCase())
    );

    return (
        <Container>
            <Typography variant="h4" sx={{ marginBottom: 2 }}>
                🏢 Customers
            </Typography>

            {/* Search Bar */}
            <TextField
                label="Search Customers"
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

            {/* Customers List */}
            <Paper elevation={3} sx={{ padding: 2, backgroundColor: "#f5f5f5" }}>
                <Typography variant="h6">
                    <BusinessIcon sx={{ verticalAlign: "middle", marginRight: 1 }} />
                    Customer List
                </Typography>
                <Divider sx={{ marginY: 1 }} />
                <List>
                    {filteredCustomers.map((customer) => (
                        <ListItem
                            key={customer.id}
                            sx={{
                                borderRadius: 2,
                                backgroundColor: "#fff",
                                marginBottom: 1,
                                boxShadow: 1,
                            }}
                        >
                            <ListItemText
                                primary={customer.name}
                                secondary={`Email: ${customer.email} | Phone: ${customer.phone}`}
                            />
                        </ListItem>
                    ))}
                </List>
            </Paper>
        </Container>
    );
};

export default Customers;
