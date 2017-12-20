var express = require('express');
var router = express.Router();
var db = require('../controllers/db')

router.get('/', function(req, res, next) {
  db.getAllWine(req.query, function (err, data) {
    if (err != null) {    console.log(err); return res.status(500).end();    }
    res.json(data)
  })
});

router.get('/:id', function(req, res, next) {
  db.getWineByID(req.params.id, function (err, data) {
    if (err != null) {    console.log(err); return res.status(500).end();    }
    res.json(data)
  })
});

router.post('/', function(req, res, next) {
  db.insertWine( req.body, function(err, data) {
    if (err != null) {    console.log(err); return res.status(500).end();    }
    res.json(data)
  })
})

module.exports = router;
