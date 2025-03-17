import React from "react";
import { Routes, Route, Navigate } from "react-router-dom";
import Sidebar from "./components/Sidebar";
import Dashboard from "./pages/Dashboard";
import Tickets from "./pages/Tickets";
import Customers from "./pages/Customers";
import Buildings from "./pages/Buildings";
import Workers from "./pages/Workers";
import Login from "./pages/Login";
import { Container, Grid } from "@mui/material";

const App = () => {
    const isAuthenticated = localStorage.getItem("token"); // Kullanıcı giriş yapmış mı?

    return (
        <Grid container>
            {/* Eğer giriş yapılmışsa Sidebar göster */}
            {isAuthenticated && (
                <Grid item xs={2}>
                    <Sidebar />
                </Grid>
            )}

            <Grid item xs={isAuthenticated ? 10 : 12}>
                <Container sx={{ padding: 3 }}>
                    <Routes>
                        <Route path="/login" element={<Login />} />
                        <Route path="/dashboard" element={isAuthenticated ? <Dashboard /> : <Navigate to="/login" />} />
                        <Route path="/tickets" element={isAuthenticated ? <Tickets /> : <Navigate to="/login" />} />
                        <Route path="/customers" element={isAuthenticated ? <Customers /> : <Navigate to="/login" />} />
                        <Route path="/buildings" element={isAuthenticated ? <Buildings /> : <Navigate to="/login" />} />
                        <Route path="/workers" element={isAuthenticated ? <Workers /> : <Navigate to="/login" />} />
                        {/* Varsayılan olarak login sayfasına yönlendir */}
                        <Route path="*" element={<Navigate to="/login" />} />
                    </Routes>
                </Container>
            </Grid>
        </Grid>
    );
};

export default App;
