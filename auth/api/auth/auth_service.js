const fs = require('fs')
const path = require('path')

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

    return user
}

const CheckUser = (phone) => {
    let users = GetUsers()
    let isUserExist = false

    users.forEach(user => {
        if (user.phone == phone && !isUserExist) {
            isUserExist = true
        }
    });

    return isUserExist
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

module.exports = {
    CreateUser,
    GetUsers,
    CheckUser,
    GeneratePassword
}