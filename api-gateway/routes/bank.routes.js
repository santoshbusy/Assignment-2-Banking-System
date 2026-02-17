const express = require('express');
const router = express.Router();
const goClient = require('../services/goClient');

const authMiddleware = require('../middleware/auth.middleware');

// Create Bank
router.post('/', authMiddleware, async (req, res) => {
    try {
        const response = await goClient.post('/banks', req.body);
        res.status(response.status).json(response.data);
    } catch (error) {
        res.status(error.response?.status || 500).json({
            success: false,
            message: error.message
        });
    }
});

// Get All Banks
router.get('/', authMiddleware, async (req, res) => {
    try {
        const response = await goClient.get('/banks');
        res.status(response.status).json(response.data);
    } catch (error) {
        res.status(error.response?.status || 500).json({
            success: false,
            message: error.message
        });
    }
});

module.exports = router;