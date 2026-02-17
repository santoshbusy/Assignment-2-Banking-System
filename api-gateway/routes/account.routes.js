const express = require('express');
const router = express.Router();
const goClient = require('../services/goClient');

router.post('/', async (req, res) => {
    const response = await goClient.post('/accounts', req.body);
    res.status(response.status).json(response.data);
});

router.post('/:id/deposit', async (req, res) => {
    const response = await goClient.post(`/accounts/${req.params.id}/deposit`, req.body);
    res.status(response.status).json(response.data);
});

router.post('/:id/withdraw', async (req, res) => {
    const response = await goClient.post(`/accounts/${req.params.id}/withdraw`, req.body);
    res.status(response.status).json(response.data);
});

router.get('/:id', async (req, res) => {
    const response = await goClient.get(`/accounts/${req.params.id}`);
    res.status(response.status).json(response.data);
});

router.get('/:id/transactions', async (req, res) => {
    const response = await goClient.get(`/accounts/${req.params.id}/transactions`);
    res.status(response.status).json(response.data);
});

router.post('/transfer', async (req, res) => {
    const response = await goClient.post('/accounts/transfer', req.body);
    res.status(response.status).json(response.data);
});

module.exports = router;