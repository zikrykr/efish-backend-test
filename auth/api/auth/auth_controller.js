const svc = require('./auth_service')

const register = (req, res) => {
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

    const isUserExist = svc.CheckUser(phone)

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

module.exports = {
    register
}