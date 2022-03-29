var express = require('express');
var router = express.Router();

const authController = require('../api/auth/auth_controller')

router.post('/register', (req, res, next) => {
  return authController.register(req, res)
});

router.post('/login', (req, res, next) => {
  return authController.login(req, res)
})

module.exports = router;
