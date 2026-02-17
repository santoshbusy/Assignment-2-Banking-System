const express = require('express');
const router = express.Router();
const goClient = require('../services/goClient');

router.post('/', async (req, res) => {
    try {
        const response = await goClient.post('/customers', req.body);
        res.status(response.status).json(response.data);
    } catch (error) {
        res.status(error.response?.status || 500).json({
            success: false,
            message: error.message,
        });
    }
});

router.get('/:id', async (req, res) => {
    try {
        const response = await goClient.get(`/customers/${req.params.id}`);
        res.status(response.status).json(response.data);
    } catch (error) {
        res.status(error.response?.status || 500).json({
            success: false,
            message: error.message,
        });
    }
});

module.exports = router;