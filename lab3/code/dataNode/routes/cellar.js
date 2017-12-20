var express = require('express');
var router = express.Router();
var db = require('../controllers/db')

router.get('/', function(req, res, next) {
  db.getAllCellar(req.query, function (err, data) {
    if (err != null) {    console.log(err); return res.status(500).end();   }
    res.json(data)
  })
});

router.get('/:id', function(req, res, next) {
  db.getCellarByID(req.params.id, function (err, data) {
    if (err != null) {    console.log(err); return res.status(500).end();   }
    res.json(data)
  })
});

router.post('/', function(req, res, next) {
  db.insertCellar( req.body, function(err, data) {
    if (err != null) {    console.log(err); return res.status(500).end();    }
    res.json(data)
  })
})

module.exports = router;
