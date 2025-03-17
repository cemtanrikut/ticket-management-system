import React, { useEffect, useState } from "react";
import { getBuildings, createBuilding, updateBuilding, deleteBuilding } from "../api/api";
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
import ApartmentIcon from "@mui/icons-material/Apartment";
import SearchIcon from "@mui/icons-material/Search";
import BusinessIcon from "@mui/icons-material/Business";
import AddCircleIcon from "@mui/icons-material/AddCircle";
import EditIcon from "@mui/icons-material/Edit";
import DeleteIcon from "@mui/icons-material/Delete";

const Buildings = () => {
    const [buildings, setBuildings] = useState([]);
    const [searchQuery, setSearchQuery] = useState("");
    const [newBuilding, setNewBuilding] = useState("");

    useEffect(() => {
        fetchBuildings();
    }, []);

    const fetchBuildings = async () => {
        try {
            const response = await getBuildings();
            setBuildings(response.data);
        } catch (error) {
            console.error("Error fetching buildings:", error);
        }
    };

    const handleAddBuilding = async () => {
        if (!newBuilding.trim()) return;

        try {
            await createBuilding({ name: newBuilding });
            setNewBuilding("");
            fetchBuildings();
        } catch (error) {
            console.error("Error adding building:", error);
        }
    };

    const handleDeleteBuilding = async (buildingId) => {
        try {
            await deleteBuilding(buildingId);
            fetchBuildings();
        } catch (error) {
            console.error("Error deleting building:", error);
        }
    };

    const filteredBuildings = buildings.filter((building) =>
        building.name.toLowerCase().includes(searchQuery.toLowerCase())
    );

    return (
        <Container>
            <Typography variant="h4" sx={{ marginBottom: 2 }}>
                🏢 Buildings
            </Typography>

            {/* Search Bar */}
            <TextField
                label="Search Buildings"
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

            {/* New Building Input */}
            <TextField
                label="New Building Name"
                fullWidth
                variant="outlined"
                margin="normal"
                value={newBuilding}
                onChange={(e) => setNewBuilding(e.target.value)}
            />
            <Button
                variant="contained"
                color="primary"
                startIcon={<AddCircleIcon />}
                fullWidth
                onClick={handleAddBuilding}
            >
                Add Building
            </Button>

            {/* Buildings List */}
            <Paper elevation={3} sx={{ padding: 2, backgroundColor: "#f5f5f5", marginTop: 2 }}>
                <Typography variant="h6">
                    <ApartmentIcon sx={{ verticalAlign: "middle", marginRight: 1 }} />
                    Building List
                </Typography>
                <Divider sx={{ marginY: 1 }} />
                <List>
                    {filteredBuildings.map((building) => (
                        <ListItem
                            key={building.id}
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
                                primary={building.name}
                                secondary={
                                    <>
                                        <BusinessIcon sx={{ fontSize: 14, marginRight: 1 }} />
                                        {`Customer ID: ${building.customerId}`}
                                    </>
                                }
                            />
                            {/* Action Buttons */}
                            <IconButton color="primary">
                                <EditIcon />
                            </IconButton>
                            <IconButton color="error" onClick={() => handleDeleteBuilding(building.id)}>
                                <DeleteIcon />
                            </IconButton>
                        </ListItem>
                    ))}
                </List>
            </Paper>
        </Container>
    );
};

export default Buildings;
