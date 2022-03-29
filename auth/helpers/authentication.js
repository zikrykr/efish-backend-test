const jwt = require('jsonwebtoken')

const JWT_ALGORITHM = 'HS256'
const EXPIRATION_TIME = 10 * 60 * 60

const GenerateAccessToken = (name, phone, role, createdAt) => {
    const secretKey = process.env.JWT_SECRET || 'jwt_secret_default'
    const exp = Date.now() + EXPIRATION_TIME
    const jwtPayload = {
        name, phone, role, createdAt, exp
    }

    return jwt.sign(jwtPayload, secretKey, { algorithm: JWT_ALGORITHM })
}

module.exports = {
    GenerateAccessToken
}