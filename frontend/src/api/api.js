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

// Müşteri listesi getir
export const getCustomers = async () => {
    return await axios.get(`${API_URL}/customers`);
};

// Çalışan listesi getir
export const getWorkers = async () => {
    return await axios.get(`${API_URL}/workers`);
};

// Bina listesi getir
export const getBuildings = async () => {
    return await axios.get(`${API_URL}/buildings`);
};
