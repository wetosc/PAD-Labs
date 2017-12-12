var express = require('express');
var router = express.Router();
var db = require('../controllers/db')

router.get('/', function(req, res, next) {
  db.getAllCellar( function (err, data) {
    console.log(err)
    res.json(data)
  })
});

router.get('/:id', function(req, res, next) {
  db.getCellarByID(req.params.id, function (err, data) {
    res.json(data)
  })
});

router.post('/', function(req, res, next) {
  db.insertCellar( req.body, function(err, data) {
    res.json(data)
  })
})

module.exports = router;
