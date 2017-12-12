var mysql = require('mysql');

var client = mysql.createPool({
  connectionLimit : 100,
  host: "localhost",
  user: "pad",
  password: "PAD",
  database: "pad3"
});

const uuid = require('uuid/v4');


const qWineAll = 'SELECT id, color, flavor, name, price FROM wine';
const qWineID =  'SELECT id, color, flavor, name, price FROM wine WHERE id = ?';
const qWineCellar =  'select c.* from cellar as c inner join wine_to_cellar as w2c on c.id = w2c.cellar_id where w2c.wine_id = ?';

const qCellarAll = 'SELECT id, area, location, name, owner FROM cellar';
const qCellarID =  'SELECT id, area, location, name, owner FROM cellar WHERE id = ?';
const qCellarWine =  'select w.* from wine as w inner join wine_to_cellar as w2c on w.id = w2c.wine_id where w2c.cellar_id = ?';

const qNewWine = 'INSERT INTO wine(id, color, flavor, name, price) VALUES (?, ?, ?, ?, ?)'
const qNewCellar = 'INSERT INTO cellar(id, area, location, name, owner) VALUES (?, ?, ?, ?, ?)'

exports.getAllWine = function (callback) {
    client.query(qWineAll, [ ], function(err, rows, columns) {
        console.log(err)
        callback(err, rows)
    });
}

exports.getWineByID = function (id, callback) {
    var collector = {data1: null, data2: null}
    var myErr
    
    function myCallback() {
        if (collector.data1 && collector.data2) {
            var obj = collector.data1
            obj.cellars = collector.data2
            return callback(myErr, obj)
        }
    }
    client.query(qWineID, [ id ], function(err, rows, columns) {
        console.log(err)
        collector.data1 = rows[0]
        myErr = myErr || err
        myCallback()
    });
    client.query(qWineCellar, [id], function(err, rows, columns) {
        console.log(err)
        collector.data2 = rows
        myErr = myErr || err
        myCallback()
    })
}

exports.getAllCellar = function (callback) {
    client.query(qCellarAll, [], function(err, rows, columns) {
        console.log(err)
        callback(err, rows)
    });
}

exports.getCellarByID = function (id, callback) {
    var collector = {data1: null, data2: null}
    var myErr
    
    function myCallback() {
        if (collector.data1 && collector.data2) {
            var obj = collector.data1
            obj.wines = collector.data2
            return callback(myErr, obj)
        }
    }


    client.query(qCellarID, [ id ], function(err, rows, columns) {
        console.log(err)
        collector.data1 = rows[0]
        myErr = myErr || err
        myCallback()
    })

    client.query(qCellarWine, [id], function(err, rows, columns) {
        console.log(err)
        collector.data2 = rows
        myErr = myErr || err
        myCallback()
    })    
}

exports.insertWine = function (d, callback) {
    d.id = uuid()
    d.price = parseFloat(d.price)
    client.query(qNewWine, [d.id, d.color, d.flavor, d.name, d.price], function(err, rows, columns) {
        console.log(err)
        callback(err, d)
    })
}

exports.insertCellar = function (d, callback) {
    d.id = uuid()
    d.area = parseFloat(d.area)
    client.query(qNewCellar, [d.id, d.area, d.location, d.name, d.owner], function(err, rows, columns) {
        console.log(err)
        callback(err, d)
    })
}
