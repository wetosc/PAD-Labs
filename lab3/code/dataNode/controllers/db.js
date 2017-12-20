const Sequelize = require('sequelize');
const Op = Sequelize.Op
const sequelize = new Sequelize('pad', 'pad', 'PAD', {
    host: 'localhost',
    dialect: 'postgres',
    logging: false,
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

exports.getAllWine = function (params, callback) {
    params = limitBuilder(params)
    params = whereBuilder(params, "name", params.qName)
    params = whereBuilder(params, "flavor", params.qFlavor)
    params = whereBuilder(params, "color", params.qColor)
    console.log(params)
    Wine.findAll(params).then(rows => {
        callback(null, rows)
    });
}

exports.getAllCellar = function (params, callback) {
    params = limitBuilder(params)
    params = whereBuilder(params, "name", params.qName)
    params = whereBuilder(params, "location", params.qLocation)
    params = whereBuilder(params, "owner", params.qOwner)
    console.log(params)
    Cellar.findAll(params).then(rows => {
        callback(null, rows)
    });
}
function limitBuilder(params) {
    params.offset = params.offset || 0
    return params
}
function whereBuilder(params, field, value) {
    if (!value) { return params }
    if (!params.where) {params.where = {}}
    params.where[field] = {
        [Op.iLike]: "%"+ value + "%"
    }
    return params
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
    Wine.create(d)
    .then(rows => {
        callback(null, rows)
    })
    .catch(error => {
        callback(error, null)
    })
}

exports.insertCellar = function (d, callback) {
    Cellar.create(d)
    .then(rows => {
        callback(null, rows)
    })
    .catch(error => {
        callback(error, null)
    })
}

// hack to remove unwanted autogenerated fields
function filterIntermediary(array, subParam) {
    for (let j = 0; j < array[subParam].length; j++) {
        array[subParam][j]["wine_cellar"] = undefined
    }
    return array
}