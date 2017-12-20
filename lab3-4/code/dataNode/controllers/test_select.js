const cassandra = require('cassandra-driver');
const client = new cassandra.Client({ contactPoints: ['localhost'], keyspace: 'lab3' });
client.connect()

const qWineAll = 'SELECT * FROM wine_to_cellar';

client.execute(qWineAll, [], function(err, result) {
    result.rows.forEach(element => {
        console.log(element);
    });
}); 