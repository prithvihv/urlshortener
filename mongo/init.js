db.tinyURLs.createIndex( { "createdat": 1 }, { expireAfterSeconds: 86400 }) 
db.tinyURLs.createIndex( { "tinyurluid": 1 }, { unique: true } )