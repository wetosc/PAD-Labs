var mysql = require('mysql');

var client = mysql.createPool({
  connectionLimit : 100,
  host: "localhost",
  user: "pad",
  password: "PAD",
  database: "pad3"
});

const qWineRandom = 'select id from wine order by rand();'
const qCellarRandom = 'select id from cellar order by rand();'

var a = false
var b = false

var items1
var items2

client.query(qWineRandom, [], function(err, fields, columns) {
    console.log(err)
    a = true
    items1 = fields
    completion()
}); 

client.query(qCellarRandom, [], function(err, fields, columns) {
    console.log(err)
    b = true
    items2 = fields
    completion()
});


function completion() {
    if (!a || !b) {  return  }

    var randomID1 = items1.sort(() => .5 - Math.random()).slice(0,50).map((x) => x.id)
    var randomID2 = items2.sort(() => .5 - Math.random()).slice(0,50).map((x) => x.id)
    
    randomID1 = randomID1.concat(items1.sort(() => .5 - Math.random()).slice(0,50).map((x) => x.id))
    randomID2 = randomID2.concat(items2.sort(() => .5 - Math.random()).slice(0,50).map((x) => x.id))

    for (let i = 0; i < 100; i++) {
        const el1 = randomID1[i]
        const el2 = randomID2[i]
        client.query('INSERT INTO wine_to_cellar (wine_id, cellar_id) VALUES(?, ?)',[el1, el2])
    }
}