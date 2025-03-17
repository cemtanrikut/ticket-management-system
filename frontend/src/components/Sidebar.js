import React from "react";
import { Drawer, List, ListItem, ListItemText } from "@mui/material";
import { Link } from "react-router-dom";

const Sidebar = () => {
    return (
        <Drawer variant="permanent">
            <List>
                <ListItem button component={Link} to="/dashboard">
                    <ListItemText primary="Dashboard" />
                </ListItem>
                <ListItem button component={Link} to="/tickets">
                    <ListItemText primary="Tickets" />
                </ListItem>
                <ListItem button component={Link} to="/customers">
                    <ListItemText primary="Customers" />
                </ListItem>
                <ListItem button component={Link} to="/buildings">
                    <ListItemText primary="Buildings" />
                </ListItem>
                <ListItem button component={Link} to="/workers">
                    <ListItemText primary="Workers" />
                </ListItem>
            </List>
        </Drawer>
    );
};

export default Sidebar;
