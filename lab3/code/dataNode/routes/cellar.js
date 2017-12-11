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
  var id = req.params.id
  if ( isNaN(id) || id <= 0) {
    res.status(400)
    res.end()
    return
  }
  db.getCellarByID(id, function (err, data) {
    console.log(err)
    res.json(data)
  })
});

module.exports = router;
