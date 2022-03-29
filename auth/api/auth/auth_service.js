const fs = require('fs')
const path = require('path')
const dateTimeHelper = require('../../helpers/datetime')

const jsonFile = path.join(__dirname, '../../auth.json')

const PASSWORD_LENGTH = 4

const CreateUser = (phone, name, role) => {
    const user = {
        phone,
        name,
        role,
        createdAt: Date.now(),
        password: GeneratePassword(PASSWORD_LENGTH)
    }

    let users = GetUsers()

    users.push(user)

    users = JSON.stringify({ users: users })

    fs.writeFileSync(jsonFile, users, 'utf8')

    user.createdAt = dateTimeHelper.DateNumberConverter(user.createdAt)

    return user
}

const GetUsers = () => {
    let users = fs.readFileSync(jsonFile, 'utf8')
    let parsedUsers = JSON.parse(users)

    return parsedUsers.users
}

const GeneratePassword = (length) => {
    let password = ""
    const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
    
    for (let i = 0; i < length; i++) {
        password += charset.charAt(Math.floor(Math.random() * charset.length))
    }

    return password
}

const GetUser = (phone) => {
    let users = GetUsers()
    let userData = null

    users.forEach(user => {
        if(!userData && user.phone == phone) {
            userData = user
        }
    })

    return userData
}

module.exports = {
    CreateUser,
    GetUsers,
    GeneratePassword,
    GetUser
}