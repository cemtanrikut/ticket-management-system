import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { login } from "../api/api";
import { Container, TextField, Button, Typography } from "@mui/material";

const Login = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const navigate = useNavigate();

    const handleLogin = async () => {
        try {
            const response = await login(email, password);
            localStorage.setItem("token", response.data.token);
            navigate("/dashboard");
        } catch (error) {
            alert("Login failed!");
        }
    };

    return (
        <Container>
            <Typography variant="h4">Login</Typography>
            <TextField label="Email" fullWidth margin="normal" onChange={(e) => setEmail(e.target.value)} />
            <TextField label="Password" type="password" fullWidth margin="normal" onChange={(e) => setPassword(e.target.value)} />
            <Button variant="contained" color="primary" fullWidth onClick={handleLogin}>Login</Button>
        </Container>
    );
};

export default Login;
