const express = require('express');
const router = express.Router();
const goClient = require('../services/goClient');

// Create branch
router.post('/', async (req, res) => {
    try {
        const response = await goClient.post('/branches', req.body);
        res.status(response.status).json(response.data);
    } catch (error) {
        res.status(error.response?.status || 500).json({
            success: false,
            message: error.message,
        });
    }
});

// Get branches by bank
router.get('/', async (req, res) => {
    try {
        const response = await goClient.get('/branches', {
            params: req.query
        });
        res.status(response.status).json(response.data);
    } catch (error) {
        res.status(error.response?.status || 500).json({
            success: false,
            message: error.message,
        });
    }
});

module.exports = router;