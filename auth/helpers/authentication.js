const jwt = require('jsonwebtoken')

const JWT_SECRET = process.env.JWT_SECRET || 'jwt_secret_default'
const JWT_ALGORITHM = 'HS256'
const EXPIRATION_TIME = 10 * 60 * 60

const GetToken = (req) => {
    let authorization = req.headers['authorization'] || null
    if (!authorization) {
        return null
    }

    return authorization.split(" ")[1]
}

const JWTClaim = (token) => {
    return jwt.decode(token, JWT_SECRET, { algorithm: JWT_ALGORITHM })
}

const GenerateAccessToken = (name, phone, role, createdAt) => {
    const exp = Date.now() + EXPIRATION_TIME
    const jwtPayload = {
        name, phone, role, createdAt, exp
    }

    return jwt.sign(jwtPayload, JWT_SECRET, { algorithm: JWT_ALGORITHM })
}

module.exports = {
    GetToken,
    JWTClaim,
    GenerateAccessToken
}