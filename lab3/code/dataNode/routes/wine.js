var express = require('express');
var router = express.Router();
var db = require('../controllers/db')

router.get('/', function(req, res, next) {
  db.getAllWine( function (err, data) {
    res.json(data)
  })
});

router.get('/:id', function(req, res, next) {
  db.getWineByID(req.params.id, function (err, data) {
    res.json(data)
  })
});

router.post('/', function(req, res, next) {
  db.insertWine( req.body, function(err, data) {
    res.json(data)
  })
})

module.exports = router;
