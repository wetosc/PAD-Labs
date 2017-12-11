const cassandra = require('cassandra-driver');
const client = new cassandra.Client({ contactPoints: ['localhost'], keyspace: 'lab3' });
client.connect()

const qWineAll = 'SELECT id FROM wine';
const qCellarAll = 'SELECT id FROM cellar';

const qNewWine = 'INSERT INTO wine(color, flavor, name, price) VALUES (?, ?, ?, ?)'

var a = false
var b = false

var items1
var items2

client.execute(qWineAll, [], function(err, result) {
    console.log(err)
    a = true
    items1 = result.rows
    completion()
}); 

client.execute(qCellarAll, [], function(err, result) {
    console.log(err)
    b = true
    items2 = result.rows
    completion()
});


function completion() {
    if (!a || !b) {  return  }

    var randomID1 = items1.sort(() => .5 - Math.random()).slice(0,50).map((x) => x.id)
    var randomID2 = items2.sort(() => .5 - Math.random()).slice(0,50).map((x) => x.id)
    
    randomID1 = randomID1.concat(items1.sort(() => .5 - Math.random()).slice(0,50).map((x) => x.id))
    randomID2 = randomID2.concat(items2.sort(() => .5 - Math.random()).slice(0,50).map((x) => x.id))

    var queries = Array()
    for (let i = 0; i < 100; i++) {
        const el1 = randomID1[i]
        const el2 = randomID2[i]
        queries.push({query: 'INSERT INTO wine_to_cellar (wine_id, cellar_id) VALUES(?, ?)', params: [el1, el2]})
    }

    client.batch(queries, { prepare: true }, function (err) {
        console.log(err)    
    });
}