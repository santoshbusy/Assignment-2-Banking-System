require('dotenv').config();
const express = require('express');

const app = express();   // ✅ FIRST create app

app.use(express.json()); // ✅ then middlewares

// Routes
const bankRoutes = require('./routes/bank.routes');
const branchRoutes = require('./routes/branch.routes');
const customerRoutes = require('./routes/customer.routes');
const accountRoutes = require('./routes/account.routes');
const loanRoutes = require('./routes/loan.routes');
const authRoutes = require('./routes/auth.routes'); // <-- your JWT route

// Register routes AFTER app is created
app.use('/auth', authRoutes);
app.use('/banks', bankRoutes);
app.use('/branches', branchRoutes);
app.use('/customers', customerRoutes);
app.use('/accounts', accountRoutes);
app.use('/loans', loanRoutes);

// Root test route
app.get('/', (req, res) => {
    res.send('API Gateway Running');
});

const PORT = process.env.PORT || 3000;

app.listen(PORT, () => {
    console.log(`Node API Gateway running on port ${PORT}`);
});