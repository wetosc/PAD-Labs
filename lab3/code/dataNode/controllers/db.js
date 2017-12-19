const Sequelize = require('sequelize');
const sequelize = new Sequelize('pad', 'pad', 'PAD', {
    host: 'localhost',
    dialect: 'postgres',
    pool: {
      max: 5,
      min: 0,
      acquire: 30000,
      idle: 10000
    },
  });

const uuid = require('uuid/v4');

const Wine = sequelize.define('wine', {
    'id': {
        type: Sequelize.UUID,
        primaryKey: true,
        allowNull: false,
        defaultValue: Sequelize.UUIDV4
    },
    'color': Sequelize.STRING,
    'flavor': Sequelize.STRING,
    'name': Sequelize.STRING,
    'price': Sequelize.FLOAT
})

const Cellar = sequelize.define('cellar', {
    'id': {
        type: Sequelize.UUID,
        primaryKey: true,
        allowNull: false,
        defaultValue: Sequelize.UUIDV4
    },
    'area': Sequelize.FLOAT,
    'location': Sequelize.STRING,
    'name': Sequelize.STRING,
    'owner': Sequelize.STRING
})

Cellar.belongsToMany(Wine, {through: 'wine_cellar'});
Wine.belongsToMany(Cellar, {through: 'wine_cellar'});

sequelize.sync({force: false})

exports.sequelize = sequelize
exports.Wine = Wine
exports.Cellar = Cellar

exports.getAllWine = function (callback) {
    Wine.findAll().then(rows => {
        callback(null, rows)
    });
}

exports.getAllCellar = function (callback) {
    Cellar.findAll().then(rows => {
        callback(null, rows)
    });
}

exports.getWineByID = function (id, callback) {
    Wine.findById(id, {include: Cellar})
    .then(rows => {
        callback(null, filterIntermediary(rows.toJSON(), "cellars"))
    })
    .catch(error => {
        callback(error, null)
    })
}

exports.getCellarByID = function (id, callback) {
    Cellar.findById(id, {include: Wine})
    .then(rows => {
        callback(null, filterIntermediary(rows.toJSON(), "wines"))
    })
    .catch(error => {
        callback(error, null)
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


function filterIntermediary(array, subParam) {
    for (let j = 0; j < array[subParam].length; j++) {
        array[subParam][j]["wine_cellar"] = undefined
    }
    return array
}