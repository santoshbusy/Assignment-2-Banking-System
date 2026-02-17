const jwt = require('jsonwebtoken');
const bcrypt = require('bcryptjs');

// Dummy user (for now)
const user = {
    id: 1,
    email: "admin@bank.com",
    password: bcrypt.hashSync("password123", 10)
};

exports.login = async (req, res) => {
    const { email, password } = req.body;

    if (email !== user.email) {
        return res.status(401).json({ message: "Invalid email" });
    }

    const isMatch = bcrypt.compareSync(password, user.password);

    if (!isMatch) {
        return res.status(401).json({ message: "Invalid password" });
    }

    const token = jwt.sign(
        { id: user.id, email: user.email },
        process.env.JWT_SECRET,
        { expiresIn: "1h" }
    );

    res.json({ token });
};