const express = require('express');
const router = express.Router();
const goClient = require('../services/goClient');

router.post('/', async (req, res) => {
    const response = await goClient.post('/loans', req.body);
    res.status(response.status).json(response.data);
});

router.post('/:id/repay', async (req, res) => {
    const response = await goClient.post(`/loans/${req.params.id}/repay`, req.body);
    res.status(response.status).json(response.data);
});

router.get('/:id', async (req, res) => {
    const response = await goClient.get(`/loans/${req.params.id}`);
    res.status(response.status).json(response.data);
});

module.exports = router;