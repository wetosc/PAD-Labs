const cassandra = require('cassandra-driver');
const uuid = require('uuid/v4');

const client = new cassandra.Client({ contactPoints: ['localhost'], keyspace: 'lab3' });
client.connect()

const qWineAll = 'SELECT id, color, flavor, name, price FROM wine';
const qWineID =  'SELECT id, color, flavor, name, price FROM wine WHERE id = ?';

const qCellarAll = 'SELECT id, area, location, name, owner FROM cellar';
const qCellarID =  'SELECT id, area, location, name, owner FROM cellar WHERE id = ?';

const qNewWine = 'INSERT INTO wine(id, color, flavor, name, price) VALUES (?, ?, ?, ?, ?)'
const qNewCellar = 'INSERT INTO cellar(id, area, location, name, owner) VALUES (?, ?, ?, ?, ?)'

exports.getAllWine = function (callback) {
    client.execute(qWineAll, [ ], function(err, result) {
        console.log(err)
        callback(err, result.rows)
    });
}

exports.getWineByID = function (id, callback) {
    client.execute(qWineID, [ id ],  { prepare : true }, function(err, result) {
        console.log(err)
        callback(err, result.rows[0])
    });    
}

exports.getAllCellar = function (callback) {
    client.execute(qCellarAll, [], function(err, result) {
        console.log(err)
        callback(err, result.rows)
    });
}

exports.getCellarByID = function (id, callback) {
    var total
    client.execute(qCellarID, [ id ],  { prepare : true }, function(err, result) {
        console.log(err)
        callback(err, result.rows[0])
        Object.keys(dictionary).forEach( function(key) {
            total[key] = dictionary[key]
        });
    });    
}



exports.insertWine = function (d, callback) {
    d.id = uuid()
    d.price = parseFloat(d.price)
    client.execute(qNewWine, [d.id, d.color, d.flavor, d.name, d.price], {prepare: true}, function (err, result) {
        console.log(err)
        callback(err, d)
    })
}

exports.insertCellar = function (d, callback) {
    d.id = uuid()
    d.area = parseFloat(d.area)
    client.execute(qNewWine, [d.id, d.area, d.location, d.name, d.owner], {prepare: true}, function (err, result) {
        console.log(err)
        callback(err, d)
    })
}
