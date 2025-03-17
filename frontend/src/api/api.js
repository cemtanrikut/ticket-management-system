import axios from "axios";

const API_URL = process.env.REACT_APP_API_URL;

// Kullanıcı giriş işlemi
export const login = async (email, password) => {
    return await axios.post(`${API_URL}/login`, { email, password });
};

// Ticket listesi getir
export const getTickets = async () => {
    return await axios.get(`${API_URL}/tickets`);
};

// Yeni ticket oluştur
export const createTicket = async (ticketData) => {
    return await axios.post(`${API_URL}/tickets`, ticketData);
};

// Ticket detaylarını getir
export const getTicketById = async (ticketId) => {
    return await axios.get(`${API_URL}/tickets/${ticketId}`);
};

// Ticket için mesajları getir
export const getMessages = async (ticketId) => {
    return await axios.get(`${API_URL}/messages/${ticketId}`);
};

// Ticket için mesaj gönder
export const sendMessage = async (ticketId, messageData) => {
    return await axios.post(`${API_URL}/messages/${ticketId}`, messageData);
};
