#  expire in one day

db.tinyURLs.createIndex( { "createdAt": 1 }, { expireAfterSeconds: 86400 }) 

# create unique tinyURLuid
db.tinyURLs.createIndex( { "tinyURLuid": 1 }, { unique: true } )
