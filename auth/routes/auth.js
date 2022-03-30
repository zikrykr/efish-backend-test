var express = require('express');
var router = express.Router();

const authController = require('../api/auth/auth_controller')

router.post('/register', (req, res, next) => {
  return authController.Register(req, res)
});

router.post('/login', (req, res, next) => {
  return authController.Login(req, res)
})

router.get('/verify-token', (req, res, next) => {
  return authController.Verify(req, res)
})

module.exports = router;
