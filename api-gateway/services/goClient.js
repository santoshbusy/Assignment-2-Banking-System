const axios = require('axios');
require('dotenv').config();

const goClient = axios.create({
    baseURL: process.env.GO_BACKEND_URL,
    timeout: 5000,
});

module.exports = goClient;