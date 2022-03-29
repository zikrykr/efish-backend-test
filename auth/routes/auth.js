var express = require('express');
var router = express.Router();

const authController = require('../api/auth/auth_controller')

router.post('/register', function(req, res, next) {
  return authController.register(req, res)
});

module.exports = router;
