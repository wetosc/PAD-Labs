const Sequelize = require('sequelize');
const db = require('./db')
const sequelize = db.sequelize

var wines =
[{"id":"f48455b9-57f7-4b1d-a15c-d534d5fd510c","name":"Blue Wild Indigo","flavor":"Nut - Pine Nuts, Whole","color":"Red","price":154.49},
{"id":"c961a4f8-e98a-4a3f-a18a-7c5f51644039","name":"Bullich","flavor":"Wine - Red, Pinot Noir, Chateau","color":"Yellow","price":484.4},
{"id":"40c1201f-0b57-4c37-a2c1-771c0f75877d","name":"Hawthorn","flavor":"Juice - Lime","color":"Khaki","price":87.01},
{"id":"fdf758c9-5ebe-4dfe-83fc-c4e8464ff651","name":"Matting Rosette Grass","flavor":"Sugar - White Packet","color":"Purple","price":476.49},
{"id":"eaac5e9b-1cc3-4a02-ba25-14660a5e1de7","name":"Macoun's Heterocladium Moss","flavor":"Tomatoes - Grape","color":"Yellow","price":398.22},
{"id":"9d61d612-d6c3-4f06-b827-936256124452","name":"Toumey's Century Plant","flavor":"Campari","color":"Fuscia","price":495.14},
{"id":"f9c63756-c951-4acc-a4c1-b5df73387a01","name":"Alpine Forget-me-not","flavor":"Crab - Claws, 26 - 30","color":"Khaki","price":333.87},
{"id":"089b481d-6447-47fa-ad1c-a29372f14423","name":"Cliff Thistle","flavor":"Lumpfish Black","color":"Goldenrod","price":68.7},
{"id":"2e0fcb4d-f6e8-452b-91b1-aafa1f242e5a","name":"Twospike Crabgrass","flavor":"Pepsi, 355 Ml","color":"Maroon","price":164.75},
{"id":"4ea23002-022a-4430-aafe-f2ca428700ce","name":"Browneyes","flavor":"Wine - Mas Chicet Rose, Vintage","color":"Pink","price":231.47},
{"id":"2227391d-3b7b-4b12-8439-372f12252211","name":"Saltwater False Willow","flavor":"Sloe Gin - Mcguinness","color":"Aquamarine","price":27.52},
{"id":"d0245f36-e3b4-40af-8632-a06fc91f5b72","name":"Bristle Fern","flavor":"Olives - Black, Pitted","color":"Purple","price":360.61},
{"id":"9e2a981e-2de7-4325-aa6b-da0ce1d15f00","name":"Whitetop-box","flavor":"Irish Cream - Butterscotch","color":"Teal","price":192.09},
{"id":"bb0c60ca-0b62-47c6-9e85-c66b96e07c93","name":"Wright's Thistle","flavor":"Pepper - Red, Finger Hot","color":"Maroon","price":212.21},
{"id":"5fc938fd-9602-4125-86e1-60558e43e1b2","name":"Hybrid Hickory","flavor":"Cranberries - Fresh","color":"Mauv","price":101.2},
{"id":"92c483e3-0d40-4a5e-a384-c027d14d5657","name":"Clinton's Woodfern","flavor":"Eggroll","color":"Yellow","price":206.3},
{"id":"224f16d7-6c21-4e8b-b641-560e06cb27a2","name":"Orange Lichen","flavor":"French Kiss Vanilla","color":"Pink","price":95.09},
{"id":"c1fbbb57-fe34-4c9f-8ff1-9f54db8f0543","name":"Purple Stickpea","flavor":"Sauce - Cranberry","color":"Teal","price":11.26},
{"id":"d4bc05ba-8ada-4a98-97c1-74a245831ee4","name":"American Snowbell","flavor":"Peach - Halves","color":"Purple","price":278.88},
{"id":"df673db3-87fe-4f45-94e1-89268277a7ba","name":"Myrtle Willow","flavor":"Tea - Decaf 1 Cup","color":"Fuscia","price":92.64}]

var cellars = 
[{"id":"5f1efec2-31da-401c-a7e1-2aa9ec6a318b","name":"Gabspot","location":"Tsagaan-Olom","owner":"Alexandr Hatherleigh","area":6.53},
{"id":"498d7261-a8a2-4bcf-be20-f0337ccc38a3","name":"Eire","location":"Pshada","owner":"Tiffy McClintock","area":6.06},
{"id":"49a48eb1-8b07-45fe-806d-2b29bcad5d49","name":"Topiclounge","location":"Varpaisjärvi","owner":"Shari Jowling","area":3.61},
{"id":"a6d07f4b-96c5-4de4-ac7f-af906864ce31","name":"Avamm","location":"Krajenka","owner":"Cecilius Eayres","area":5.9},
{"id":"3c538907-a72d-4f12-bd49-320e69235e9f","name":"InnoZ","location":"Tsarychanka","owner":"Josefina Nunes Nabarro","area":3.42},
{"id":"7731cf71-4e99-4efe-9a4c-3322ca3013ec","name":"Trudoo","location":"Wakefield","owner":"Pernell Mapam","area":6.22},
{"id":"7f9347f7-bcfe-4b35-8fc6-62e93d45bf70","name":"Realbridge","location":"Wassu","owner":"Gertie Ewbank","area":7.43},
{"id":"c9436d5d-7a3e-4471-a7f9-530356f1ad45","name":"Chatterbridge","location":"Berlin","owner":"Rodrigo Purse","area":3.55},
{"id":"c4dff387-c5c5-4d80-8ba5-84a065e64ffc","name":"Zoombeat","location":"Bayangol","owner":"Kennie Yurkiewicz","area":3.95},
{"id":"95801c93-d454-4f00-a328-5b96c5fbb983","name":"Wikizz","location":"Francisco Morato","owner":"Annalise Grimsditch","area":7.21},
{"id":"fd64bb1d-40a2-449b-8112-7a141c67e876","name":"Latz","location":"Blagodarnyy","owner":"Elicia Lydford","area":9.05},
{"id":"508ae62c-a4c6-4726-9766-9ec70e65f527","name":"Divanoodle","location":"Lens","owner":"Mic Kuban","area":4.43},
{"id":"154394f5-5037-4f05-b6da-a8e00e9185db","name":"Podcat","location":"Évosmos","owner":"Jeralee Haddrill","area":1.76},
{"id":"1d263499-e9dc-481c-8914-39a334f65f45","name":"Edgepulse","location":"Abay","owner":"Kaleb Tomaszczyk","area":7.36},
{"id":"53c1ff59-9319-4db4-baf8-3b11610853bd","name":"Abatz","location":"Tambulatana","owner":"Annadiana Tavener","area":7.92},
{"id":"e9f7c036-3b14-41a6-84d7-4f91c92efb39","name":"Innojam","location":"Tunbao","owner":"Mohandas Grace","area":8.01},
{"id":"bb463a8d-b9ce-4f37-8c7d-bc6885042a06","name":"Thoughtblab","location":"Huayana","owner":"Dana Chiese","area":8.03},
{"id":"6dcb3618-0ca2-4be4-a696-094a2925885e","name":"Realfire","location":"Roriz","owner":"Noell Brigshaw","area":3.31},
{"id":"940ed2ac-1309-4273-9d80-002938248c05","name":"Rhyzio","location":"Lëbushë","owner":"Ainslie Spick","area":4.82},
{"id":"d0d2ddce-d03e-44b8-9838-a528d21acc96","name":"InnoZ","location":"Tečovice","owner":"Dell Heam","area":8.42}]


// db.Wine.bulkCreate(wines).then(response => {
//     console.log(response)
// })

// db.Cellar.bulkCreate(cellars).then(response => {
//     console.log(response)
// })


// setTimeout(fillRelations, 1000);

function fillRelations() {
    db.Wine.findAll().then(wines => {
        db.Cellar.findAll().then(cellars => {
            for (let i = 0; i < wines.length; i++) {
                const wine = wines[i];
                let sel = shuffle(cellars, 3)//.map(x=> {x.id}).filter(x=> {x != null})
                wine.setCellars(sel)
                wine.save()
            }   
        })
    })
}

function shuffle (array, n) {
    const shuffled = array.sort(() => .5 - Math.random());// shuffle  
    let selected = shuffled.slice(0,n) ;
    return selected
}
