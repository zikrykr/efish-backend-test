const svc = require('./auth_service')
const authHelper = require('../../helpers/authentication')

const Register = (req, res) => {
    const {
        phone = null,
        name = null,
        role = null
    } = req.body

    if (!name || name == 0){
        return res.json({
            error: {
                message: "invalid name"
            }
        })
    }

    if (!phone || phone == 0) {
        return res.json({
            error: {
                message: "invalid phone"
            }
        })
    }

    if (!role || role == 0) {
        return res.json({
            error: {
                message: "invalid role"
            }
        })
    }

    const isUserExist = svc.GetUser(phone)

    if (isUserExist) {
        return res.json({
            error: {
                message: "user already exist"
            }
        })
    }

    const result = svc.CreateUser(phone, name, role)

    return res.json(result)
}

const Login = (req, res) => {
    const { phone, password } = req.body

    if (!phone) {
        return res.status(401).json({
            error: {
                message: "phone is required"
            }
        })
    }

    if (!password) {
        return res.status(401).json({
            error: {
                message: "password is required"
            }
        })
    }

    const user = svc.GetUser(phone)
    if (!user) {
        return res.json({
            error: {
                message: "user not found"
            }
        })
    }

    if (user.password != password) {
        return res.json({
            error: {
                message: "wrong password"
            }
        })
    }

    const accessToken = authHelper.GenerateAccessToken(user.name, user.phone, user.role, user.createdAt)

    return res.json({
        phone: user.phone,
        accessToken: accessToken
    })
}

const Verify = (req, res) => {
    const token = authHelper.GetToken(req)
    if(!token) {
        return res.status(401).json({
            error: {
                message: "authorization required"
            }
        })
    }

    const privateClaim = authHelper.JWTClaim(token)
    if(!privateClaim) {
        return res.status(403).json({
            error: {
                message: "invalid token"
            }
        })
    }
    return res.json(privateClaim)
}

module.exports = {
    Register,
    Login,
    Verify
}